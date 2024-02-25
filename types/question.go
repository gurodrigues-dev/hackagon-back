package types

import (
	"github.com/google/uuid"
)

type Question struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        string    `json:"date"`
	Level       string    `json:"level"`
}

type QuestionCreateRequest struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Level       string    `json:"level"`
	Date        string    `json:"date"`
}

type Param struct {
}
