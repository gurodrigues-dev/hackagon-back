package utils

import (
	"fmt"
	"gin/config"
	"gin/models"
	"gin/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyLogin(login *models.Login) error {

	err := repository.VerifyLoginByNickname(login)

	if err != nil {
		return err
	}

	return nil

}

func CreateJWT(login *models.Login) (string, error) {

	config.LoadEnvironmentVariables()

	var secretKey = config.GetSecretKeyApi()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nickname": &login.Nickname,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	jwt, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return jwt, nil

}

func FoundUserByJwtNickname(c *gin.Context) (interface{}, error) {

	user, found := c.Get("nickname")

	if !found {
		return "", fmt.Errorf("error to found nickname in jwt")
	}

	return user, nil

}
