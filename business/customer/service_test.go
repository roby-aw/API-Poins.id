package customer_test

import (
	customerBusiness "api-redeem-point/business/customer"
	"api-redeem-point/utils"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

var service customerBusiness.Service
var customer1, customer2, customer3 customerBusiness.Customers
var history1, history2, history3 customerBusiness.History_Transaction
var stockProduct1, stockProduct2 customerBusiness.StockProduct
var InsertCustomer, insertCustomernull, insCusErrorPin customerBusiness.RegisterCustomer
var login, loginnull customerBusiness.AuthLogin
var UpdateCustomer, failUpdateCustomer customerBusiness.UpdateCustomer
var pagination utils.Pagination
var RedeemPulsa, RedeemData, RedeemPulsaLess, RedeemWrongIDPulsaData, RedeemPulsaDataNull, RedeemWrongPin, RedeemUnliPoin customerBusiness.RedeemPulsaData
var RedeemBank, RedeemEmoney, RedeemBankEmoneyLess, RedeemWrongIDBankEmoney, RedeemBankEmoneyNull, RedeemBankEmoneyWrongPin, RedeemBankEmoneyUnliPoin customerBusiness.InputTransactionBankEmoney

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetCustomerByID(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.FindCustomersByID(int(customer3.ID))
		if result.ID != customer3.ID {
			t.Error("Expect found customer id 3")
		}
	})
}

func TestSignCustomer(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.LoginCustomer(&login)
		if result == nil {
			t.Error("error cant login customer")
		}
	})
	t.Run("Expect error validation", func(t *testing.T) {
		_, err := service.LoginCustomer(&loginnull)
		if err == nil {
			t.Error("expect found error")
		}
	})
}

func TestCreateCustomer(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.CreateCustomer(&InsertCustomer)
		if result == nil {
			t.Error("expect found the result")
		}
	})
	t.Run("Expect error validation struct", func(t *testing.T) {
		_, err := service.CreateCustomer(&insertCustomernull)
		if err == nil {
			t.Error("expect found error")
		}
	})
	t.Run("Expect error validation pin", func(t *testing.T) {
		_, err := service.CreateCustomer(&insCusErrorPin)
		if err == nil {
			t.Error("expect found error len pin")
		}
	})
}

func TestUpdateCustomers(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.UpdateCustomer(&UpdateCustomer)
		if result == nil {
			t.Error("expect found the result")
		}
	})
	t.Run("Expect error validation struct", func(t *testing.T) {
		_, err := service.UpdateCustomer(&failUpdateCustomer)
		if err == nil {
			t.Error("expect found error")
		}
	})
}

func TestGetHistoryCustomers(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.HistoryCustomer(int(customer1.ID), pagination)
		if len(result) != 3 {
			t.Error("expect result is lenght 3")
		}
	})
	t.Run("Expect got error nil result", func(t *testing.T) {
		_, err := service.HistoryCustomer(5, pagination)
		if err == nil {
			t.Error("expect found error")
		}
	})
}

func TestGetDetailHistory(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.DetailHistoryCustomer(history2.ID_Transaction)
		if result == nil {
			t.Error("expect found the result")
		}
	})
}

func TestRedeemPulsa(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		err := service.RedeemPulsa(&RedeemPulsa)
		if err != nil {
			t.Error("got error ", err)
		}
	})
	t.Run("Expect error validation", func(t *testing.T) {
		err := service.RedeemPulsa(&RedeemPulsaDataNull)
		if err == nil {
			t.Error("expect found error")
		}
	})
	t.Run("Expect error not found result", func(t *testing.T) {
		err := service.RedeemPulsa(&RedeemWrongIDPulsaData)
		if err == nil {
			t.Error("expect found error not found")
		}
	})
	t.Run("Expect error poin kurang", func(t *testing.T) {
		err := service.RedeemPulsa(&RedeemPulsaLess)
		if err == nil {
			t.Error("expect found error")
		}
	})
	t.Run("Expect error claim pulsa", func(t *testing.T) {
		err := service.RedeemPulsa(&RedeemWrongPin)
		if err == nil {
			t.Error("expect found error wrong pin")
		}
	})
	t.Run("Expect error stock not available", func(t *testing.T) {
		err := service.RedeemPulsa(&RedeemUnliPoin)
		if err == nil {
			t.Error("expect found error stock not available")
		}
	})
}

