package controllers

import (
	"hospital-management/config"
	"hospital-management/methods"
	"hospital-management/models"
	"hospital-management/mutex"
	"net/http"

	"hospital-management/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
var existingPatient models.Patient
var doctor models.Doctor
func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mutex.Mu.Lock()
	defer mutex.Mu.Unlock()
	
	if len(patient.ContactNo) < 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contact number must be at least 10 characters"})
		return
	}
	if(!utils.IsValidName(patient.Name)){
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please Enter a Valid Name"})
		return
	}
	if err := config.DB.Where("id = ?", patient.ID).First(&existingPatient).Error; err == nil {
		utils.ErrorResponse(c, http.StatusConflict, "Patient with the same ID already exists")
		return
	}
	if err := config.DB.Where("contact_no = ?", patient.ContactNo).First(&existingPatient).Error; err == nil {
        utils.ErrorResponse(c, http.StatusConflict, "Patient with the same contact number already exists")
        return
    }
	if err := config.DB.Where("id = ?", patient.DoctorID).First(&doctor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Please enter a valid doctor ID"})
		return
	}
	patient.ID = uuid.New().String()[:5]
	methods.CreatePatient(&patient)
	c.JSON(http.StatusOK, patient)
}

func GetPatient(c *gin.Context) {
	id := c.Param("id")
	patient, err := methods.GetPatientByID(id)
	if len(id) != 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be exactly 5 characters"})
		return
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func UpdatePatientDetails(c *gin.Context) {
	id := c.Param("id")
	if err := utils.ValidateIDLength(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body models.Patient
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := methods.UpdatePatientDetails(id, body.ContactNo, body.Address, body.DoctorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient details"})
		return
	}

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
