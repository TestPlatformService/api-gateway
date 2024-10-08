package handler

import (
	"api/api/token"
	pb "api/genproto/notification"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		log.Printf("WebSocket so'rovi keldi. Origin: %s", r.Header.Get("Origin"))
		return true // Ishlab chiqarish muhitida bu funksiyani xavfsizroq qiling
	},
}

type WebSocketMessage struct {
	Action string `json:"action"`
	ID     string `json:"id"`
	Token  string `json:"token"`
}

func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Printf("WebSocket ulanish so'rovi: %s", r.URL)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade xatosi: %v", err)
		http.Error(w, fmt.Sprintf("WebSocket upgrade failed: %v", err), http.StatusBadRequest)
		return
	}
	defer conn.Close()

	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// Autentifikatsiya
	var userID string
	stopChan := make(chan struct{}) // Go-routine ni to'xtatish uchun channel
	go func() {
		for {
			select {
			case <-stopChan:
				return // Channel yopilganda go-routine to'xtaydi
			default:
				time.Sleep(5 * time.Second)
				h.sendNotifications(conn, userID)
			}
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break // Ulanishni to'xtatish
		}

		var msg WebSocketMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("unmarshal:", err)
			continue
		}

		if msg.Action == "auth" {
			log.Printf("Auth so'rovi keldi. Token: %s", msg.Token)
			userID, _, err = token.GetUserInfoFromAccessToken(msg.Token)
			if err != nil {
				log.Printf("Noto'g'ri access token: %v", err)
				conn.WriteMessage(websocket.TextMessage, []byte("Invalid access token"))
				return
			}
			log.Printf("Foydalanuvchi autentifikatsiyadan o'tdi: %s", userID)
			break
		}
	}

	// Ulanishni saqlash
	h.ConnMutex.Lock()
	h.Connections[userID] = conn
	h.ConnMutex.Unlock()

	defer func() {
		h.ConnMutex.Lock()
		delete(h.Connections, userID)
		h.ConnMutex.Unlock()
		close(stopChan) // Channelni yopish
	}()

	// Dastlabki bildirishnomalarni yuborish
	log.Printf("Dastlabki bildirishnomalar yuborilmoqda userID: %s uchun", userID)
	h.sendNotifications(conn, userID)

	// Ping yuborish uchun go-routine
	go func() {
		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()
		for {
			<-ticker.C
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}()

	// Asosiy xabarlarni o'qish tsikli
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Xatoni o'qish: %v", err)
			}
			break
		}

		var msg WebSocketMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("unmarshal:", err)
			continue
		}

		if msg.Action == "markAsRead" {
			log.Printf("markAsRead so'rovi keldi. ID: %s", msg.ID)
			_, err = h.Notification.MarkNotificationAsRead(context.Background(), &pb.MarkNotificationAsReadReq{NotificationId: msg.ID})
			if err != nil {
				log.Println("mark as read:", err)
				conn.WriteMessage(websocket.TextMessage, []byte("Failed to mark as read"))
			} else {
				log.Println("Xabar o'qilgan deb belgilandi")
				h.sendNotifications(conn, userID)
			}
		}
	}
}

func (h *Handler) sendNotifications(conn *websocket.Conn, userID string) {
	log.Printf("sendNotifications funksiyasi chaqirildi. UserID: %s", userID)
	notifications, err := h.Notification.GetAllNotifications(context.Background(), &pb.GetNotificationsReq{UserId: userID})
	if err != nil {
		log.Printf("Bildirishnomalarni olishda xatolik: %v", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Failed to get notifications"))
		return
	}
	log.Printf("Olingan bildirishnomalar soni: %d", len(notifications.Notifications))

	conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err := conn.WriteJSON(map[string]interface{}{
		"action":        "updateNotifications",
		"notifications": notifications.Notifications,
	}); err != nil {
		log.Printf("Bildirishnomalarni yuborishda xatolik: %v", err)
	} else {
		log.Printf("Bildirishnomalar muvaffaqiyatli yuborildi")
	}
}

func (h *Handler) CreateNotification(ctx context.Context, req *pb.CreateNotificationsReq) (*pb.CreateNotificationsRes, error) {
	res, err := h.Notification.CreateNotification(ctx, req)
	if err != nil {
		return nil, err
	}

	go h.broadcastNewNotification(req.UserId)

	return res, nil
}

func (h *Handler) broadcastNewNotification(userID string) {
	h.ConnMutex.Lock()
	conn, ok := h.Connections[userID]
	h.ConnMutex.Unlock()

	if ok {
		h.sendNotifications(conn, userID)
	}
}
