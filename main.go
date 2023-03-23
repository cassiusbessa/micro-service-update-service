package main

import (
	"github.com/cassiusbessa/micro-service-update-service/handlers"
	"github.com/cassiusbessa/micro-service-update-service/logs"
	"github.com/cassiusbessa/micro-service-update-service/repositories"
	"github.com/sirupsen/logrus"
)

var file = logs.Init()

func main() {
	defer file.Close()
	r := handlers.Router()
	repositories.Repo.Ping()
	r.PUT("/services/:company", handlers.UpdateService)
	r.StaticFile("/logs", "./logs/logs.log")
	if err := r.Run(":8080"); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
