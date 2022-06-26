package admin

import (
	"api-redeem-point/business/admin"
	"api-redeem-point/business/customermitra"
	"api-redeem-point/config"
	"api-redeem-point/utils"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type PosgresRepository struct {
	db *gorm.DB
}

func NewPosgresRepository(db *gorm.DB) *PosgresRepository {
	return &PosgresRepository{
		db: db,
	}
}

func (repo *PosgresRepository) Dashboard() (*int, error) {
	var today64 int64
	now := time.Now().Format("2006-01-02")
	err := repo.db.Model(&customermitra.History_Transaction{}).Where("created_at > ?", now+" 00:00:00").Where("status_poin = ?", "OUT").Count(&today64).Error
	if err != nil {
		return nil, err
	}
	var today int
	today = int(today64)
	return &today, err
}

func (repo *PosgresRepository) TransactionPending() ([]*admin.TransactionPending, error) {
	var History []customermitra.History_Transaction
	var Pending []*admin.TransactionPending
	err := repo.db.Model(&customermitra.History_Transaction{}).Where("Status_transaction = ?", "PENDING").Preload("Customers", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "fullname")
	}).Find(&Pending).Error
	fmt.Println(History)
	if err != nil {
		return nil, err
	}
	return Pending, nil
}

func (repo *PosgresRepository) RemoveAdmin(id int) error {
	var admin *admin.Admin
	err := repo.db.Where("ID = ?", id).First(&admin).Error
	if err != nil {
		return err
	}
	fmt.Println(admin)
	err = repo.db.Delete(admin, id).Error
	if err != nil {
		return err
	}
	return err
}
func (repo *PosgresRepository) InsertAdmin(Admins *admin.RegisterAdmin) (*admin.RegisterAdmin, error) {
	password, _ := Hash(Admins.Password)
	var admin = admin.Admin{Email: Admins.Email, Password: string(password), Fullname: Admins.Fullname, No_hp: Admins.No_hp}
	err := repo.db.Create(&admin).Error
	if err != nil {
		return nil, fmt.Errorf("failed insert data")
	}
	return Admins, nil
}

func (repo *PosgresRepository) LoginAdmin(Auth *admin.AuthLogin) (*admin.ResponseLogin, error) {
	var Admin admin.Admin
	err := repo.db.Where("email = ?", Auth.Email).First(&Admin).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("Email salah")
			return nil, err
		}
	}

	err = VerifyPassword(Admin.Password, Auth.Password)
	if err != nil {
		err = errors.New("Password salah")
		return nil, err
	}
	expirationTime := time.Now().Add(5 * time.Hour)

	claims := &admin.Claims{
		ID:    int(Admin.ID),
		Email: Admin.Email,
		Role:  "Admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SECRET_KEY := config.GetConfig().Secrettoken.Token
	token_jwt, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, err
	}
	response := admin.ResponseLogin{
		ID:       int(Admin.ID),
		Email:    Admin.Email,
		Fullname: Admin.Fullname,
		Password: Admin.Password,
		No_hp:    Admin.No_hp,
		Token:    token_jwt,
	}

	return &response, nil
}

func (repo PosgresRepository) RenewAdmin(id int, admin *admin.Admin) (*admin.Admin, error) {
	err := repo.db.Model(*admin).Where("ID = ?", id).Updates(admin).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Where("ID = ?", id).First(admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo *PosgresRepository) AcceptTransaction(idtransaction string) error {
	var transaction *customermitra.History_Transaction
	err := repo.db.Where("ID_Transaction = ?", idtransaction).Take(&transaction).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("wrong id transaction")
			return err
		}
	}
	if transaction.Status_Transaction == "COMPLETED" {
		err = errors.New("Transaction already completed")
		return err
	}
	transaction.Status_Transaction = "COMPLETED"
	err = repo.db.Updates(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *PosgresRepository) GetCustomers() ([]*customermitra.Customers, error) {
	var customers []*customermitra.Customers
	err := repo.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (repo *PosgresRepository) GetHistoryCustomers(pagination utils.Pagination) ([]admin.CustomerHistory, error) {
	var CustomerHistory []admin.CustomerHistory
	var History_Transaction []*customermitra.History_Transaction
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := repo.db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	err := queryBuider.Where("Status_Poin = ?", "OUT").Order("created_at desc").Preload("Customers", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "fullname")
	}).Find(&History_Transaction).Error
	if err != nil {
		return nil, err
	}
	for _, v := range History_Transaction {
		var tmpHistory admin.CustomerHistory
		tmpHistory.Customer_id = v.Customer_id
		tmpHistory.Customers.ID = v.Customers.ID
		tmpHistory.Customers.Email = v.Customers.Email
		tmpHistory.Customers.Fullname = v.Customers.Fullname
		tmpHistory.Description = v.Description
		tmpHistory.Nomor = v.Nomor
		tmpHistory.CreatedAt = v.CreatedAt
		tmpHistory.Status_Transaction = v.Status_Transaction
		tmpHistory.Poin_redeem = v.Poin_Redeem

		CustomerHistory = append(CustomerHistory, tmpHistory)
	}

	return CustomerHistory, nil
}

