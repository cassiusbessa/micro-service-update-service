package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cassiusbessa/micro-service-update-service/entities"
	"github.com/cassiusbessa/micro-service-update-service/errors"
	"github.com/cassiusbessa/micro-service-update-service/repositories"
	"github.com/gin-gonic/gin"
)

var _, cancel = repositories.MongoConnection()

func UpdateService(c *gin.Context) {
	var service entities.Service
	if err := c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := service.Validate(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	db := c.Param("company")

	result, err := repositories.UpdateService(db, service.Id.Hex(), service)
	if err != nil {
		repositories.SaveError(db, *errors.NewError(http.StatusInternalServerError, "Error Mongo updating service", "UpdateService", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !result {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service updated successfully"})

	defer cancel()
}

func main() {
	r := gin.Default()
	r.PUT("/services/:company", UpdateService)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
