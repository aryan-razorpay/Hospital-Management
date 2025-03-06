package controllers

import (
	"hospital-management/methods"
	"hospital-management/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor.ID = uuid.New().String()[:5]
	methods.CreateDoctor(&doctor)

	c.JSON(http.StatusOK, doctor)
}

func GetDoctor(c *gin.Context) {
}

func UpdateDoctorContact(c *gin.Context) {
}
