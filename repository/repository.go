package repository

import (
	_ "github.com/lib/pq"
)

type Repository interface {
	CreateQuestion() error
	ReadQuestion() error
	UpdateQuestion() error
	DeleteQuestion() error
	CreateTest() error
	ReadTest() error
	UpdateTest() error
	DeleteTest() error
	CreateUserResponse() error
	ReadUserResponse() error
	UpdateUserResponse() error
	DeleteUserResponse() error
}
