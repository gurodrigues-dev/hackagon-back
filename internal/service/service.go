package service

import (
	"context"
	"fmt"
	"gin/config"
	"gin/repository"
	"gin/types"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
	repository repository.Repository
	cloud      repository.Cloud
}

func New(repo repository.Repository, cloud repository.Cloud) *Service {
	return &Service{
		repository: repo,
		cloud:      cloud,
	}
}

func (s *Service) CreateQuestion(ctx context.Context, question *types.Question) error {

	id, err := uuid.NewV7()

	if err != nil {
		return err
	}

	question.ID = id

	return s.repository.CreateQuestion(ctx, question)

}

func (s *Service) ReadQuestion(ctx context.Context) (*types.Question, error) {

	question, err := s.repository.ReadQuestion(ctx)

	return question, err
}

func (s *Service) UpdateQuestion() error {
	return nil
}

func (s *Service) DeleteQuestion(ctx context.Context, id uuid.UUID) error {
	return s.repository.DeleteQuestion(ctx, id)
}

func (s *Service) CreateUser(ctx context.Context, user *types.User) error {

	user.Password = user.HashPassword()

	return s.repository.CreateUser(ctx, user)
}

func (s *Service) ReadUser(ctx context.Context, id *int) (*types.User, error) {
	return s.repository.ReadUser(ctx, id)
}

func (s *Service) UpdateUser(ctx context.Context, id *int) error {
	return s.repository.UpdateUser(ctx, id)
}

func (s *Service) DeleteUser(ctx context.Context, nickname *string) error {
	return s.repository.DeleteUser(ctx, nickname)
}

func (s *Service) VerifyLogin(ctx context.Context, user *types.User) error {

	user.Password = user.HashPassword()

	return s.repository.VerifyLogin(ctx, user)
}

func (s *Service) CreateTokenJwt(ctx context.Context, user *types.User) (string, error) {

	conf := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nickname": user.Nickname,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	jwt, err := token.SignedString([]byte(conf.Server.Secret))

	if err != nil {
		return "", err
	}

	return jwt, nil

}

func (s *Service) ParserJwt(ctx *gin.Context) error {

	_, found := ctx.Get("nickname")

	if !found {
		return fmt.Errorf("error while veryfing token")
	}

	return nil

}

func (s *Service) CreateAnswer(ctx context.Context, answer *types.Answer) error {

	id, err := uuid.NewV7()

	if err != nil {
		return err
	}

	answer.ID = id

	answer.CreatedAt = time.Now().Format("02-01-2006-15-04-05")

	return s.repository.CreateAnswer(ctx, answer)
}

func (s *Service) DeleteAnswer(ctx context.Context, id uuid.UUID) error {
	return s.repository.DeleteAnswer(ctx, id)
}

func (s *Service) VerifyAnswer(ctx context.Context, question *types.Question, nickname *string) (*types.Answer, error) {
	answerResponse, err := s.repository.VerifyAnswer(ctx, question, nickname)
	return answerResponse, err
}

func (s *Service) IncreaseScore(ctx context.Context, nickname *string, points *int) error {
	return s.repository.IncreaseScore(ctx, nickname, points)
}

func (s *Service) GetRank(ctx context.Context, nickname *string) ([]types.Rank, error) {
	return s.repository.GetRank(ctx, nickname)
}

func (s *Service) SendEmail(ctx context.Context, email *string) error {

	return nil
}