func (repo *PosgresRepository) TransactionDate() ([]admin.TransactionDate, error) {
	var transaction []admin.TransactionDate
	repo.db.Model(&customermitra.History_Transaction{}).Where("created_at > ?", "2022-01-01").Where("created_at < ?", "2023-01-01").Order("created_at asc").Find(&transaction)
	return transaction, nil
}

func (repo *PosgresRepository) TransactionByDate(startdate string, enddate string) ([]admin.TransactionDate, error) {
	var transaction []admin.TransactionDate
	repo.db.Raw("select * from history_transactions where created_at BETWEEN ? AND ?", startdate+" 00:00:00", enddate+" 23:59:59").Find(&transaction)
	return transaction, nil
}

func (repo *PosgresRepository) UpdateCustomer(data customermitra.Customers) (*customermitra.Customers, error) {
	err := repo.db.Model(&customermitra.Customers{}).Where("ID = ?", data.ID).Updates(&customermitra.Customers{Email: data.Email, Fullname: data.Fullname, No_hp: data.No_hp}).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (repo *PosgresRepository) UpdateCustomerPoint(id int, point int) (*int, error) {
	var data customermitra.Customers
	err := repo.db.Model(&customermitra.Customers{}).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	hasil := data.Poin + point
	data.Poin = hasil
	err = repo.db.Model(&customermitra.Customers{}).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return nil, err
	}
	return &hasil, err
}

func (repo *PosgresRepository) GetProduct() ([]admin.StockProduct, error) {
	var stock []admin.StockProduct
	repo.db.Find(&stock)
	return stock, nil
}

func (repo *PosgresRepository) UpdateStock(id int, stock int) (*admin.StockProduct, error) {
	var product admin.StockProduct
	err := repo.db.Model(&customermitra.StockProduct{}).Where("id = ?", id).Find(&product).Error
	if err != nil {
		return nil, err
	}
	sum := product.Balance + stock
	product.Balance = sum
	repo.db.Model(&customermitra.StockProduct{}).Where("id = ?", id).Updates(customermitra.StockProduct{Balance: sum})
	return &product, nil
}

func (repo *PosgresRepository) TestDB() ([]admin.TransactionMonth, error) {
	var TransactionMonth []admin.TransactionMonth
	year := time.Now().Year()
	err := repo.db.Raw("SELECT TO_CHAR(created_at, 'Month') AS Month, COUNT(1) AS count FROM history_transactions where EXTRACT(YEAR From created_at) = ? AND status_poin = ? GROUP BY TO_CHAR(created_at, 'Month')", year, "OUT").Find(&TransactionMonth).Error
	if err != nil {
		return nil, err
	}
	return TransactionMonth, nil
}

func (repo *PosgresRepository) HistoryStore() ([]admin.HistoryStore, error) {
	var tmpHistory []admin.HistoryStore
	err := repo.db.Model(customermitra.History_Transaction{}).Where("status_poin = ?", "IN").Select("created_at", "poin_redeem", "amount", "store_id", "customer_id").Preload("customer").Preload("store").Find(&tmpHistory).Error
	if err != nil {
		return nil, err
	}
	return tmpHistory, nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
