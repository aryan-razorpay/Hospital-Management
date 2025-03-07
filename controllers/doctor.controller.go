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
	id := c.Param("id")
	doctor, err := methods.GetDoctorByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}
	c.JSON(http.StatusOK, doctor)
}

func UpdateDoctorContact(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		ContactNo string `json:"contact_no"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	methods.UpdateDoctorContact(id, body.ContactNo)
	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}
