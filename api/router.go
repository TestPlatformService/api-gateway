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
	user := router.Group("/api/user")
	user.Use(middleware.Check)
	user.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		user.POST("/register", hand.Register)
		user.POST("/login", hand.Login)
		user.GET("/getprofile", hand.GetProfile)
		user.GET("/all", hand.GetAllUsers)
		user.PUT("/updateprofile", hand.UpdateProfile)
		user.PUT("/update", hand.UpdateProfileAdmin)
		user.DELETE("/delete/:id", hand.DeleteProfile)
	}

	
	return router
}
