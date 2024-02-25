package service

import (
	"context"
	"gin/repository"
	"gin/types"

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

func (s *Service) DeleteQuestion() error {
	return nil
}

func (s *Service) CreateUser(ctx context.Context, user *types.User) {

}

func (s *Service) ReadUser(ctx context.Context) {

}

func (s *Service) UpdateUser(ctx context.Context) {

}

func (s *Service) DeleteUser(ctx context.Context) {

}
