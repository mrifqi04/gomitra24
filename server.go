package main

import (
	"mitra24/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1/")

	v1.GET("/", controller.Ping)

	v1.POST("login", controller.Login)
	v1.POST("register", controller.Register)

	r.Run(":3000")
}
