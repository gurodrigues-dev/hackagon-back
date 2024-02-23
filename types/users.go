package types

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Nickname string    `json:"nickname"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
