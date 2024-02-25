package types

import (
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) HashPassword() string {
	hasher := sha256.New()
	hasher.Write([]byte(u.Password))
	return hex.EncodeToString(hasher.Sum(nil))
}
