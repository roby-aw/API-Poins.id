package store

import "github.com/go-playground/validator/v10"

type Repository interface {
	SignStore(auth *AuthStore) (*ResponseLoginStore, error)
	InputPoin(input *InputPoin) (*int, error)
}

type Service interface {
	LoginStore(auth *AuthStore) (*ResponseLoginStore, error)
	InputPoin(input *InputPoin) (*int, error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) LoginStore(auth *AuthStore) (*ResponseLoginStore, error) {
	return s.repository.SignStore(auth)
}

func (s *service) InputPoin(input *InputPoin) (*int, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}
	return s.repository.InputPoin(input)
}
