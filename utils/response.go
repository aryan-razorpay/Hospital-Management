package utils

import (
	"github.com/gin-gonic/gin"
)


func ErrorResponse(c *gin.Context, statusCode int, errorMsg string) {
	c.JSON(statusCode, gin.H{
		"error": errorMsg,
	})
}
