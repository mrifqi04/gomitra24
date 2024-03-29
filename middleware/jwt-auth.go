package middleware

import (
	"log"
	"mitra24/helper"
	"mitra24/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("failed process request", "no token found", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("claim[user_id] : ", claims["user_id"])
			log.Println("claim[issuer] : ", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
