package main

import (
	"api/api"
	"api/api/handler"
	"api/casbin"
	"api/config"
	"api/logs"
	"api/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("API Gateway started successfully!")
	logger := logs.NewLogger()
	logger.Info("API Gateway started successfully!")

	enforcer, err := casbin.CasbinEnforcer(logger)
	if err != nil {
        log.Println("Error initializing casbin enforcer", "error", err.Error())
		logger.Error("Error initializing enforcer", "error", err.Error())
		return
    }

	config := config.Load()
	serviceManager, err := service.NewServiceManager()
	if err != nil {
		log.Println("Error initializing service manager", "error", err.Error())
		logger.Error("Error initializing service manager", "error", err.Error())
		return
	}


	handler := handler.NewHandler(serviceManager.UserService(), logger, enforcer)
	controller := api.NewController(gin.Default())
	controller.SetupRoutes(*handler, logger)
	controller.StartServer(config)

}