package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type Login struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
