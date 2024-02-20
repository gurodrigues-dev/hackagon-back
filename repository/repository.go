package repository

import (
	_ "github.com/lib/pq"
)

type Repository interface {
	CreateQuestion()
	ReadQuestion()
	UpdateQuestion()
	DeleteQuestion()
	CreateTest()
	ReadTest()
	UpdateTest()
	DeleteTest()
	CreateUserResponse()
	ReadUserResponse()
	UpdateUserResponse()
	DeleteUserResponse()
}
