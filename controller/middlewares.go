package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (app *App) JWTMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing auth token"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid or expired token"})
		return
	}

	c.Next()
}