func TestRedeemPaketData(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		err := service.RedeemPaketData(&RedeemData)
		if err != nil {
			t.Error("Expect error not nil")
		}
	})
	t.Run("Expect error validation", func(t *testing.T) {
		err := service.RedeemPaketData(&RedeemPulsaDataNull)
		if err == nil {
			t.Error("expect found error")
		}
	})
	t.Run("Expect error not found result", func(t *testing.T) {
		err := service.RedeemPaketData(&RedeemWrongIDPulsaData)
		if err == nil {
			t.Error("expect found error not found")
		}
	})
	t.Run("Expect error poin kurang", func(t *testing.T) {
		err := service.RedeemPaketData(&RedeemPulsaLess)
		if err == nil {
			t.Error("expect found error")
		}
	})
	t.Run("Expect error claim pulsa", func(t *testing.T) {
		err := service.RedeemPaketData(&RedeemWrongPin)
		if err == nil {
			t.Error("expect found error wrong pin")
		}
	})
	t.Run("Expect error stock not available", func(t *testing.T) {
		err := service.RedeemPaketData(&RedeemUnliPoin)
		if err == nil {
			t.Error("expect found error stock not available")
		}
	})
}

func TestRedeemBank(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.RedeemBank(&RedeemBank)
		if result == nil {
			t.Error("expect not found the result")
		}
	})
	t.Run("Expect error validation", func(t *testing.T) {
		_, err := service.RedeemBank(&RedeemBankEmoneyNull)
		if err == nil {
			t.Error("expect found error")
		}
	})
}

func TestRedeemEmoney(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.ToOrderEmoney(&RedeemBank)
		if result == nil {
			t.Error("expect not found the result")
		}
	})
	t.Run("Expect error validation", func(t *testing.T) {
		_, err := service.ToOrderEmoney(&RedeemBankEmoneyNull)
		if err == nil {
			t.Error("expect found error")
		}
	})
}

func setup() {
	customer1.ID = 1
	customer1.CreatedAt = time.Now()
	customer1.Email = "testcustomer1@gmail.com"
	customer1.Fullname = "testcustomer1"
	customer1.Password = "testpassword1"
	customer1.Pin = 1234
	customer1.Poin = 500000

	customer2.ID = 2
	customer2.CreatedAt = time.Now()
	customer2.Email = "testcustomer2@gmail.com"
	customer2.Fullname = "testcustomer2"
	customer2.Password = "testpassword2"
	customer2.Pin = 9876
	customer2.Poin = 30000

	customer3.ID = 3
	customer3.CreatedAt = time.Now()
	customer3.Email = "testcustomer3@gmail.com"
	customer3.Fullname = "testcustomer3"
	customer3.Password = "testpassword3"
	customer3.Pin = 6366
	customer3.Poin = 1000

	history1.ID = 1
	history1.ID_Transaction = "T12345"
	history1.Customer_id = 1
	history1.Transaction_type = "Redeem Pulsa"
	history1.Bank_Provider = "TELKOMSEL"
	history1.Nomor = "085696865698"
	history1.Poin_Account = 10000
	history1.Poin_Redeem = 10000
	history1.Amount = 10000
	history1.Description = "TELKOMSEL - 10000"
	history1.Status_Transaction = "COMPLETE"
	history1.Status_Poin = "OUT"

	history2.ID = 2
	history2.ID_Transaction = "T123565"
	history2.Customer_id = 1
	history2.Transaction_type = "Redeem Pulsa"
	history2.Bank_Provider = "TELKOMSEL"
	history2.Nomor = "08569686546"
	history2.Poin_Account = 10000
	history2.Poin_Redeem = 10000
	history2.Amount = 10000
	history2.Description = "TELKOMSEL - 10000"
	history2.Status_Transaction = "PENDING"
	history2.Status_Poin = "OUT"

	history3.ID = 3
	history3.ID_Transaction = "T136768"
	history3.Customer_id = 1
	history3.Transaction_type = "Redeem Pulsa"
	history3.Bank_Provider = "TELKOMSEL"
	history3.Nomor = "08569686546"
	history3.Poin_Account = 10000
	history3.Poin_Redeem = 10000
	history3.Amount = 10000
	history3.Description = "TELKOMSEL - 10000"
	history3.Status_Transaction = "PENDING"
	history3.Status_Poin = "OUT"

	stockProduct1.ID = 1
	stockProduct1.Product = "PulsaAndPaketData"
	stockProduct1.Balance = 50000

	stockProduct2.ID = 2
	stockProduct2.Product = "CashoutAndEmoney"
	stockProduct2.Balance = 50000

	repo := newInMemoryRepository()
	service = customerBusiness.NewService(&repo)

	login.Email = customer2.Email
	login.Password = customer2.Password

	InsertCustomer.Email = "insertcustomer@gmail.com"
	InsertCustomer.Fullname = "insertcustomer"
	InsertCustomer.Password = "insertpassword"
	InsertCustomer.Pin = 1234
	InsertCustomer.No_hp = "0235685685"

	insCusErrorPin.Email = "insCinsCusErrorPin@gmail.com"
	insCusErrorPin.Fullname = "insCinsCusErrorPin"
	insCusErrorPin.Password = "insertpassword"
	insCusErrorPin.Pin = 123464
	insCusErrorPin.No_hp = "0235685685"

	UpdateCustomer.ID = 2
	UpdateCustomer.Email = "updatecustomer@gmail.com"
	UpdateCustomer.Fullname = "updatecustomer"
	UpdateCustomer.Password = "updatepassword"
	UpdateCustomer.No_hp = "085623569898"

	RedeemPulsa.Customer_id = 1
	RedeemPulsa.Amount = 10000
	RedeemPulsa.Bank_Provider = "TELKOMSEL"
	RedeemPulsa.Nomor = "08569686546"
	RedeemPulsa.Pin = customer1.Pin
	RedeemPulsa.Poin_redeem = 10000

	RedeemData.Customer_id = 1
	RedeemData.Amount = 10
	RedeemData.Bank_Provider = "TELKOMSEL"
	RedeemData.Nomor = "08569686546"
	RedeemData.Pin = customer1.Pin
	RedeemData.Poin_redeem = 10000

	RedeemWrongIDPulsaData.Customer_id = 56
	RedeemWrongIDPulsaData.Amount = 10000
	RedeemWrongIDPulsaData.Bank_Provider = "TELKOMSEL"
	RedeemWrongIDPulsaData.Nomor = "08569686546"
	RedeemWrongIDPulsaData.Pin = customer1.Pin
	RedeemWrongIDPulsaData.Poin_redeem = 10000

	RedeemPulsaLess.Customer_id = 3
	RedeemPulsaLess.Amount = 10000
	RedeemPulsaLess.Bank_Provider = "TELKOMSEL"
	RedeemPulsaLess.Nomor = "08569686546"
	RedeemPulsaLess.Pin = customer3.Pin
	RedeemPulsaLess.Poin_redeem = 10000

	RedeemWrongPin.Customer_id = 1
	RedeemWrongPin.Amount = 10000
	RedeemWrongPin.Bank_Provider = "TELKOMSEL"
	RedeemWrongPin.Nomor = "08569686546"
	RedeemWrongPin.Pin = customer2.Pin
	RedeemWrongPin.Poin_redeem = 10000

	RedeemUnliPoin.Customer_id = 1
	RedeemUnliPoin.Amount = 154566
	RedeemUnliPoin.Bank_Provider = "TELKOMSEL"
	RedeemUnliPoin.Nomor = "08569686546"
	RedeemUnliPoin.Pin = customer1.Pin
	RedeemUnliPoin.Poin_redeem = 164646

	RedeemBank.Customer_id = 1
	RedeemBank.Amount = 1000
	RedeemBank.Bank_Provider = "TELKOMSEL"
	RedeemBank.AN_Rekening = customer1.Fullname
	RedeemBank.Nomor = "08569686546"
	RedeemBank.Pin = customer1.Pin
	RedeemBank.Poin_redeem = 1000
}

