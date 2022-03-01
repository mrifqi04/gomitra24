package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login request",
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Register Request",
	})
}
