package admin

import (
	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindAdmins() (admins []Admin, err error)
	FindAdminByID(id int) (*Admin, error)
	InsertAdmin(admin *Admin) (*Admin, error)
	RemoveAdmin(id int) error
	RenewAdmin(id int, admin *Admin) (*Admin, error)
	CreateToken(admins *Admin) (string, error)
}

type Service interface {
	GetAdmins() (Admins []Admin, err error)
	GetAdminByID(id int) (*Admin, error)
	CreateAdmin(admin *Admin) (*Admin, error)
	DeleteAdmin(id int) error
	UpdateAdmin(id int, admin *Admin) (*Admin, error)
	GetToken(admins *Admin) (string, error)
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

func (s *service) GetAdmins() (admins []Admin, err error) {
	admins, err = s.repository.FindAdmins()
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (s *service) GetAdminByID(id int) (*Admin, error) {
	return s.repository.FindAdminByID(id)
}

func (s *service) CreateAdmin(admin *Admin) (*Admin, error) {
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

func (s *service) GetToken(admins *Admin) (string, error) {
	tokens, err := s.repository.CreateToken(admins)
	return tokens, err
}
