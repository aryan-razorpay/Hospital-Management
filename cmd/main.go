package main

import (
	"hospital-management/config"
	"hospital-management/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	routes.RegisterRoutes(r)
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
