package repository

import "database/sql"

type Postgres struct {
	conn *sql.DB
}

func NewPostgres() (*Postgres, error) {

	var db *sql.DB

	repo := &Postgres{
		conn: db,
	}

	return repo, nil
}
