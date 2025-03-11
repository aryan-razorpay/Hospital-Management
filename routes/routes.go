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
	router.GET("/fetchPatientByDoctorId/:doctorId", controllers.GetPatientsByDoctorID)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome to the Hospital Management System!")
	})
}
