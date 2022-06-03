package user

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name        string `json:"name"`
	Picture_url string `json:"picture_url"`
	City        string `json:"city"`
	Price       string `json:"price"`
	Open_time   string `json:"open_time"`
	Latitude    string `json:"latitude"`
	Longtitude  string `json:"Longtitude"`
	Rating      string `json:"rating"`
	Visited     int    `json:"visited"`
}

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	No_hp    string `json:"no_hp"`
	Password string `json:"password"`
	Pin      int    `json:"pin"`
}

type Login struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Password string `json:"password"`
	Poin     int    `json:"poin"`
	Pin      int    `json:"pin"`
}
type AuthLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type History struct {
	ID             int       `json:"id"`
	Tipe_transaksi string    `json:"tipe_transaksi"`
	Tanggal        time.Time `json:"tanggal"`
	Status         string    `json:"status"`
}

type DetailTransaction struct {
	ID                int    `json:"id"`
	Jenis_transaction string `json:"jenis_transaction"`
	Nama_bank         string `json:"nama_bank"`
	No_rekening       int    `json:"no_rekening"`
	Poin_account      int    `json:"poin_account"`
	Poin_redeem       int    `json:"poin_redeem"`
}

type ProductCashout struct {
	ID    int `json:"id"`
	Harga int `json:"harga"`
	Poin  int `json:"poin"`
}
type ProductEmoney struct {
	ID    int `json:"id"`
	Harga int `json:"harga"`
	Poin  int `json:"poin"`
}
type ProductPulsa struct {
	ID    int `json:"id"`
	Harga int `json:"harga"`
	Poin  int `json:"poin"`
}

type ProductPaketData struct {
	ID    int    `json:"id"`
	Nama  string `json:"Internet"`
	Kuota string `json:"kuota"`
	Harga int    `json:"harga"`
	Poin  int    `json:"poin"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	No_hp    string `json:"no_hp"`
	Password string `json:"password"`
	Pin      int    `json:"pin"`
}

type Bank struct {
	BankCode    string `json:"bankcode"`
	No_rekening string `json:"no_rekening"`
	Amount      int    `json:"amount"`
}

type Disbursement struct {
	UserID                  string `json:"user_id"`
	ExternalID              string `json:"external_id"`
	Amount                  int    `json:"amount"`
	BankCode                string `json:"bank_code"`
	AccountHolderName       string `json:"account_holder_name"`
	DisbursementDescription string `json:"disbursement_description"`
	Status                  string `json:"status"`
	ID                      string `json:"id"`
}

type TransactionBank struct {
	ID                uint64 `gorm:"primaryKey"`
	ID_Transaction    string `json:"id_transaction"`
	ID_User           string `json:"id_user"`
	Jenis_transaction string `json:"jenis_transaction" validate:"required"`
	Nama_bank         string `json:"nama_bank" validate:"required"`
	AN_Bank           string `json:"AN_Bank" validate:"required"`
	No_rekening       string `json:"no_rekening" validate:"required"`
	Amount            int    `json:"amount"`
	Status            string `json:"status"`
}

type InputTransactionBank struct {
	ID_User           string `json:"id_user"`
	Jenis_transaction string `json:"jenis_transaction" validate:"required"`
	Nama_bank         string `json:"nama_bank" validate:"required"`
	AN_Bank           string `json:"AN_Bank" validate:"required"`
	No_rekening       string `json:"no_rekening" validate:"required"`
	Amount            int    `json:"amount"`
}
