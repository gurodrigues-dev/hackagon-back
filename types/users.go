package types

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Points   int    `json:"points"`
}

type Answer struct {
	ID         uuid.UUID `json:"id"`
	Nickname   string    `json:"nickname"`
	QuestionID uuid.UUID `json:"questionid"`
	Status     string    `json:"status"`
	CreatedAt  string    `json:"time"`
	Points     int       `json:"points"`
}

type Rank struct {
	Nickname string `json:"nickname"`
	Points   int    `json:"points"`
	Position int    `json:"position"`
}

func (u *User) HashPassword() string {
	hasher := sha256.New()
	hasher.Write([]byte(u.Password))
	return hex.EncodeToString(hasher.Sum(nil))
}