type inMemoryRepository struct {
	Customer    map[int]customerBusiness.Customers
	AllCustomer []customerBusiness.Customers
	History     map[string]customerBusiness.History_Transaction
	AllHistory  []customerBusiness.History_Transaction
	Product     map[int]customerBusiness.StockProduct
	AllProduct  []customerBusiness.StockProduct
}

func newInMemoryRepository() inMemoryRepository {
	var repo inMemoryRepository
	repo.Customer = make(map[int]customerBusiness.Customers)
	repo.Customer[int(customer1.ID)] = customer1
	repo.Customer[int(customer2.ID)] = customer2
	repo.Customer[int(customer3.ID)] = customer3

	repo.AllCustomer = []customerBusiness.Customers{}
	repo.AllCustomer = append(repo.AllCustomer, customer1)
	repo.AllCustomer = append(repo.AllCustomer, customer2)
	repo.AllCustomer = append(repo.AllCustomer, customer3)

	repo.History = make(map[string]customerBusiness.History_Transaction)
	repo.History[history1.ID_Transaction] = history1
	repo.History[history2.ID_Transaction] = history2
	repo.History[history3.ID_Transaction] = history3

	repo.AllHistory = []customerBusiness.History_Transaction{}
	repo.AllHistory = append(repo.AllHistory, history1)
	repo.AllHistory = append(repo.AllHistory, history2)
	repo.AllHistory = append(repo.AllHistory, history3)

	repo.Product = make(map[int]customerBusiness.StockProduct)
	repo.Product[int(stockProduct1.ID)] = stockProduct1
	repo.Product[int(stockProduct2.ID)] = stockProduct2

	repo.AllProduct = []customerBusiness.StockProduct{}
	repo.AllProduct = append(repo.AllProduct, stockProduct1)
	repo.AllProduct = append(repo.AllProduct, stockProduct2)

	return repo
}

