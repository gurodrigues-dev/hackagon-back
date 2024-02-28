package controllers

import (
	"fmt"
	"gin/config"
	"gin/internal/service"
	"gin/types"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (ct *controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (ct *controller) AuthUser(c *gin.Context) {

	var input types.User

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error parsing body content: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ct.service.VerifyLogin(c, &input)

	if err != nil {
		log.Printf("wrong nickname or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	jwt, err := ct.service.CreateTokenJwt(c, &input)

	if err != nil {
		log.Printf("error while creating jwt: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwt})

}

func (ct *controller) createQuestion(c *gin.Context) {

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

	question := input.ToQuestion()

	err := ct.service.CreateQuestion(c, &question)

	if err != nil {
		log.Printf("error creating question: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, question)

}

func (ct *controller) getQuestion(c *gin.Context) {

	var input types.Answer

	err := ct.service.ParserJwt(c)

	if err != nil {
		log.Printf("error to parser jwt: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	nicknameInterface, _ := c.Get("nickname")

	nickname := fmt.Sprint(nicknameInterface)

	input.Nickname = nickname

	question, err := ct.service.ReadQuestion(c)

	if err != nil {
		log.Printf("error searching question %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	verify, err := ct.service.VerifyAnswer(c, question, &nickname)

	if err != nil {
		log.Printf("error verify answer %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if verify.Status == "SUCCESS" || verify.Status == "RUNNING" {
		log.Println("It was not possible to answer the question, this user has already answered this question")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":                  "It was not possible to answer the question, this user has already answered this question",
			"currentStatus":            verify.Status,
			"currentResponseCreatedAt": verify.CreatedAt,
		})
		return
	}

	c.JSON(http.StatusOK, question)

}

func (ct *controller) updateQuestion(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "update done successfully"})

}

func (ct *controller) deleteQuestion(c *gin.Context) {

	input := c.Param("id")

	id, err := uuid.Parse(input)

	if err != nil {
		log.Printf("error parsing path ID: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = ct.service.DeleteQuestion(c, id)

	if err != nil {
		log.Printf("error while deleting question: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "question deleted done successfully"})
}

func (ct *controller) createUser(c *gin.Context) {

	var input types.User

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error parsing body content: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err := ct.service.CreateUser(c, &input)

	if err != nil {
		log.Printf("error creating question: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create done successfully"})

}

func (ct *controller) getUser(c *gin.Context) {

	err := ct.service.ParserJwt(c)

	if err != nil {
		log.Printf("error to parser jwt: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	input := c.Param("id")

	id, err := strconv.Atoi(input)
	if err != nil {
		log.Printf("error to parse id: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid id"})
		return
	}

	user, err := ct.service.ReadUser(c, &id)

	if err != nil {
		log.Printf("error to find user: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "don't find user"})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (ct *controller) updateUser(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "user update done successfully"})

}

func (ct *controller) deleteUser(c *gin.Context) {

	err := ct.service.ParserJwt(c)

	if err != nil {
		log.Printf("error to parser jwt: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	nicknameInterface, _ := c.Get("nickname")

	nickname := fmt.Sprint(nicknameInterface)

	err = ct.service.DeleteUser(c, &nickname)
	if err != nil {
		log.Printf("error while deleting account: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user delete done successfully"})

}

func (ct *controller) CreateAnswer(c *gin.Context) {

	var input types.Answer

	err := ct.service.ParserJwt(c)

	if err != nil {
		log.Printf("error to parser jwt: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	nicknameInterface, _ := c.Get("nickname")

	nickname := fmt.Sprint(nicknameInterface)

	input.Nickname = nickname

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error parsing body content: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	err = ct.service.CreateAnswer(c, &input)

	if err != nil {
		log.Printf("error while creating answer: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = ct.service.IncreaseScore(c, &input.Nickname, &input.Points)

	if err != nil {
		log.Printf("error while increasing score: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "answer created"})

}

func (ct *controller) DeleteAnswer(c *gin.Context) {

	input := c.Param("id")

	id, err := uuid.Parse(input)
	if err != nil {
		log.Printf("error parsing path ID: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = ct.service.DeleteAnswer(c, id)
	if err != nil {
		log.Printf("error while deleted answer: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "answer deleted!"})

}

func (ct *controller) GetRank(c *gin.Context) {

	err := ct.service.ParserJwt(c)

	if err != nil {
		log.Printf("error to parser jwt: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	nicknameInterface, _ := c.Get("nickname")

	nickname := fmt.Sprint(nicknameInterface)

	rank, err := ct.service.GetRank(c, &nickname)

	if err != nil {
		log.Printf("error while mount rank: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, rank)

}

func (ct *controller) Start() {

	conf := config.Get()

	router := gin.Default()

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

	router.Run(fmt.Sprintf(":%d", conf.Server.Port))

}
