package routes

import (
	"hospital-management/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/doctor/", controllers.CreateDoctor)
	router.GET("/doctor/:id", controllers.GetDoctor)
	router.PATCH("/doctor/:id", controllers.UpdateDoctorContact)

	router.POST("/patient/", controllers.CreatePatient)
	router.GET("/patient/:id", controllers.GetPatient)
	router.PATCH("/patient/:id", controllers.UpdatePatientDetails)
	router.GET("/patientbyDoctorId/:doctorId", controllers.GetPatientsByDoctorID)
}
