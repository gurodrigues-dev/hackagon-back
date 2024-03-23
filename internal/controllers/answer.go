package controllers

import (
	"fmt"
	"gin/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
