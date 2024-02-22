package service

import (
	"context"
	"gin/repository"
	"gin/types"
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

	return nil
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
