package controllers

import (
	"fmt"
	"gin/types"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Nickname string `json:"nickname"`
	jwt.StandardClaims
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

func (ct *controller) createUser(c *gin.Context) {

	var input types.User

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error parsing body content: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	exists, _ := ct.service.VerifyEmailExists(c, &input.Email)

	if exists {
		log.Println("email already exists!")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err := ct.service.CheckEmail(c, &input.Email)

	if err != nil {
		log.Printf("error while sending email: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = ct.service.CreateUser(c, &input)

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

func (ct *controller) updatePassword(c *gin.Context) {
	var input types.User

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error parsing body content: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	email := c.GetHeader("Email")

	if email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email não fornecido"})
		c.Abort()
		return
	}

	input.Email = email

	err := ct.service.NewPassword(c, &input)

	if err != nil {
		log.Printf("error while updating password in database: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password updated"})

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

func (ct *controller) sendEmailToRecovery(c *gin.Context) {

	var input types.User

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error parsing body content: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body content"})
		return
	}

	exists, err := ct.service.VerifyEmailExists(c, &input.Email)

	if !exists {
		log.Printf("error while searching email: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	token, err := ct.service.GenerateRandomToken()

	if err != nil {
		log.Printf("error at creating token: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	email := types.Email{
		Subject:   "Hackagon - Token to recovery your account",
		Body:      fmt.Sprintf("Hey! How are you? Let me see... You forgot your password. Oh, it's a shame. Your token's here: %s", token),
		Recipient: &input.Email,
	}

	err = ct.service.SendEmailToRecovery(c, &email)

	if err != nil {
		log.Printf("error at creating token: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	err = ct.service.SaveRedis(c, token, input.Email)

	if err != nil {
		log.Printf("error while saving token: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "token send w/ successfully"})

}

func (ct *controller) verifyToken(c *gin.Context) {

	var input types.User

	input.Email = c.GetHeader("email")
	token := c.Param("token")

	err := ct.service.VerifyTokenRedis(c, token, input.Email)

	if err != nil {
		log.Printf("error while verifying token: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "token success verified"})

}
