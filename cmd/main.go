package main

import (
	"hospital-management/config"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
