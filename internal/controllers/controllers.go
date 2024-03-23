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

func (ct *controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (ct *controller) Start() {

	conf := config.Get()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		c.Writer.Header().Set("Cross-Origin-Opener-Policy", "same-origin")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	authMiddleware := func(c *gin.Context) {

		secret := []byte(conf.Server.Secret)

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
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

	api := router.Group("/api/v1")
	api.GET("/ping", ct.Ping)
	api.POST("/login", ct.AuthUser)
	api.POST("/question", ct.createQuestion)
	api.GET("/question", authMiddleware, ct.getQuestion)
	api.PATCH("/question/:id", ct.updateQuestion)
	api.DELETE("/question/:id", ct.deleteQuestion)
	api.POST("/user", ct.createUser)
	api.GET("/user/:id", authMiddleware, ct.getUser)
	api.PATCH("/user", authMiddleware, ct.updateUser)
	api.DELETE("/user", authMiddleware, ct.deleteUser)
	api.POST("/answer", authMiddleware, ct.CreateAnswer)
	api.DELETE("/answer/:id", authMiddleware, ct.DeleteAnswer)
	api.GET("/rank", authMiddleware, ct.GetRank)
	api.POST("/password", ct.sendEmailToRecovery)
	api.GET("/password/:token", ct.verifyToken)
	api.PATCH("/password", ct.updatePassword)

	router.Run(fmt.Sprintf(":%d", conf.Server.Port))

}
