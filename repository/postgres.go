package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/config"
	"gin/types"
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

func (p *Postgres) CreateQuestion(ctx context.Context, question *types.Question) error {
	return nil
}

func (p *Postgres) ReadQuestion(ctx context.Context, id *int) error {
	return nil
}

func (p *Postgres) UpdateQuestion(ctx context.Context, id *int, dataToChange *types.Question) error {
	return nil
}

func (p *Postgres) DeleteQuestion(ctx context.Context, id *int) error {
	return nil
}
