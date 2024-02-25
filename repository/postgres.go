package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/config"
	"gin/types"
	"log"
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
	sql := `INSERT INTO questions (id, title, description, date, level) VALUES ($1, $2, $3, $4, $5)`
	_, err := p.conn.Exec(sql, question.ID, question.Title, question.Description, question.Date, question.Level)
	return err
}

func (p *Postgres) ReadQuestion(ctx context.Context) (*types.Question, error) {
	sql := `SELECT id, title, description, date, level FROM questions WHERE date = $1 LIMIT 1`

	var question types.Question

	err := p.conn.QueryRow(sql, "2024-02-24").Scan(
		&question.ID,
		&question.Title,
		&question.Description,
		&question.Date,
		&question.Level,
	)
	if err != nil {
		log.Fatal(err)
	}

	return &question, nil
}

func (p *Postgres) UpdateQuestion(ctx context.Context, id *int, dataToChange *types.Question) error {
	return nil
}

func (p *Postgres) DeleteQuestion(ctx context.Context, id *int) error {
	return nil
}

func (p *Postgres) CreateUser(ctx context.Context, user *types.User) {

}

func (p *Postgres) ReadUser(ctx context.Context) {

}

func (p *Postgres) UpdateUser(ctx context.Context) {

}

func (p *Postgres) DeleteUser(ctx context.Context) {

}
