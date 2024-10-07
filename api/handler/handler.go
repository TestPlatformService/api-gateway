package handler

import (
	"api/genproto/group"
	"api/genproto/notification"
	"api/genproto/question"
	"api/genproto/subject"
	"api/genproto/task"
	"api/genproto/topic"
	"api/genproto/user"
	"log"
	"log/slog"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User           user.UsersClient
	Group          group.GroupServiceClient
	Subject        subject.SubjectServiceClient
	Notification   notification.NotificationsClient
	Topic          topic.TopicServiceClient
	Question       question.QuestionServiceClient
	QuestionOutput question.OutputServiceClient
	QuestionInput  question.InputServiceClient
	TestCase       question.TestCaseServiceClient
	Task           task.TaskServiceClient
	Log            *slog.Logger
	Enforcer       *casbin.Enforcer
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("CORS middleware triggered")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400") // 24 soat

		// WebSocket ulanishlari uchun qo'shimcha headerlar
		if strings.ToLower(c.Request.Header.Get("Connection")) == "upgrade" &&
			strings.ToLower(c.Request.Header.Get("Upgrade")) == "websocket" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
			c.Writer.Header().Set("Sec-Websocket-Extensions", c.Request.Header.Get("Sec-Websocket-Extensions"))
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
