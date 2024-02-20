package service

import "gin/repository"

type Service struct {
	repository repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) CreateQuestion() error {
	return nil
}

func (s *Service) ReadQuestion() error {
	return nil
}

func (s *Service) UpdateQuestion() error {
	return nil
}

func (s *Service) DeleteQuestion() error {
	return nil
}

func (s *Service) CreateTest() error {
	return nil
}

func (s *Service) ReadTest() error {
	return nil
}

func (s *Service) UpdateTest() error {
	return nil
}

func (s *Service) DeleteTest() error {
	return nil
}

func (s *Service) CreateUserResponse() error {
	return nil
}

func (s *Service) ReadUserResponse() error {
	return nil
}

func (s *Service) UpdateUserResponse() error {
	return nil
}

func (s *Service) DeleteUserResponse() error {
	return nil
}
