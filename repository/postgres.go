package repository

import (
	"database/sql"
	"fmt"
	"gin/config"
)

type Postgres struct {
	conn *sql.DB
}

func NewPostgres() (*Postgres, error) {

	conf := config.Get()

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.Name),
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	repo := &Postgres{
		conn: db,
	}

	return repo, nil
}

func (p *Postgres) CreateQuestion() error {
	return nil
}

func (p *Postgres) ReadQuestion() error {
	return nil
}

func (p *Postgres) UpdateQuestion() error {
	return nil
}

func (p *Postgres) DeleteQuestion() error {
	return nil
}

func (p *Postgres) CreateTest() error {
	return nil
}

func (p *Postgres) ReadTest() error {
	return nil
}

func (p *Postgres) UpdateTest() error {
	return nil
}

func (p *Postgres) DeleteTest() error {
	return nil
}

func (p *Postgres) CreateUserResponse() error {
	return nil
}

func (p *Postgres) ReadUserResponse() error {
	return nil
}

func (p *Postgres) UpdateUserResponse() error {
	return nil
}

func (p Postgres) DeleteUserResponse() error {
	return nil
}
