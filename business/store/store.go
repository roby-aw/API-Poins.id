package store

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type History_Transaction struct {
	ID                 uint      `json:"id"`
	CreatedAt          time.Time `json:"createdat"`
	UpdatedAt          time.Time `json:"updatedat"`
	ID_Transaction     string    `json:"id_transaction"`
	Customer_id        int       `json:"customer_id"`
	Customers          Customers `json:"customers" gorm:"foreignkey:ID;references:Customer_id"`
	Store_id           int       `json:"store_id"`
	Store              Store     `json:"store" gorm:"foreignkey:ID;references:Store_id"`
	Transaction_type   string    `json:"transaction_type"`
	Bank_Provider      string    `json:"bank_provider"`
	Nomor              string    `json:"nomor"`
	Poin_Account       int       `json:"poin_account"`
	Poin_Redeem        int       `json:"poin_redeem"`
	Amount             int       `json:"amount"`
	Description        string    `json:"description"`
	Status_Transaction string    `json:"status_transaction"`
	Status_Poin        string    `json:"status_poin"`
}

type AuthStore struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResponseLoginStore struct {
	Store Store  `json:"store"`
	Token string `json:"token"`
}

type Store struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Store    string `json:"store"`
	Alamat   string `json:"alamat"`
}
type InputPoin struct {
	Customer_id int `json:"customer_id" validate:"required"`
	Store_id    int `json:"store_id" validate:"required"`
	Amount      int `json:"amount" validate:"required"`
}
type ClaimsMitra struct {
	ID    int
	Email string
	Store bool
	jwt.StandardClaims
}

type Customers struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	Password  string `json:"password"`
	No_hp     string `json:"no_hp"`
	Poin      int    `json:"poin"`
	Pin       int    `json:"pin"`
}
