package customermitra

import (
	"github.com/go-playground/validator/v10"
)

type Repository interface {
	SignCustomer(login *AuthLogin) (*ResponseLogin, error)
	InsertCustomer(Data *RegisterCustomer) (*RegisterCustomer, error)
	UpdateCustomer(Data *UpdateCustomer) (*UpdateCustomer, error)
	ClaimPulsa(Data *RedeemPulsaData) error
	ClaimPaketData(Data *RedeemPulsaData) error
	ClaimBank(emoney *InputTransactionBankEmoney) (*InputTransactionBankEmoney, error)
	TakeCallback(data *Disbursement) (*Disbursement, error)
	GetOrderEmoney(emoney *InputTransactionBankEmoney) (*InputTransactionBankEmoney, error)
}

type Service interface {
	LoginCustomer(login *AuthLogin) (*ResponseLogin, error)
	CreateCustomer(Data *RegisterCustomer) (*RegisterCustomer, error)
	UpdateCustomer(Data *UpdateCustomer) (*UpdateCustomer, error)
	RedeemPulsa(Data *RedeemPulsaData) error
	RedeemPaketData(Data *RedeemPulsaData) error
	RedeemBank(Data *InputTransactionBankEmoney) (*InputTransactionBankEmoney, error)
	GetCallback(data *Disbursement) (*Disbursement, error)
	ToOrderEmoney(emoney *InputTransactionBankEmoney) (*InputTransactionBankEmoney, error)
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

func (s *service) LoginCustomer(login *AuthLogin) (*ResponseLogin, error) {
	err := s.validate.Struct(login)
	if err != nil {
		return nil, err
	}
	result, err := s.repository.SignCustomer(login)
	return result, err
}

func (s *service) CreateCustomer(Data *RegisterCustomer) (*RegisterCustomer, error) {
	err := s.validate.Struct(Data)
	if err != nil {
		return nil, err
	}
	return s.repository.InsertCustomer(Data)
}

func (s *service) UpdateCustomer(Data *UpdateCustomer) (*UpdateCustomer, error) {
	err := s.validate.Struct(Data)
	if err != nil {
		return nil, err
	}
	return s.repository.UpdateCustomer(Data)
}

func (s *service) RedeemPulsa(Data *RedeemPulsaData) error {
	err := s.validate.Struct(Data)
	if err != nil {
		return err
	}
	return s.repository.ClaimPulsa(Data)
}

func (s *service) RedeemPaketData(Data *RedeemPulsaData) error {
	err := s.validate.Struct(Data)
	if err != nil {
		return err
	}
	return s.repository.ClaimPaketData(Data)
}

func (s *service) RedeemBank(Data *InputTransactionBankEmoney) (*InputTransactionBankEmoney, error) {
	err := s.validate.Struct(Data)
	if err != nil {
		return nil, err
	}
	Data, err = s.repository.ClaimBank(Data)
	return Data, nil
}

func (s *service) GetCallback(data *Disbursement) (*Disbursement, error) {
	return s.repository.TakeCallback(data)
}

func (s *service) ToOrderEmoney(emoney *InputTransactionBankEmoney) (*InputTransactionBankEmoney, error) {
	err := s.validate.Struct(emoney)
	if err != nil {
		return nil, err
	}
	emoney, err = s.repository.GetOrderEmoney(emoney)
	return emoney, err
}
