package utils

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, data any, message string, status int) {
	c.JSON(status, gin.H{
		"code":    status,
		"message": message,
		"data":    data,
	})
}
