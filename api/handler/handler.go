package handler

import (
	"api/genproto/group"
	"api/genproto/user"
	"log"
	"log/slog"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User     user.UsersClient
	Group    group.GroupServiceClient
	Log      *slog.Logger
	Enforcer *casbin.Enforcer
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Cors middleware triggered")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
