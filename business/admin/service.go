package admin

import (
	"api-redeem-point/business/customermitra"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	Dashboard() ([]*Dashboard, error)
	InsertAdmin(admin *customermitra.Admin) (*customermitra.Admin, error)
	RenewAdmin(id int, admin *Admin) (*Admin, error)
	LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error)
}

type Service interface {
	Dashboard() ([]*Dashboard, error)
	CreateAdmin(admin *customermitra.Admin) (*customermitra.Admin, error)
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

func (s *service) Dashboard() ([]*Dashboard, error) {
	return s.repository.Dashboard()
}

func (s *service) CreateAdmin(admin *customermitra.Admin) (*customermitra.Admin, error) {
	err := s.validate.Struct(admin)
	if err != nil {
		return nil, err
	}
	admin, err = s.repository.InsertAdmin(admin)
	return admin, err
}

func (s *service) UpdateAdmin(id int, admin *Admin) (*Admin, error) {
	return s.repository.RenewAdmin(id, admin)
}

func (s *service) LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error) {
	tokens, err := s.repository.LoginAdmin(Auth)
	return tokens, err
}
