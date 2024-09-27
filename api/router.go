package api

import (
	"api/api/handler"
	"api/api/middleware"

	_ "api/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @title ALL
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
	user.Use(handler.CORSMiddleware())
	{
		user.POST("/register", hand.Register)
		user.GET("/getprofile", hand.GetProfile)
		user.GET("/all", hand.GetAllUsers)
		user.PUT("/updateprofile", hand.UpdateProfile)
		user.PUT("/update", hand.UpdateProfileAdmin)
		user.DELETE("/delete/:id", hand.DeleteProfile)
		user.DELETE("/photo", hand.DeleteUserPhoto)
		user.POST("/photo", hand.UploadPhotoToUser)
	}

	all := router.Group("/all/user")
	{
		all.POST("/login", hand.Login)
		all.POST("/refresh", hand.Refresh)
	}

	// websocket
	router.GET("/ws", func(c *gin.Context) {
		hand.HandleWebSocket(c.Writer, c.Request)
	})

	group := router.Group("/api/groups")
	group.Use(middleware.Check)
	group.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	group.Use(handler.CORSMiddleware())
	{
		group.POST("/create", hand.CreateGroup)
		group.PUT("/update", hand.UpdateGroup)
		group.DELETE("/delete", hand.DeleteGroup)
		group.GET("/getById", hand.GetGroupById)
		group.GET("/getAll", hand.GetAllGroups)
		group.POST("/add-student", hand.AddStudentToGroup)
		group.DELETE("delete-student", hand.DeleteStudentFromGroup)
		group.POST("add-teacher", hand.AddTeacherToGroup)
		group.DELETE("delete-teacher", hand.DeleteTeacherFromGroup)
		group.GET("student-groups", hand.GetStudentGroups)
		group.GET("tacher-groups", hand.GetTeacherGroups)
		group.GET("students", hand.GetGroupStudents)
	}

	topic := router.Group("/api/topics")
	group.Use(middleware.Check)
	group.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	topic.Use(handler.CORSMiddleware())
	{
		topic.POST("/create", hand.CreateTopic)
		topic.PUT("/update", hand.UpdateTopic)
		topic.DELETE("/delete", hand.DeleteTopic)
		topic.GET("/getAll", hand.GetAllTopics)
	}

	return router
}
