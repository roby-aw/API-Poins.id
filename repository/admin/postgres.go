package admin

import (
	"api-redeem-point/business/admin"
	"api-redeem-point/business/customermitra"
	"api-redeem-point/config"
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

func (repo *PosgresRepository) Dashboard() ([]*admin.Dashboard, error) {
	var History []customermitra.History_Transaction
	var Dashboard []*admin.Dashboard
	err := repo.db.Model(&customermitra.History_Transaction{}).Where("Status_transaction = ?", "PENDING").Preload("Customers").Find(&Dashboard).Error
	fmt.Println(History)
	if err != nil {
		return nil, err
	}
	return Dashboard, nil
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

func (repo *PosgresRepository) GetHistoryCustomers() ([]admin.CustomerHistory, error) {
	var CustomerHistory []admin.CustomerHistory
	var History_Transaction []*customermitra.History_Transaction
	err := repo.db.Where("Status_Poin = ?", "OUT").Order("created_at desc").Preload("Customers", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name")
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

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
