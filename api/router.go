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
	router.Use(handler.CORSMiddleware())
	// user
	user := router.Group("/api/user")
	user.Use(middleware.Check)
	user.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
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
	{
		group.POST("/create", hand.CreateGroup)
		group.PUT("/update", hand.UpdateGroup)
		group.DELETE("/delete", hand.DeleteGroup)
		group.GET("/getById/:group_id", hand.GetGroupById)
		group.GET("/getAll", hand.GetAllGroups)
		group.POST("/add-student", hand.AddStudentToGroup)
		group.DELETE("delete-student", hand.DeleteStudentFromGroup)
		group.POST("add-teacher", hand.AddTeacherToGroup)
		group.DELETE("delete-teacher", hand.DeleteTeacherFromGroup)
		group.GET("student-groups/:hh_id", hand.GetStudentGroups)
		group.GET("teacher-groups/:id", hand.GetTeacherGroups)
		group.GET("students/:group_id", hand.GetGroupStudents)
	}

	topic := router.Group("/api/topics")
	topic.Use(middleware.Check)
	topic.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		topic.POST("/create", hand.CreateTopic)
		topic.PUT("/update", hand.UpdateTopic)
		topic.DELETE("/delete/:topic_id", hand.DeleteTopic)
		topic.GET("/getAll", hand.GetAllTopics)
	}

	subject := router.Group("/api/subjects")
	subject.Use(middleware.Check)
	subject.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		subject.POST("/create", hand.CreateSubject)
		subject.GET("/get/:id", hand.GetSubject)
		subject.GET("/getall", hand.GetAllSubjects)
		subject.PUT("/update/:id", hand.UpdateSubject)
		subject.DELETE("/delete/:id", hand.DeleteSubject)
	}

	question := router.Group("/api/questions")
	question.Use(middleware.Check)
	question.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		question.POST("/create", hand.CreateQuestion)
		question.GET("/:id", hand.GetQuestionById)
		question.PUT("/update/:id", hand.UpdateQuestion)
		question.DELETE("/delete/:id", hand.DeleteQuestion)
		question.GET("/getAll", hand.GetAllQuestions)
		question.POST("/upload-image/:id", hand.UploadImageToQuestion)
		question.DELETE("/delete-image/:id", hand.DeleteImageFromQuestion)
	}

	questionInput := router.Group("/api/question-inputs")
	questionInput.Use(middleware.Check)
	questionInput.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		questionInput.GET("/:id", hand.GetQuestionInputById)
		questionInput.DELETE("/delete/:id", hand.DeleteQuestionInput)
		questionInput.GET("/question/:question_id", hand.GetQuestionInputsByQuestionId)
		questionInput.POST("/create", hand.CreateQuestionInput)
	}

	testCase := router.Group("/api/test-cases")
	testCase.Use(middleware.Check)
	testCase.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		testCase.POST("/create", hand.CreateTestCase)
		testCase.GET("/:id", hand.GetTestCaseById)
		testCase.DELETE("/delete/:id", hand.DeleteTestCase)
		testCase.GET("/question/:question_id", hand.GetTestCasesByQuestionId)
	}

	task := router.Group("/api/task")
	task.Use(middleware.Check)
	task.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		task.POST("/create", hand.CreateTask)
		task.DELETE("/delete", hand.DeleteTask)
		task.GET("get", hand.GetTask)
	}

	check := router.Group("/api/check")
	check.Use(middleware.Check)
	check.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		check.POST("/submit", hand.ProxyChecker)
	}

	return router
}
