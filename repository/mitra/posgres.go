package mitra

import (
	"api-redeem-point/business/mitra"
	"api-redeem-point/repository"
	"api-redeem-point/utils"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type PosgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PosgresRepository {
	return &PosgresRepository{
		db: db,
	}
}

func (repo *PosgresRepository) SignStore(store *mitra.AuthStore) (*mitra.ResponseLoginStore, error) {
	var tmpStore *mitra.Store
	err := repo.db.Where("email = ?", store.Email).First(&tmpStore).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("Email salah")
			return nil, err
		}
	}
	err = VerifyPassword(tmpStore.Password, store.Password)
	if err != nil {
		err = errors.New("Password salah")
		return nil, err
	}
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &mitra.ClaimsMitra{
		ID:    int(tmpStore.ID),
		Email: tmpStore.Email,
		Store: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SECRET_KEY := os.Getenv("SECRET_JWT")
	token_jwt, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, err
	}
	Response := mitra.ResponseLoginStore{
		Store: *tmpStore,
		Token: token_jwt,
	}
	return &Response, nil
}

func (repo *PosgresRepository) InputPoin(input *mitra.InputPoin) (*int, error) {
	var tmpCustomer mitra.Customers
	err := repo.db.Model(mitra.Customers{}).Where("ID = ?", input.Customer_id).First(&tmpCustomer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("wrong id customer")
			return nil, err
		}
	}
	var store mitra.Store
	err = repo.db.Model(repository.Store{}).Where("ID = ?", input.Store_id).First(&store).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("wrong id store")
			return nil, err
		}
	}
	var i int
	price := input.Amount
	for i = 0; price >= 100; i = i + 1 {
		price = price - 100
	}
	random := utils.Randomstring()
	transaction := mitra.History_Transaction{
		ID_Transaction: "IP" + random,
		Customer_id:    input.Customer_id,
		Store_id:       input.Store_id,
		Amount:         input.Amount,
		Poin_Redeem:    i,
		Status_Poin:    "IN",
		Poin_Account:   tmpCustomer.Poin,
	}
	err = repo.db.Create(&transaction).Error
	if err != nil {
		return nil, err
	}
	tmpCustomer.Poin = tmpCustomer.Poin + i
	err = repo.db.Model(tmpCustomer).Select("Poin").Updates(tmpCustomer).Error
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
