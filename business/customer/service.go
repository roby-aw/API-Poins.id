package customer

import (
	"api-redeem-point/utils"
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	GetCustomersByID(id int) (*Customers, error)
	SignCustomer(login *AuthLogin) (*ResponseLogin, error)
	InsertCustomer(Data *RegisterCustomer) (*RegisterCustomer, error)
	UpdateCustomer(Data *UpdateCustomer) (*UpdateCustomer, error)
	HistoryCustomer(id int, pagination utils.Pagination) ([]History, error)
	DetailHistoryCustomer(idtransaction string) (*DetailHistory, error)
	ClaimPulsa(Data *RedeemPulsaData) error
	ClaimPaketData(Data *RedeemPulsaData) error
	ClaimBank(emoney *InputTransactionBankEmoney) (*InputTransactionBankEmoney, error)
	TakeCallback(data *Disbursement) (*Disbursement, error)
	GetOrderEmoney(emoney *InputTransactionBankEmoney) (*InputTransactionBankEmoney, error)
	InsertStore(store *RegisterStore) (*RegisterStore, error)
	DecraseStock(id int, stock int) error
}

type Service interface {
	FindCustomersByID(id int) (*Customers, error)
	LoginCustomer(login *AuthLogin) (*ResponseLogin, error)
	CreateCustomer(Data *RegisterCustomer) (*RegisterCustomer, error)
	UpdateCustomer(Data *UpdateCustomer) (*UpdateCustomer, error)
	HistoryCustomer(id int, pagination utils.Pagination) ([]History, error)
	DetailHistoryCustomer(idtransaction string) (*DetailHistory, error)
	RedeemPulsa(Data *RedeemPulsaData) error
	RedeemPaketData(Data *RedeemPulsaData) error
	RedeemBank(Data *InputTransactionBankEmoney) (*InputTransactionBankEmoney, error)
	GetCallback(data *Disbursement) (*Disbursement, error)
	ToOrderEmoney(emoney *InputTransactionBankEmoney) (*InputTransactionBankEmoney, error)
	DecraseStock(id int, stock int) error
	CreateStore(store *RegisterStore) (*RegisterStore, error)
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

func (s *service) FindCustomersByID(id int) (*Customers, error) {
	return s.repository.GetCustomersByID(id)
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
	err = s.validate.Var(strconv.Itoa(Data.Pin), "len=4")
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
	if Data.Email != "" {
		err = s.validate.Var(Data.Email, "email")
	}
	return s.repository.UpdateCustomer(Data)
}

func (s *service) HistoryCustomer(id int, pagination utils.Pagination) ([]History, error) {
	result, err := s.repository.HistoryCustomer(id, pagination)
	if len(result) == 0 {
		err = errors.New("Tidak ada transaksi")
		return nil, err
	}
	return result, nil
}

func (s *service) DetailHistoryCustomer(idtransaction string) (*DetailHistory, error) {
	return s.repository.DetailHistoryCustomer(idtransaction)
}

func (s *service) RedeemPulsa(Data *RedeemPulsaData) error {
	err := s.validate.Struct(Data)
	if err != nil {
		return err
	}
	result, err := s.repository.GetCustomersByID(Data.Customer_id)
	if err != nil {
		return err
	}
	if result.Poin < Data.Poin_redeem {
		err := errors.New("Poin kurang")
		return err
	}
	err = s.repository.ClaimPulsa(Data)
	if err != nil {
		return err
	}
	err = s.repository.DecraseStock(2, Data.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) RedeemPaketData(Data *RedeemPulsaData) error {
	err := s.validate.Struct(Data)
	if err != nil {
		return err
	}
	result, err := s.repository.GetCustomersByID(Data.Customer_id)
	if err != nil {
		return err
	}
	if result.Poin < Data.Poin_redeem {
		err := errors.New("Poin kurang")
		return err
	}
	err = s.repository.ClaimPaketData(Data)
	if err != nil {
		return err
	}
	err = s.repository.DecraseStock(2, Data.Poin_redeem)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) RedeemBank(Data *InputTransactionBankEmoney) (*InputTransactionBankEmoney, error) {
	err := s.validate.Struct(Data)
	if err != nil {
		return nil, err
	}
	result, err := s.repository.GetCustomersByID(Data.Customer_id)
	if err != nil {
		return nil, err
	}
	if result.Poin < Data.Poin_redeem {
		err := errors.New("Poin kurang")
		return nil, err
	}
	Data, err = s.repository.ClaimBank(Data)
	if err != nil {
		return nil, err
	}
	err = s.repository.DecraseStock(1, Data.Amount)
	if err != nil {
		return nil, err
	}
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
	result, err := s.repository.GetCustomersByID(emoney.Customer_id)
	if err != nil {
		return nil, err
	}
	if result.Poin < emoney.Poin_redeem {
		err := errors.New("Poin kurang")
		return nil, err
	}
	emoney, err = s.repository.GetOrderEmoney(emoney)
	if err != nil {
		return nil, err
	}
	err = s.repository.DecraseStock(1, emoney.Amount)
	if err != nil {
		return nil, err
	}
	return emoney, err
}

func (s *service) CreateStore(store *RegisterStore) (*RegisterStore, error) {
	err := s.validate.Struct(store)
	if err != nil {
		return nil, err
	}
	return s.repository.InsertStore(store)
}

func (s *service) DecraseStock(id int, stock int) error {
	return s.repository.DecraseStock(id, stock)
}
