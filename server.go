package main

import (
	"mitra24/config"
	"mitra24/controller"
	"mitra24/repository"
	"mitra24/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1/")

	v1.GET("/", controller.Ping)

	v1.POST("register", authController.Register)
	v1.POST("login", authController.Login)

	r.Run(":3000")
}
