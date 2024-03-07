package types

import (
	"github.com/google/uuid"
)

type Params struct {
	Params   []string `json:"params"`
	Response string   `json:"response"`
}

type Inputs struct {
	Test1 Params `json:"test1"`
	Test2 Params `json:"test2"`
	Test3 Params `json:"test3"`
}

type Question struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Date            string    `json:"date"`
	Level           string    `json:"level"`
	Inputs          Inputs    `json:"inputs"`
	UsernameCognito string    `json:"username"`
	PasswordCognito string    `json:"password"`
}

type QuestionCreateRequest struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Level           string    `json:"level"`
	Date            string    `json:"date"`
	Inputs          Inputs    `json:"inputs"`
	UsernameCognito string    `json:"username"`
	PasswordCognito string    `json:"password"`
}

func (qcr *QuestionCreateRequest) ToQuestion() Question {
	return Question{
		ID:              qcr.ID,
		Title:           qcr.Title,
		Description:     qcr.Description,
		Date:            qcr.Date,
		Level:           qcr.Level,
		Inputs:          qcr.Inputs,
		UsernameCognito: qcr.UsernameCognito,
		PasswordCognito: qcr.PasswordCognito,
	}
}
