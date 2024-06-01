package controllers

import (
	"fmt"
	"gin/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

	err := ct.service.VerifyCognitoUser(c, &question)

	fmt.Println(question.UsernameCognito, question.PasswordCognito)

	if err != nil {
		log.Printf("error while verifying cognito: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "cognito user or password wrong"})
		return
	}

	err = ct.service.CreateQuestion(c, &question)

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

		question.AnswerQuestion = true

		c.JSON(http.StatusOK, question)

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
