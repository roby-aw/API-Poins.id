package admin

import (
	"api-redeem-point/business/admin"
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

func (repo *PosgresRepository) FindAdmins() (Admins []admin.Admin, err error) {
	result := repo.db.Find(&Admins)
	if result.Error != nil {
		return nil, result.Error
	}
	return Admins, nil
}

func (repo *PosgresRepository) FindAdminByID(id int) (*admin.Admin, error) {
	var admin *admin.Admin
	err := repo.db.Where("ID = ? ", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
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
func (repo *PosgresRepository) InsertAdmin(Admins *admin.Admin) (*admin.Admin, error) {
	err := repo.db.Create(&Admins).Error
	if err != nil {
		return nil, fmt.Errorf("failed insert data")
	}
	return Admins, nil
}

func (repo *PosgresRepository) CreateToken(Admins *admin.Admin) (string, error) {
	err := repo.db.Where("username =? AND password = ?", Admins.Username, Admins.Password).First(&Admins).Error
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &admin.Claims{
		Username: Admins.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SECRET_KEY := config.GetConfig().Secrettoken.Token
	token_jwt, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return token_jwt, err
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
