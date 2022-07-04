package admin_test

import (
	"api-redeem-point/business/admin"
	"api-redeem-point/business/customermitra"
	"api-redeem-point/utils"
	"errors"
	"os"
	"reflect"
	"testing"
)

var service admin.Service
var admin1, admin2, admin3, updateadmin admin.Admin
var customer1, customer2, customer3 admin.Customers
var InsertAdmin admin.Admin
var insertSpec, updateSpec, failedSpec, errorspec admin.RegisterAdmin
var stockProduct1, stockProduct2 admin.StockProduct
var TransactionMonth1, TransactionMonth2 admin.TransactionMonth

var errorFindID int

var errorInsert error = errors.New("error on insert")
var errorFind error = errors.New("error on find")

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetAdminByID(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.FindAdminByID(int(admin1.ID))
		if !reflect.DeepEqual(*result, admin1) {
			t.Error("Expect content has to be equal with content1", result, admin1)
		}
	})
	t.Run("Expect not found the result", func(t *testing.T) {
		result, err := service.FindAdminByID(int(100))
		if err != nil {
			t.Error("Expect error is nil. Error", err)
		} else if result != nil {
			t.Error("Expect result must be not found (nil)")
		}
	})
}
func TestInsertAdmin(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, err := service.CreateAdmin(&insertSpec)
		if err != nil {
			t.Error("Cannot insert admin")
		}
		if result.Email != insertSpec.Email {
			t.Error("Expect email admin is equal to email insert admin")
		}
		NewAdmin, _ := service.FindAdminByID(4)
		if NewAdmin == nil {
			t.Error("expect admins is not nil after inserted")
			t.FailNow()
		}
	})
}

func TestDashboard(t *testing.T) {
	t.Run("Expect found the result", func(t *testing.T) {
		result, _ := service.Dashboard()
		if result == nil {
			t.Error("expect result is nil")
		}
		if result.Month == nil {
			t.Error("expect data month is nil")
		}
		if result.Stock == nil {
			t.Error("expect stock is nil")
		}
		if len(result.Stock) != 2 {
			t.Error("expect lenght stock is 2", len(result.Stock))
		}
	})
}

func setup() {
	admin1.ID = 1
	admin1.Email = "testemail@gmail.com"
	admin1.Fullname = "testname"
	admin1.Password = "testpassword"
	admin1.No_hp = "08565895685"

	admin2.ID = 2
	admin2.Email = "testemail2@gmail.com"
	admin2.Fullname = "testname2"
	admin2.Password = "testpassword"
	admin2.No_hp = "08565897685"

	admin3.ID = 3
	admin3.Email = "testemail3@gmail.com"
	admin3.Fullname = "testname3"
	admin3.Password = "testpassword"
	admin3.No_hp = "08565897665"

	stockProduct1.ID = 1
	stockProduct1.Product = "PulsaAndPaketData"
	stockProduct1.Balance = 500000

	stockProduct2.ID = 2
	stockProduct2.Product = "CashoutAndEmoney"
	stockProduct2.Balance = 500000

	TransactionMonth1.Day = "01"
	TransactionMonth2.Count = 5

	TransactionMonth1.Day = "02"
	TransactionMonth2.Count = 10

	repo := newInMemoryRepository()
	service = admin.NewService(&repo)

	insertSpec.Email = "insertadmin@gmail.com"
	insertSpec.Fullname = "insertfullname"
	insertSpec.No_hp = "0854696963"
	insertSpec.Password = "insertpassword"
}

type inMemoryRepository struct {
	Admin            map[int]admin.Admin
	AllAdmin         []admin.Admin
	Product          map[int]admin.StockProduct
	AllProduct       []admin.StockProduct
	TransactionMonth []admin.TransactionMonth
}

