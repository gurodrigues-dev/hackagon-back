package service

import (
	"context"
	"gin/repository"
	"gin/types"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repository repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) CreateQuestion(ctx context.Context, question *types.Question) error {

	id, err := uuid.NewV7()

	if err != nil {
		return err
	}

	today := time.Now().Format("2006-01-02")

	question.ID = id
	question.Date = today

	return s.repository.CreateQuestion(ctx, question)

}

func (s *Service) ReadQuestion(ctx context.Context, id *int) error {
	return nil
}

func (s *Service) UpdateQuestion() error {
	return nil
}

func (s *Service) DeleteQuestion() error {
	return nil
}

func (s *Service) CreateUser(ctx context.Context) {

}

func (s *Service) ReadUser(ctx context.Context) {

}

func (s *Service) UpdateUser(ctx context.Context) {

}

func (s *Service) DeleteUser(ctx context.Context) {

}
