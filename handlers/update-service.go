package handlers

import (
	"net/http"

	"github.com/cassiusbessa/micro-service-update-service/entities"
	"github.com/cassiusbessa/micro-service-update-service/logs"
	"github.com/cassiusbessa/micro-service-update-service/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UpdateService(c *gin.Context) {
	var service entities.Service
	db := c.Param("company")
	logrus.Warnf("Updating Service on %s", db)
	if err := c.BindJSON(&service); err != nil {
		logrus.Errorf("Error decoding Service %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := service.Validate(); err != nil {
		logrus.Errorf("Error validating Service on %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result, err := repositories.UpdateService(db, service.Id.Hex(), service)
	if err != nil {
		logrus.Errorf("Error Updating Service on %v: %v", db, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !result {
		logrus.Errorf("Service not found on %v, with id: %v", db, service.Id.String())
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	defer logs.Elapsed("Updated Schedule")()
	c.JSON(http.StatusOK, gin.H{"message": "Service updated successfully"})

}
