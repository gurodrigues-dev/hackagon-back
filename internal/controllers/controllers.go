package controllers

import (
	"fmt"
	"gin/config"
	"gin/internal/service"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type controller struct {
	service *service.Service
}

func New(s *service.Service) *controller {
	return &controller{
		service: s,
	}
}

type Claims struct {
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		var secretKey []byte

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("nickname", claims.Nickname)
		c.Set("isAuthenticated", true)
		c.Next()

	}
}

func (ct *controller) Start() {

	config := config.Get()

	router := gin.Default()

	api := router.Group("/api/v1")
	api.GET("/ping", ct.Ping)
	api.POST("/question")
	api.GET("/question")
	api.PUT("/question")
	api.DELETE("/question")

	router.Run(fmt.Sprintf(":%d", config.Server.Port))

}

func (ct *controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (ct *controller) CreateQuestion() {

}

// func (ct *controller) ReadQuestion() {

// }

// func (ct *controller) UpdateQuestion() {

// }

// func (ct *controller) DeleteQuestion() {

// }
