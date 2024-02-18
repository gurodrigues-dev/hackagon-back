package repository

import (
	"database/sql"
	"fmt"
	"gin/config"
	"gin/models"
	"log"
	"reflect"

	_ "github.com/lib/pq"
)

func SaveUser(user *models.User) error {
	err := config.LoadEnvironmentVariables()
	if err != nil {
		return err
	}

	var (
		userdb   = config.GetDatabaseUser()
		password = config.GetDatabasePassword()
		dbname   = config.GetDatabaseName()
		host     = config.GetDatabaseHost()
		port     = config.GetDatabasePort()
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, userdb, password, dbname)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (name, email, password, nickname) VALUES ($1, $2, $3, $4)",
		&user.Name, &user.Email, &user.Password, &user.Nickname)
	if err != nil {
		return err
	}

	return nil
}

func FindUserByNick(nick *interface{}) (*models.User, error) {
	err := config.LoadEnvironmentVariables()
	if err != nil {
		return nil, err
	}

	var (
		userdb   = config.GetDatabaseUser()
		password = config.GetDatabasePassword()
		dbname   = config.GetDatabaseName()
		host     = config.GetDatabaseHost()
		port     = config.GetDatabasePort()
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, userdb, password, dbname)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT name, email, nickname
		FROM users
		WHERE nickname = $1
	`, *nick)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User

	found := false

	for rows.Next() {
		found = true
		err := rows.Scan(&user.Name, &user.Email, &user.Nickname)

		if err != nil {
			log.Fatal(err)
		}
	}

	if !found {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil

}

type nickInterface struct {
	Nickname interface{}
}

func UpdateUser(user *models.User) error {

	newUser := nickInterface{
		Nickname: user.Nickname,
	}

	err := config.LoadEnvironmentVariables()
	if err != nil {
		return err
	}

	var (
		userdb   = config.GetDatabaseUser()
		password = config.GetDatabasePassword()
		dbname   = config.GetDatabaseName()
		host     = config.GetDatabaseHost()
		port     = config.GetDatabasePort()
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, userdb, password, dbname)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	dataOldUser, err := FindUserByNick(&newUser.Nickname)

	if err != nil {
		return err
	}

	userUpdated := func(old, new *models.User) *models.User {

		updatedUser := old

		t := reflect.TypeOf(new)

		for i := 0; i < t.NumField(); i++ {
			fieldName := t.Field(i).Name
			fieldValue := reflect.ValueOf(new).FieldByName(fieldName)

			if reflect.ValueOf(old).FieldByName(fieldName).Interface() != fieldValue.Interface() {
				reflect.ValueOf(&updatedUser).Elem().FieldByName(fieldName).Set(fieldValue)
			}
		}

		return updatedUser

	}(dataOldUser, user)

	_, err = db.Exec(`UPDATE users SET email=$1, password=$2, nickname=$3 WHERE nickname=$4`, userUpdated.Email, userUpdated.Password, userUpdated.Nickname, dataOldUser.Nickname)

	if err != nil {
		return err
	}

	return nil

}

func DeleteUser(nick *interface{}) error {

	err := config.LoadEnvironmentVariables()
	if err != nil {
		return err
	}

	var (
		userdb   = config.GetDatabaseUser()
		password = config.GetDatabasePassword()
		dbname   = config.GetDatabaseName()
		host     = config.GetDatabaseHost()
		port     = config.GetDatabasePort()
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, userdb, password, dbname)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE nickname = $1",
		*nick)

	if err != nil {
		log.Fatal("")
		return err
	}

	return nil
}

func VerifyLoginByNickname(login *models.Login) error {

	err := config.LoadEnvironmentVariables()
	if err != nil {
		return err
	}

	var (
		userdb   = config.GetDatabaseUser()
		password = config.GetDatabasePassword()
		dbname   = config.GetDatabaseName()
		host     = config.GetDatabaseHost()
		port     = config.GetDatabasePort()
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, userdb, password, dbname)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	defer db.Close()

	var passwordUser string

	err = db.QueryRow("SELECT password FROM users WHERE nickname = $1", &login.Nickname).Scan(&passwordUser)
	if err != nil {
		return err
	}

	passwordMatch := passwordUser == login.Password

	if !passwordMatch {
		return fmt.Errorf("password wrong")
	}

	return nil

}
