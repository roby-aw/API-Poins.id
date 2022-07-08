package mitra

import "github.com/go-playground/validator/v10"

type Repository interface {
	SignStore(store *AuthStore) (*ResponseLoginStore, error)
	InputPoin(input *InputPoin) (*int, error)
}

type Service interface {
	LoginStore(store *AuthStore) (*ResponseLoginStore, error)
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

func (s *service) LoginStore(store *AuthStore) (*ResponseLoginStore, error) {
	return s.repository.SignStore(store)
}

func (s *service) InputPoin(input *InputPoin) (*int, error) {
	err := s.validate.Struct(input)
	if err != nil {
		return nil, err
	}
	return s.repository.InputPoin(input)
}
