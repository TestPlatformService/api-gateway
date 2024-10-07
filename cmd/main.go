package main

import (
	"api/api"
	"api/api/handler"
	"api/casbin"
	"api/config"
	"api/genproto/group"
	"api/genproto/notification"
	"api/genproto/question"
	"api/genproto/subject"
	"api/genproto/task"
	"api/genproto/topic"
	"api/genproto/user"
	"api/logs"
	"log"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf := config.Load()
	hand := NewHandler()
	router := api.Router(hand)
	log.Printf("server is running...")
	log.Fatal(router.Run(conf.API_ROUTER))
}

func NewHandler() *handler.Handler {
	conf := config.Load()
	connUser, err := grpc.NewClient(conf.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	connQuestion, err := grpc.NewClient(conf.QUESTION_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	User := user.NewUsersClient(connUser)
	Notification := notification.NewNotificationsClient(connUser)
	Group := group.NewGroupServiceClient(connUser)
	Topic := topic.NewTopicServiceClient(connQuestion)
	Subject := subject.NewSubjectServiceClient(connQuestion)
	Question := question.NewQuestionServiceClient(connQuestion)
	QuestionOutput := question.NewOutputServiceClient(connQuestion)
	QuestionInput := question.NewInputServiceClient(connQuestion)
	QuestionTest := question.NewTestCaseServiceClient(connQuestion)
	Task := task.NewTaskServiceClient(connQuestion)

	logs := logs.NewLogger()
	en, err := casbin.CasbinEnforcer(logs)
	if err != nil {
		log.Fatal("error in creating casbin enforcer", err)
	}
	return &handler.Handler{
		User:           User,
		Notification:   Notification,
		Group:          Group,
		Log:            logs,
		Enforcer:       en,
		Question:       Question,
		QuestionOutput: QuestionOutput,
		QuestionInput:  QuestionInput,
		TestCase:       QuestionTest,
		Topic:          Topic,
		Subject:        Subject,
		Task:           Task,
		Connections:    make(map[string]*websocket.Conn),
	}
}
