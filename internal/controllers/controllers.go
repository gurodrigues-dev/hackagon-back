package controllers

import (
	"fmt"
	"gin/config"
	"gin/internal/service"
	"gin/types"
	"log"
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

func (ct *controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (ct *controller) CreateQuestion(c *gin.Context) {

	var input types.QuestionCreateRequest

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error parsing body content: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	if err := input.ValidateCreate(); err != nil {
		log.Printf("invalid input: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question := types.Question{
		Title:       input.Title,
		Description: input.Description,
		Level:       input.Level,
		Date:        input.Date,
	}

	err := ct.service.CreateQuestion(c, &question)

	if err != nil {
		log.Printf("error creating question: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, question)

}

func (ct *controller) GetQuestion(c *gin.Context) {

	question, err := ct.service.ReadQuestion(c)

	if err != nil {
		log.Printf("error searching question %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, question)

}

func (ct *controller) UpdateQuestion(c *gin.Context) {

}

func (ct *controller) DeleteQuestion(c *gin.Context) {

}

func (ct *controller) Start() {

	config := config.Get()

	router := gin.Default()

	api := router.Group("/api/v1")
	api.GET("/ping", ct.Ping)
	api.POST("/question", ct.CreateQuestion)
	api.GET("/question", ct.GetQuestion)
	api.PUT("/question")
	api.DELETE("/question")

	router.Run(fmt.Sprintf(":%d", config.Server.Port))

}
