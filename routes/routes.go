package routes

import (
	"gin/config"
	"gin/controllers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Nickname string `json:"nickname"`
	jwt.StandardClaims
}

func HandleRequests() {

	config.LoadEnvironmentVariables()

	var secretKey = []byte(config.GetSecretKeyApi())

	r := gin.Default()

	authMiddleware := func(c *gin.Context) {
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

	api := r.Group("/api/v1")
	api.GET("/ping", controllers.Ping)
	api.POST("/user", controllers.CreateUser)
	api.GET("/user", authMiddleware, controllers.GetUser)
	// api.PUT("/user", authMiddleware, controllers.UpdateUser)
	api.DELETE("/user", authMiddleware, controllers.DeleteUser)
	api.POST("/login", controllers.Login)

	r.Run(":9692")
}
