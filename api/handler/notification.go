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

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type WebSocketMessage struct {
	Action string `json:"action"`
	ID     string `json:"id"`
}

func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	accestoken := r.Header.Get("Authorization")
	if accestoken == "" {
		http.Error(w, "Authorization token is required", http.StatusUnauthorized)
		return
	}
	userID, _, err := token.GetUserInfoFromAccessToken(accestoken)
	if err != nil {
		http.Error(w, "Invalid access token", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Send initial notifications
	h.sendNotifications(conn, userID)

	// Start a goroutine to send notifications every 5 seconds
	go func() {
		for {
			time.Sleep(5 * time.Second)
			h.sendNotifications(conn, userID)
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		var msg WebSocketMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("unmarshal:", err)
			continue
		}

		if msg.Action == "markAsRead" {
			_, err = h.Notification.MarkNotificationAsRead(context.Background(), &pb.MarkNotificationAsReadReq{NotificationId: msg.ID})
			if err != nil {
				log.Println("mark as read:", err)
			}
		}

		h.sendNotifications(conn, userID)
	}
}

func (h *Handler) sendNotifications(conn *websocket.Conn, userID string) {
	notifications, err := h.Notification.GetAllNotifications(context.Background(), &pb.GetNotificationsReq{UserId: userID})
	if err != nil {
		log.Println(err)
		return
	}

	err = conn.WriteJSON(notifications)
	if err != nil {
		log.Println(err)
	}
}
