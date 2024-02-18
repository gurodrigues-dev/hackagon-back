package controllers

import (
	"gin/models"
	"gin/repository"
	"gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func CreateUser(c *gin.Context) {

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "error to parse json on struct",
		})

		return
	}

	user.Password = utils.HashPassword(user.Password)

	err = repository.SaveUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created with success",
	})
}

func GetUser(c *gin.Context) {

	nick, err := utils.FoundUserByJwtNickname(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user, err := repository.FindUserByNick(&nick)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {

	nick, err := utils.FoundUserByJwtNickname(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err = repository.DeleteUser(&nick)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}

func Login(c *gin.Context) {

	var login models.Login

	err := c.ShouldBindJSON(&login)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "error to parse json on struct",
		})

		return
	}

	login.Password = utils.HashPassword(login.Password)

	err = utils.VerifyLogin(&login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "login wrong",
		})

		return
	}

	jwt, err := utils.CreateJWT(&login)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while creating JWToken",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "login accepted",
		"token":   jwt,
	})

}
