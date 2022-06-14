package admin

import (
	"api-redeem-point/business/customermitra"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	InsertAdmin(admin *customermitra.Admin) (*customermitra.Admin, error)
	RemoveAdmin(id int) error
	RenewAdmin(id int, admin *Admin) (*Admin, error)
	LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error)
}

type Service interface {
	CreateAdmin(admin *customermitra.Admin) (*customermitra.Admin, error)
	DeleteAdmin(id int) error
	UpdateAdmin(id int, admin *Admin) (*Admin, error)
	LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error)
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

func (s *service) CreateAdmin(admin *customermitra.Admin) (*customermitra.Admin, error) {
	err := s.validate.Struct(admin)
	if err != nil {
		return nil, err
	}
	admin, err = s.repository.InsertAdmin(admin)
	return admin, err
}

func (s *service) DeleteAdmin(id int) error {
	return s.repository.RemoveAdmin(id)
}

func (s *service) UpdateAdmin(id int, admin *Admin) (*Admin, error) {
	return s.repository.RenewAdmin(id, admin)
}

func (s *service) LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error) {
	tokens, err := s.repository.LoginAdmin(Auth)
	return tokens, err
}
