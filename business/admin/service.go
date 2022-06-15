package admin

import (
	"api-redeem-point/business/customermitra"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	Dashboard() ([]*Dashboard, error)
	InsertAdmin(admin *RegisterAdmin) (*RegisterAdmin, error)
	AcceptTransaction(idtransaction string) error
	RenewAdmin(id int, admin *Admin) (*Admin, error)
	LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error)
	GetCustomers() ([]*customermitra.Customers, error)
}

type Service interface {
	Dashboard() ([]*Dashboard, error)
	CreateAdmin(admin *RegisterAdmin) (*RegisterAdmin, error)
	ApproveTransaction(idtransaction string) error
	UpdateAdmin(id int, admin *Admin) (*Admin, error)
	LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error)
	FindCustomers() ([]*customermitra.Customers, error)
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

func (s *service) CreateAdmin(admin *RegisterAdmin) (*RegisterAdmin, error) {
	err := s.validate.Struct(admin)
	if err != nil {
		return nil, err
	}
	admin, err = s.repository.InsertAdmin(admin)
	return admin, err
}

func (s *service) ApproveTransaction(idtransaction string) error {
	return s.repository.AcceptTransaction(idtransaction)
}

func (s *service) UpdateAdmin(id int, admin *Admin) (*Admin, error) {
	return s.repository.RenewAdmin(id, admin)
}

func (s *service) LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error) {
	err := s.validate.Struct(Auth)
	if err != nil {
		return nil, err
	}
	tokens, err := s.repository.LoginAdmin(Auth)
	return tokens, err
}

func (s *service) FindCustomers() ([]*customermitra.Customers, error) {
	return s.repository.GetCustomers()
}
