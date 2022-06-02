package user

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindFoodByName(name string) (foods []Food, err error)
	TakeCallback(data *Disbursement) (*Disbursement, error)
	GetOrderEmoney(emoney *TransactionBank) (*TransactionBank, error)
}

type Service interface {
	GetFoodByName(name string) (foods []Food, err error)
	GetCallback(data *Disbursement) (*Disbursement, error)
	ToOrderEmoney(emoney *TransactionBank) (*TransactionBank, error)
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

func (s *service) GetFoodByName(name string) (foods []Food, err error) {
	fmt.Println("service jalan")
	return s.repository.FindFoodByName(name)
}

func (s *service) GetCallback(data *Disbursement) (*Disbursement, error) {
	return s.repository.TakeCallback(data)
}

func (s *service) ToOrderEmoney(emoney *TransactionBank) (*TransactionBank, error) {
	return s.repository.GetOrderEmoney(emoney)
}