func (repo *inMemoryRepository) GetCustomersByID(id int) (*customerBusiness.Customers, error) {
	data := repo.Customer[id]
	if data.Email == "" {
		err := errors.New("data not found")
		return nil, err
	}
	return &data, nil
}

func (repo *inMemoryRepository) SignCustomer(login *customerBusiness.AuthLogin) (*customerBusiness.ResponseLogin, error) {
	var data customerBusiness.ResponseLogin
	for _, v := range repo.AllCustomer {
		if v.Email == login.Email {
			if v.Password == login.Password {
				data.ID = int(v.ID)
				data.Email = v.Email
				data.Fullname = v.Fullname
				data.No_hp = v.No_hp
				data.Password = v.Password
				data.Poin = v.Poin
				data.Pin = v.Pin
				data.Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MywiRW1haWwiOiJ0ZXN0MUBnbWFpbC5jb20iLCJDdXN0b21lciI6dHJ1ZSwiZXhwIjoxNjU2ODcxMzEzfQ.a9O_RMxG7iJR4tMdBZmL6JY2lZKsiQv3rSegnhv1C00"
			}
		}
	}
	return &data, nil
}
func (repo *inMemoryRepository) InsertCustomer(Data *customerBusiness.RegisterCustomer) (*customerBusiness.RegisterCustomer, error) {
	fmt.Println(Data)
	return Data, nil
}

func (repo *inMemoryRepository) UpdateCustomer(Data *customerBusiness.UpdateCustomer) (*customerBusiness.UpdateCustomer, error) {
	return Data, nil
}

func (repo *inMemoryRepository) HistoryCustomer(id int, pagination utils.Pagination) ([]customerBusiness.History, error) {
	var history []customerBusiness.History
	for _, v := range repo.AllHistory {
		if v.Customer_id == id {
			var tmpHistory customerBusiness.History
			tmpHistory.ID = int(v.ID)
			tmpHistory.CreatedAt = v.CreatedAt
			tmpHistory.ID_Transaction = v.ID_Transaction
			tmpHistory.Status_Transaction = v.Status_Transaction
			tmpHistory.Transaction_type = v.Transaction_type
			history = append(history, tmpHistory)
		}
	}
	return history, nil
}

func (repo *inMemoryRepository) DetailHistoryCustomer(idtransaction string) (*customerBusiness.DetailHistory, error) {
	data := repo.History[idtransaction]
	history := customerBusiness.DetailHistory{
		ID_Transaction:     data.ID_Transaction,
		Transaction_type:   data.Transaction_type,
		CreatedAt:          data.CreatedAt,
		Bank_Provider:      data.Bank_Provider,
		Nomor:              data.Nomor,
		Amount:             data.Amount,
		Poin_account:       data.Poin_Account,
		Poin_redeem:        data.Poin_Redeem,
		Description:        data.Description,
		Status_Transaction: data.Status_Transaction,
	}
	return &history, nil
}
func (repo *inMemoryRepository) ClaimPulsa(Data *customerBusiness.RedeemPulsaData) error {
	customers := repo.Customer[Data.Customer_id]
	if customers.Pin != Data.Pin {
		err := errors.New("wrong pin")
		return err
	}
	return nil
}

func (repo *inMemoryRepository) ClaimPaketData(Data *customerBusiness.RedeemPulsaData) error {
	customers := repo.Customer[Data.Customer_id]
	if customers.Pin != Data.Pin {
		err := errors.New("wrong pin")
		return err
	}
	return nil
}

func (repo *inMemoryRepository) ClaimBank(emoney *customerBusiness.InputTransactionBankEmoney) (*customerBusiness.InputTransactionBankEmoney, error) {
	customers := repo.Customer[emoney.Customer_id]
	if customers.Pin != emoney.Pin {
		err := errors.New("wrong pin")
		return nil, err
	}
	return emoney, nil
}

func (repo *inMemoryRepository) GetOrderEmoney(emoney *customerBusiness.InputTransactionBankEmoney) (*customerBusiness.InputTransactionBankEmoney, error) {
	customers := repo.Customer[emoney.Customer_id]
	if customers.Pin != emoney.Pin {
		err := errors.New("wrong pin")
		return nil, err
	}
	return emoney, nil
}

func (repo *inMemoryRepository) InsertStore(store *customerBusiness.RegisterStore) (*customerBusiness.RegisterStore, error) {
	return nil, nil
}

func (repo *inMemoryRepository) DecraseStock(id int, stock int) error {
	data := repo.Product[id]
	fmt.Println(data.Balance)
	fmt.Println(stock)
	if data.Balance < stock {
		err := errors.New("stock not available")
		return err
	}
	data.Balance = data.Balance - stock
	repo.Product[id] = data
	return nil
}

func (repo *inMemoryRepository) TakeCallback(data *customerBusiness.Disbursement) (*customerBusiness.Disbursement, error) {
	return nil, nil
}
