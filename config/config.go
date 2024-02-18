package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func GetDatabaseHost() string {
	return os.Getenv("host")
}

func GetDatabaseName() string {
	return os.Getenv("dbname")
}

func GetDatabaseUser() string {
	return os.Getenv("user")
}

func GetDatabasePassword() string {
	return os.Getenv("password")
}

func GetDatabasePort() string {
	return os.Getenv("port")
}

func GetSecretKeyApi() string {
	return os.Getenv("secret")
}
