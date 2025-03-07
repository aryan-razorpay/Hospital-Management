package controllers

import (
	"hospital-management/methods"
	"hospital-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient.ID = uuid.New().String()[:5]
	methods.CreatePatient(&patient)
	c.JSON(http.StatusOK, patient)
}

func GetPatient(c *gin.Context) {
	id := c.Param("id")
	patient, err := methods.GetPatientByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func UpdatePatientDetails(c *gin.Context) {
	id := c.Param("id")
	var body models.Patient
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	methods.UpdatePatientDetails(id, body.ContactNo, body.Address, body.DoctorID)
	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func GetPatientsByDoctorID(c *gin.Context) {
	doctorID := c.Param("doctorId")
	patients, err := methods.GetPatientsByDoctorID(doctorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	c.JSON(http.StatusOK, patients)
}
