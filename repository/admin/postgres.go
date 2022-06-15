package admin

import (
	"api-redeem-point/business/admin"
	"api-redeem-point/business/customermitra"
	"api-redeem-point/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
	err := repo.db.Where("Status_transaction = ?", "PENDING").Preload("Customers").Find(&History).Error
	fmt.Println(History)
	if err != nil {
		return nil, err
	}
	for _, v := range History {
		var tmpDashboard admin.Dashboard
		tmpDashboard.ID_Transaction = v.ID_Transaction
		tmpDashboard.Nomor = v.Nomor
		tmpDashboard.Customer_id = v.Customer_id
		tmpDashboard.Customer.ID = v.Customers.ID
		tmpDashboard.Customer.Email = v.Customers.Email
		tmpDashboard.Customer.Fullname = v.Customers.Fullname
		tmpDashboard.Description = v.Description
		tmpDashboard.Status_transaction = v.Status_Transaction

		Dashboard = append(Dashboard, &tmpDashboard)
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
func (repo *PosgresRepository) InsertAdmin(Admins *customermitra.Admin) (*customermitra.Admin, error) {
	err := repo.db.Create(&Admins).Error
	if err != nil {
		return nil, fmt.Errorf("failed insert data")
	}
	return Admins, nil
}

func (repo *PosgresRepository) LoginAdmin(Auth *admin.AuthLogin) (*admin.ResponseLogin, error) {
	var Admin admin.Admin
	err := repo.db.Where("email =? AND password = ?", Auth.Email, Auth.Password).First(&Admin).Error
	if err != nil {
		return nil, err
	}
	expirationTime := time.Now().Add(5 * time.Hour)

	claims := &admin.Claims{
		ID:    Admin.ID,
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
		ID:       Admin.ID,
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
