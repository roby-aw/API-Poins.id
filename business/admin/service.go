package admin

import (
	"api-redeem-point/business/customermitra"
	"api-redeem-point/utils"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	Dashboard() ([]*Dashboard, error)
	InsertAdmin(admin *RegisterAdmin) (*RegisterAdmin, error)
	AcceptTransaction(idtransaction string) error
	RenewAdmin(id int, admin *Admin) (*Admin, error)
	LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error)
	GetCustomers() ([]*customermitra.Customers, error)
	GetHistoryCustomers(pagination utils.Pagination) ([]CustomerHistory, error)
	TransactionDate() ([]TransactionDate, error)
	TransactionByDate(startdate string, enddate string) ([]TransactionDate, error)
	UpdateCustomer(data customermitra.Customers) (*customermitra.Customers, error)
	UpdateCustomerPoint(id int, point int) (*int, error)
	GetProduct() ([]*StockProduct, error)
	UpdateStock(id int, stock int) (*StockProduct, error)
	TestDB() ([]TransactionMonth, error)
}

type Service interface {
	Dashboard() ([]*Dashboard, error)
	CreateAdmin(admin *RegisterAdmin) (*RegisterAdmin, error)
	ApproveTransaction(idtransaction string) error
	UpdateAdmin(id int, admin *Admin) (*Admin, error)
	LoginAdmin(Auth *AuthLogin) (*ResponseLogin, error)
	FindCustomers() ([]*customermitra.Customers, error)
	FindHistoryCustomers(pagination utils.Pagination) ([]CustomerHistory, error)
	TransactionDate() ([]TransactionDate, error)
	TransactionByDate(startdate string, enddate string) ([]TransactionDate, error)
	UpdateCustomer(data customermitra.Customers) (*customermitra.Customers, error)
	UpdateCustomerPoint(id int, point int) (*int, error)
	FindProduct() ([]*StockProduct, error)
	UpdateStock(id int, stock int) (*StockProduct, error)
	TestDB() ([]TransactionMonth, error)
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

func (s *service) FindHistoryCustomers(pagination utils.Pagination) ([]CustomerHistory, error) {
	return s.repository.GetHistoryCustomers(pagination)
}

func (s *service) TransactionDate() ([]TransactionDate, error) {
	return s.repository.TransactionDate()
}

func (s *service) TransactionByDate(startdate string, enddate string) ([]TransactionDate, error) {
	return s.repository.TransactionByDate(startdate, enddate)
}

func (s *service) UpdateCustomer(data customermitra.Customers) (*customermitra.Customers, error) {
	return s.repository.UpdateCustomer(data)
}

func (s *service) UpdateCustomerPoint(id int, point int) (*int, error) {
	return s.repository.UpdateCustomerPoint(id, point)
}

func (s *service) FindProduct() ([]*StockProduct, error) {
	return s.repository.GetProduct()
}

func (s *service) UpdateStock(id int, stock int) (*StockProduct, error) {
	return s.repository.UpdateStock(id, stock)
}

func (s *service) TestDB() ([]TransactionMonth, error) {
	return s.repository.TestDB()
}