func newInMemoryRepository() inMemoryRepository {
	var repo inMemoryRepository
	repo.Admin = make(map[int]admin.Admin)
	repo.Admin[int(admin1.ID)] = admin1
	repo.Admin[int(admin2.ID)] = admin2
	repo.Admin[int(admin3.ID)] = admin3

	repo.Product = make(map[int]admin.StockProduct)
	repo.Product[stockProduct1.ID] = stockProduct1
	repo.Product[stockProduct2.ID] = stockProduct2

	repo.AllAdmin = []admin.Admin{}
	repo.AllAdmin = append(repo.AllAdmin, admin1)
	repo.AllAdmin = append(repo.AllAdmin, admin2)
	repo.AllAdmin = append(repo.AllAdmin, admin3)

	repo.AllProduct = []admin.StockProduct{}
	repo.AllProduct = append(repo.AllProduct, stockProduct1)
	repo.AllProduct = append(repo.AllProduct, stockProduct2)

	repo.TransactionMonth = []admin.TransactionMonth{}
	repo.TransactionMonth = append(repo.TransactionMonth, TransactionMonth1)
	repo.TransactionMonth = append(repo.TransactionMonth, TransactionMonth2)

	return repo
}

func (repo *inMemoryRepository) GetAdminByID(id int) (*admin.Admin, error) {
	if id == errorFindID {
		return nil, errorFind
	}

	admin, ok := repo.Admin[id]
	if !ok {
		return nil, nil
	}
	return &admin, nil
}

func (repo *inMemoryRepository) InsertAdmin(admins *admin.RegisterAdmin) (*admin.RegisterAdmin, error) {
	if admins.Fullname == errorspec.Fullname {
		return nil, errorInsert
	}
	adminInsert := admin.Admin{
		ID:       4,
		Email:    admins.Email,
		Fullname: admins.Fullname,
		Password: admins.Password,
	}
	repo.AllAdmin = append(repo.AllAdmin, adminInsert)
	repo.Admin[int(adminInsert.ID)] = adminInsert

	return admins, nil
}

func (repo *inMemoryRepository) Dashboard() (*int, error) {
	today := 4
	return &today, nil
}

func (repo *inMemoryRepository) GetProduct() ([]admin.StockProduct, error) {
	stock := repo.AllProduct
	return stock, nil
}

func (repo *inMemoryRepository) GetTransactionMonthDay() ([]admin.TransactionMonth, error) {
	month := repo.TransactionMonth
	return month, nil
}

func (repo *inMemoryRepository) TransactionPending(pagination utils.Pagination) ([]*admin.TransactionPending, error) {
	return nil, nil
}

func (repo *inMemoryRepository) AcceptTransaction(idtransaction string) error {
	return nil
}
func (repo *inMemoryRepository) LoginAdmin(Auth *admin.AuthLogin) (*admin.ResponseLogin, error) {
	return nil, nil
}
func (repo *inMemoryRepository) RenewAdmin(id int, admin *admin.Admin) (*admin.Admin, error) {
	return nil, nil
}
func (repo *inMemoryRepository) GetCustomers(pagination utils.Pagination) ([]*customermitra.Customers, error) {
	return nil, nil
}
func (repo *inMemoryRepository) GetHistoryCustomers(pagination utils.Pagination) ([]admin.CustomerHistory, error) {
	return nil, nil
}
func (repo *inMemoryRepository) DeleteCustomer(id int) error {
	return nil
}
func (repo *inMemoryRepository) TransactionDate() ([]admin.TransactionDate, error) {
	return nil, nil
}
func (repo *inMemoryRepository) TransactionByDate(startdate string, enddate string) ([]admin.TransactionDate, error) {
	return nil, nil
}
func (repo *inMemoryRepository) UpdateCustomer(data admin.UpdateCustomer) (*admin.UpdateCustomer, error) {
	return nil, nil
}
func (repo *inMemoryRepository) UpdateCustomerPoint(id int, point int) (*int, error) {
	return nil, nil
}
func (repo *inMemoryRepository) UpdateStock(id int, stock int) (*admin.StockProduct, error) {
	return nil, nil
}
func (repo *inMemoryRepository) HistoryStore(pagination utils.Pagination, name string) ([]admin.HistoryStore, error) {
	return nil, nil
}
func (repo *inMemoryRepository) DeleteStore(id int) error {
	return nil
}
func (repo *inMemoryRepository) GetStore(pagination utils.Pagination, name string) ([]*customermitra.Store, error) {
	return nil, nil
}
func (repo *inMemoryRepository) UpdateStore(store admin.UpdateStore) (*admin.UpdateStore, error) {
	return nil, nil
}
