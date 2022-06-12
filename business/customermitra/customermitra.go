package customermitra

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type History_Transaction struct {
	gorm.Model
	ID_Transaction     string `json:"transaction_id" gorm:"primaryKey;autoIncrement:false"`
	Customer_id        int    `json:"customer_id"`
	Mitra_id           int    `json:"mitra_id"`
	Transaction_type   string `json:"transaction_type"`
	Bank_Provider      string `json:"bank_provider" gorm:"size:255"`
	Nomor              string `json:"nomor" gorm:"size:20"`
	Poin_Account       int    `json:"poin_account"`
	Poin_Redeem        int    `json:"poin_redeem"`
	Amount             int    `json:"amount"`
	Description        string `json:"description" gorm:"size:255"`
	Status_Transaction string `json:"status_transaction" gorm:"size:255"`
	Status_Poin        string `json:"status_poin" gorm:"size:10"`
}

type Store struct {
	gorm.Model
	Email    string `json:"email" gorm:"size:255"`
	Password string `json:"password" gorm:"size:255"`
	Store    string `json:"store" gorm:"size:255"`
	Alamat   string `json:"alamat" gorm:"size:255"`
}

type Customer struct {
	gorm.Model
	Email    string `json:"email" gorm:"primaryKey;autoIncrement:false"`
	Fullname string `json:"fullname" gorm:"size:255"`
	Password string `json:"password" gorm:"size:255"`
	No_hp    string `json:"no_hp" gorm:"size:80"`
	Poin     int    `json:"poin" gorm:"size:50"`
	Pin      int    `json:"pin" gorm:"size:50"`
}

type StockProduct struct {
	gorm.Model
	Product string `json:"product" gorm:"size:100"`
	Balance int    `json:"balance" gorm:"size:100"`
}

type Admin struct {
	gorm.Model
	Email    string `json:"email" gorm:"size:255"`
	Fullname string `json:"fullname" gorm:"size:255"`
	Password string `json:"password" gorm:"size:255"`
	No_hp    string `json:"no_hp" gorm:"size:80"`
}

type RegisterCustomer struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required"`
	Password string `json:"password" validate:"required"`
	No_hp    string `json:"no_hp" validate:"required"`
	Pin      int    `json:"pin" validate:"required"`
}

type Login struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Password string `json:"password"`
	Poin     int    `json:"poin"`
	Pin      int    `json:"pin"`
}

type ResponseLogin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Poin     int    `json:"poin"`
	Pin      int    `json:"pin"`
	Token    string `json:"token"`
}

type AuthLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type History struct {
	ID             int       `json:"id"`
	Tipe_transaksi string    `json:"tipe_transaksi"`
	Tanggal        time.Time `json:"tanggal"`
	Status         string    `json:"status"`
}

type RedeemPulsaData struct {
	Customer_id   int    `json:"customer_id" validate:"required"`
	Bank_Provider string `json:"bank_provider" validate:"required"`
	Nomor         string `json:"nomor" validate:"required"`
	Poin_account  int    `json:"poin_account" validate:"required"`
	Poin_redeem   int    `json:"poin_redeem" validate:"required"`
	Amount        int    `json:"amount" validate:"required"`
}

type UpdateCustomer struct {
	ID    int    `json:"id" validate:"required"`
	Name  string `json:"name"`
	Email string `json:"email"`
	No_hp string `json:"no_hp"`
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
	Customer_id   int    `json:"customer_id"`
	Bank_Provider string `json:"bank_provider"`
	AN_Rekening   string `json:"an_rekening"`
	Nomor         string `json:"nomor" validate:"required"`
	Amount        int    `json:"amount"`
	Poin_account  int    `json:"poin_account" validate:"required"`
	Poin_redeem   int    `json:"poin_redeem" validate:"required"`
}

type Claims struct {
	ID    int
	Email string
	jwt.StandardClaims
}
