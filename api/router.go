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
	{
		topic.POST("/create", hand.CreateTopic)
		topic.PUT("/update", hand.UpdateTopic)
		topic.DELETE("/delete", hand.DeleteTopic)
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
		question.PUT("/update", hand.UpdateQuestion)
		question.DELETE("/delete", hand.DeleteQuestion)
		question.GET("/getAll", hand.GetAllQuestions)
		question.POST("/upload-image", hand.UploadImageToQuestion)
		question.DELETE("/delete-image", hand.DeleteImageFromQuestion)
	}

	questionOutput := router.Group("/api/question-outputs")
	questionOutput.Use(middleware.Check)
	questionOutput.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		questionOutput.POST("/create", hand.CreateQuestionOutput)
		questionOutput.GET("/:id", hand.GetQuestionOutputById)
		questionOutput.PUT("/update", hand.UpdateQuestionOutput)
		questionOutput.DELETE("/delete", hand.DeleteQuestionOutput)
		questionOutput.GET("/:question_id", hand.GetQuestionOutputsByQuestionId)
	}

	questionInput := router.Group("/api/question-inputs")
	questionInput.Use(middleware.Check)
	questionInput.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		questionInput.POST("/create", hand.CreateQuestionInput)
		questionInput.GET("/:id", hand.GetQuestionInputById)
		questionInput.PUT("/update", hand.UpdateQuestionInput)
		questionInput.DELETE("/delete", hand.DeleteQuestionInput)
		questionInput.GET("/:question_id", hand.GetQuestionInputsByQuestionId)
	}

	testCase := router.Group("/api/test-cases")
	testCase.Use(middleware.Check)
	testCase.Use(middleware.CheckPermissionMiddleware(hand.Enforcer))
	{
		testCase.POST("/create", hand.CreateTestCase)
		testCase.GET("/:id", hand.GetTestCaseById)
		testCase.PUT("/update", hand.UpdateTestCase)
		testCase.DELETE("/delete", hand.DeleteTestCase)
		testCase.GET("/:question_id", hand.GetTestCasesByQuestionId)
	}

	return router
}
