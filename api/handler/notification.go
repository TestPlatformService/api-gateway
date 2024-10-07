// file: api/handler/notification.go
package handler

import (
	"api/api/token"
	pb "api/genproto/notification"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: Ishlab chiqarish muhitida bu funksiyani xavfsizroq qiling
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocketMessage struct {
	Type    string      `json:"type"`
	Action  string      `json:"action,omitempty"`
	ID      string      `json:"id,omitempty"`
	Token   string      `json:"token,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Println("WebSocket ulanish so'rovi qabul qilindi")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrader xatosi: %v", err)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("WebSocket ulanishini yopishda xatolik: %v", err)
		}
	}()

	log.Println("WebSocket ulanishi muvaffaqiyatli o'rnatildi")

	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	var userID string

	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
					log.Println("Ping yuborishda xatolik:", err)
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			h.sendNotifications(conn, userID)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Foydalanuvchi chiqib ketdi: %v", err)
				return
			}
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Xabarni o'qishda xatolik: %v", err)
			}
			break
		}

		var msg WebSocketMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("JSON ni ochishda xatolik:", err)
			sendError(conn, "Invalid message format")
			continue
		}

		switch msg.Type {
		case "auth":
			userID, _, err = token.GetUserInfoFromAccessToken(msg.Token)
			if err != nil {
				log.Printf("Noto'g'ri access token: %v", err)
				sendError(conn, "Invalid access token")
				continue
			}
			log.Printf("Foydalanuvchi ID: %s autentifikatsiyadan o'tdi", userID)
			h.sendNotifications(conn, userID)

		case "markAsRead":
			if userID == "" {
				sendError(conn, "Not authenticated")
				continue
			}
			_, err = h.Notification.MarkNotificationAsRead(ctx, &pb.MarkNotificationAsReadReq{NotificationId: msg.ID})
			if err != nil {
				log.Println("O'qilgan deb belgilashda xatolik:", err)
				sendError(conn, "Failed to mark as read")
			} else {
				h.sendNotifications(conn, userID)
			}

		default:
			log.Printf("Noma'lum xabar turi: %s", msg.Type)
			sendError(conn, "Unknown message type")
		}
	}
}

func (h *Handler) sendNotifications(conn *websocket.Conn, userID string) {
	notifications, err := h.Notification.GetAllNotifications(context.Background(), &pb.GetNotificationsReq{UserId: userID})
	if err != nil {
		log.Println("Bildirishnomalarni olishda xatolik:", err)
		sendError(conn, "Failed to get notifications")
		return
	}

	conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err := conn.WriteJSON(WebSocketMessage{Type: "notifications", Payload: notifications}); err != nil {
		log.Println("Bildirishnomalarni yuborishda xatolik:", err)
	}
}

func sendError(conn *websocket.Conn, message string) {
	conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err := conn.WriteJSON(WebSocketMessage{Type: "error", Payload: message}); err != nil {
		log.Printf("Xato xabarini yuborishda muammo: %v", err)
	}
}
