package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/config"
	"gin/types"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
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

	err = repo.migrate(conf.Database.Schema)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (p *Postgres) migrate(filepath string) error {

	schema, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = p.conn.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) CreateQuestion(ctx context.Context, question *types.Question) error {

	params1 := strings.Join(question.Inputs.Test1.Params, ",")
	params2 := strings.Join(question.Inputs.Test2.Params, ",")
	params3 := strings.Join(question.Inputs.Test3.Params, ",")

	sqlQuery := `INSERT INTO questions (id, title, description, date, level, params1, response1, params2, response2, params3, response3) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := p.conn.Exec(sqlQuery, question.ID, question.Title, question.Description, question.Date, question.Level, params1, question.Inputs.Test1.Response, params2, question.Inputs.Test2.Response, params3, question.Inputs.Test3.Response)
	return err
}

func (p *Postgres) ReadQuestion(ctx context.Context) (*types.Question, error) {
	sqlQuery := `SELECT id, title, description, date, level, params1, response1, params2, response2, params3, response3 FROM questions WHERE date = $1 LIMIT 1`

	var (
		question   types.Question
		params1Str string
		params2Str string
		params3Str string
	)

	err := p.conn.QueryRow(sqlQuery, time.Now().Format("2006-01-02")).Scan(
		&question.ID,
		&question.Title,
		&question.Description,
		&question.Date,
		&question.Level,
		&params1Str,
		&question.Inputs.Test1.Response,
		&params2Str,
		&question.Inputs.Test2.Response,
		&params3Str,
		&question.Inputs.Test3.Response,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	question.Inputs.Test1.Params = strings.Split(params1Str, ",")
	question.Inputs.Test2.Params = strings.Split(params2Str, ",")
	question.Inputs.Test3.Params = strings.Split(params3Str, ",")

	return &question, nil
}

func (p *Postgres) UpdateQuestion(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (p *Postgres) DeleteQuestion(ctx context.Context, id uuid.UUID) error {
	tx, err := p.conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.Exec("DELETE FROM questions WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) CreateUser(ctx context.Context, user *types.User) error {
	sqlQuery := `INSERT INTO users (nickname, email, password) VALUES ($1, $2, $3)`
	_, err := p.conn.Exec(sqlQuery, user.Nickname, user.Email, user.Password)
	return err
}

func (p *Postgres) ReadUser(ctx context.Context, id *int) (*types.User, error) {
	sqlQuery := `SELECT id, nickname, email FROM users WHERE id = $1 LIMIT 1`

	var user types.User

	err := p.conn.QueryRow(sqlQuery, id).Scan(
		&user.ID,
		&user.Nickname,
		&user.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}

func (p *Postgres) UpdateUser(ctx context.Context, id *int) error {
	return nil
}

func (p *Postgres) DeleteUser(ctx context.Context, nickname *string) error {
	tx, err := p.conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.Exec("DELETE FROM users WHERE nickname = $1", nickname)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) VerifyLogin(ctx context.Context, user *types.User) (*types.User, error) {
	sqlQuery := `SELECT id, nickname, email, password FROM users WHERE nickname = $1 LIMIT 1`
	var userData types.User
	err := p.conn.QueryRow(sqlQuery, user.Nickname).Scan(
		&userData.ID,
		&userData.Nickname,
		&userData.Email,
		&userData.Password,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	match := userData.Password == user.Password
	if !match {
		return nil, fmt.Errorf("email or password wrong")
	}
	userData.Password = ""
	return &userData, nil

}

func (p *Postgres) CreateAnswer(ctx context.Context, answer *types.Answer) error {
	sqlQuery := `INSERT INTO answers (id, nickname, questionid, status, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := p.conn.Exec(sqlQuery, answer.ID, answer.Nickname, answer.QuestionID, answer.Status, answer.CreatedAt)
	return err
}

func (p *Postgres) DeleteAnswer(ctx context.Context, id uuid.UUID) error {
	tx, err := p.conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.Exec("DELETE FROM answers WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) VerifyAnswer(ctx context.Context, question *types.Question, nickname *string) (*types.Answer, error) {
	sqlQuery := `SELECT nickname, status, created_at FROM answers WHERE questionid = $1 AND nickname = $2 ORDER BY created_at DESC LIMIT 1`

	var answerResponse types.Answer

	err := p.conn.QueryRow(sqlQuery, question.ID, nickname).Scan(
		&answerResponse.Nickname,
		&answerResponse.Status,
		&answerResponse.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return &answerResponse, nil
		}

		log.Fatal(err)
		return nil, err
	}

	return &answerResponse, nil

}

func (p *Postgres) IncreaseScore(ctx context.Context, nickname *string, points *int) error {
	sqlQuery := `UPDATE users SET points = points + $1 WHERE nickname = $2`

	_, err := p.conn.Exec(sqlQuery, points, nickname)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetRank(ctx context.Context, nickname *string) ([]types.Rank, error) {
	sqlQuery := `SELECT nickname, points, RANK() OVER (ORDER BY points DESC) AS rank FROM users`

	rows, err := p.conn.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ranks []types.Rank

	for rows.Next() {
		var rank types.Rank
		err := rows.Scan(&rank.Nickname, &rank.Points, &rank.Position)
		if err != nil {
			return nil, err
		}
		ranks = append(ranks, rank)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ranks, nil
}

func (p *Postgres) NewPassword(ctx context.Context, user *types.User) error {
	sqlQuery := `UPDATE users SET password = $1 WHERE email = $2`

	_, err := p.conn.Exec(sqlQuery, user.Password, user.Email)
	return err
}

func (p *Postgres) VerifyEmailExists(ctx context.Context, email *string) (bool, error) {
	sqlQuery := `SELECT email FROM users WHERE email = $1`

	var emailDatabase string

	err := p.conn.QueryRow(sqlQuery, email).Scan(&emailDatabase)

	fmt.Println(err)

	if err != nil {
		return false, err
	}

	return true, nil

}
