package api

import (
	"api/api/handler"
	"api/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "api/api/docs"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @title User
// @version 1.0
// @description API Gateway
// BasePath: /
func Router(hand *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// user
	user := router.Group("/user")
	user.Use(middleware.Check)
	user.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		user.POST("/user/register")
		user.POST("/user/login")
		user.GET("/user/getprofile")
	}

	return router
}
