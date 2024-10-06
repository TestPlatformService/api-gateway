// file: api/handler/notification.go
package handler

import (
	"api/api/token"
	pb "api/genproto/notification"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Barcha ulanishlarga ruxsat berish
		return true
	},
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
	defer conn.Close()

	log.Println("WebSocket ulanishi muvaffaqiyatli o'rnatildi")

	var userID string

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Xabarni o'qishda xatolik:", err)
			break
		}

		var msg WebSocketMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("JSON ni ochishda xatolik:", err)
			continue
		}

		switch msg.Type {
		case "auth":
			userID, _, err = token.GetUserInfoFromAccessToken(msg.Token)
			if err != nil {
				log.Printf("Noto'g'ri access token: %v", err)
				conn.WriteJSON(WebSocketMessage{Type: "error", Payload: "Invalid access token"})
				continue
			}
			log.Printf("Foydalanuvchi ID: %s autentifikatsiyadan o'tdi", userID)
			h.sendNotifications(conn, userID)

		case "markAsRead":
			if userID == "" {
				conn.WriteJSON(WebSocketMessage{Type: "error", Payload: "Not authenticated"})
				continue
			}
			_, err = h.Notification.MarkNotificationAsRead(context.Background(), &pb.MarkNotificationAsReadReq{NotificationId: msg.ID})
			if err != nil {
				log.Println("O'qilgan deb belgilashda xatolik:", err)
				conn.WriteJSON(WebSocketMessage{Type: "error", Payload: "Failed to mark as read"})
			} else {
				h.sendNotifications(conn, userID)
			}

		default:
			log.Printf("Noma'lum xabar turi: %s", msg.Type)
		}
	}
}

func (h *Handler) sendNotifications(conn *websocket.Conn, userID string) {
	notifications, err := h.Notification.GetAllNotifications(context.Background(), &pb.GetNotificationsReq{UserId: userID})
	if err != nil {
		log.Println("Bildirishnomalarni olishda xatolik:", err)
		conn.WriteJSON(WebSocketMessage{Type: "error", Payload: "Failed to get notifications"})
		return
	}

	err = conn.WriteJSON(WebSocketMessage{Type: "notifications", Payload: notifications})
	if err != nil {
		log.Println("Bildirishnomalarni yuborishda xatolik:", err)
	}
}
