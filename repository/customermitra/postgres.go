package customermitra

import (
	"api-redeem-point/business/customermitra"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursement"
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

func (repo *PosgresRepository) SignCustomer(login *customermitra.AuthLogin) (*customermitra.ResponseLogin, error) {
	var Customer *customermitra.Customer
	err := repo.db.Where("email = ? AND password = ?", login.Email, login.Password).First(&Customer).Error
	if Customer.Email == "" {
		err = errors.New("email atau password salah")
		return nil, err
	}
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &customermitra.Claims{
		ID:    int(Customer.ID),
		Email: Customer.Email,
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
	var response customermitra.ResponseLogin
	response.ID = int(Customer.ID)
	response.Email = Customer.Email
	response.Password = Customer.Password
	response.Pin = Customer.Pin
	response.Poin = Customer.Poin
	response.Token = token_jwt
	return &response, err
}

func (repo *PosgresRepository) InsertCustomer(Data *customermitra.RegisterCustomer) (*customermitra.RegisterCustomer, error) {
	var Customer customermitra.Customer
	Customer.Email = Data.Email
	Customer.Fullname = Data.Fullname
	Customer.Password = Data.Password
	Customer.No_hp = Data.No_hp
	Customer.Pin = Data.Pin
	err := repo.db.Where("email = ?", Data.Email).First(&Customer).Error
	if err == nil {
		err = errors.New("email sudah digunakan")
		return nil, err
	}
	fmt.Println(Customer)
	repo.db.Create(&Customer)

	return Data, nil
}

func (repo *PosgresRepository) UpdateCustomer(Data *customermitra.UpdateCustomer) (*customermitra.UpdateCustomer, error) {
	err := repo.db.Model(&customermitra.Customer{}).Where("ID = ?", Data.ID).Updates(customermitra.Customer{Email: Data.Email, Fullname: Data.Name, No_hp: Data.No_hp}).Error
	if err != nil {
		return nil, err
	}
	return Data, nil
}

func (repo *PosgresRepository) ClaimPulsa(Data *customermitra.RedeemPulsaData) error {
	var tmpCustomer customermitra.Customer
	err := repo.db.Where("ID = ?", Data.Customer_id).First(&tmpCustomer).Error
	if err != nil {
		return err
	}
	random := randomstring()
	var tmpHistory customermitra.History_Transaction
	repo.db.Where("ID_Transaction = ?", "P"+random).First(&tmpHistory)
	if tmpHistory.ID_Transaction != "" {
		inthasil, _ := strconv.Atoi(random)
		inthasil = inthasil + 1
		random = strconv.Itoa(inthasil)
	}
	err = nil
	transaction := customermitra.History_Transaction{
		Customer_id:        Data.Customer_id,
		ID_Transaction:     "P" + random,
		Transaction_type:   "Redeem Pulsa",
		Bank_Provider:      Data.Bank_Provider,
		Nomor:              Data.Nomor,
		Poin_Account:       Data.Poin_account,
		Poin_Redeem:        Data.Poin_redeem,
		Amount:             Data.Amount,
		Description:        Data.Bank_Provider + " - " + strconv.Itoa(Data.Amount),
		Status_Transaction: "PENDING",
		Status_Poin:        "OUT",
	}
	err = repo.db.Create(&transaction).Error
	return err
}

func (repo *PosgresRepository) ClaimPaketData(Data *customermitra.RedeemPulsaData) error {
	var tmpCustomer customermitra.Customer
	err := repo.db.Where("ID = ?", Data.Customer_id).First(&tmpCustomer).Error
	if err != nil {
		return err
	}
	random := randomstring()
	var tmpHistory customermitra.History_Transaction
	repo.db.Where("ID_Transaction = ?", "P"+random).First(&tmpHistory)
	if tmpHistory.ID_Transaction != "" {
		inthasil, _ := strconv.Atoi(random)
		inthasil = inthasil + 1
		random = strconv.Itoa(inthasil)
	}
	err = nil
	transaction := customermitra.History_Transaction{
		Customer_id:        Data.Customer_id,
		ID_Transaction:     "PD" + random,
		Transaction_type:   "Redeem Paket Data",
		Bank_Provider:      Data.Bank_Provider,
		Nomor:              Data.Nomor,
		Poin_Account:       Data.Poin_account,
		Poin_Redeem:        Data.Poin_redeem,
		Amount:             Data.Amount,
		Description:        Data.Bank_Provider + " - " + strconv.Itoa(Data.Amount) + "GB",
		Status_Transaction: "PENDING",
		Status_Poin:        "OUT",
	}
	err = repo.db.Create(&transaction).Error
	return err
}

func (repo *PosgresRepository) TakeCallback(data *customermitra.Disbursement) (*customermitra.Disbursement, error) {
	var TransactionBank customermitra.History_Transaction
	TransactionBank.Status_Transaction = data.Status

	err := repo.db.Model(TransactionBank).Where("ID_Transaction = ?", data.ExternalID).Updates(TransactionBank).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *PosgresRepository) GetOrderEmoney(emoney *customermitra.InputTransactionEmoney) (*customermitra.InputTransactionEmoney, error) {
	random := randomstring()
	inputdata := customermitra.History_Transaction{
		ID_Transaction:     "EM" + random,
		Transaction_type:   "Redeem Emoney",
		Customer_id:        emoney.Customer_id,
		Bank_Provider:      emoney.Bank_Provider,
		Nomor:              emoney.Nomor,
		Amount:             emoney.Amount,
		Poin_Account:       emoney.Poin_account,
		Poin_Redeem:        emoney.Poin_redeem,
		Description:        emoney.Bank_Provider + " - " + strconv.Itoa(emoney.Amount),
		Status_Transaction: "PENDING",
		Status_Poin:        "OUT",
	}
	xendit.Opt.SecretKey = "xnd_development_cUiYsYw0nFqaykCMXpl3cqoxlIy7zciDRVaTHemLUUXhh3iKKILDJvbYKo8U9t"

	createData := disbursement.CreateParams{
		IdempotencyKey:    "disbursement" + time.Now().String(),
		ExternalID:        inputdata.ID_Transaction,
		BankCode:          inputdata.Bank_Provider,
		AccountHolderName: "",
		AccountNumber:     inputdata.Nomor,
		Description:       "Redeem Emoney" + " - " + inputdata.ID_Transaction,
		Amount:            float64(emoney.Amount),
	}
	fmt.Println(createData)
	resp, err := disbursement.Create(&createData)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	errdb := repo.db.Create(&inputdata).Error
	if errdb != nil {
		return nil, errdb
	}
	return emoney, nil
}

func (repo *PosgresRepository) ClaimBank(emoney *customermitra.InputTransactionBank) (*customermitra.InputTransactionBank, error) {
	random := randomstring()
	inputdata := customermitra.History_Transaction{
		ID_Transaction:     "EM" + random,
		Transaction_type:   "Redeem Bank",
		Customer_id:        emoney.Customer_id,
		Bank_Provider:      emoney.Bank_Provider,
		Nomor:              emoney.Nomor,
		Amount:             emoney.Amount,
		Poin_Account:       emoney.Poin_account,
		Poin_Redeem:        emoney.Poin_redeem,
		Description:        emoney.Bank_Provider + " - " + emoney.AN_Rekening,
		Status_Transaction: "PENDING",
		Status_Poin:        "OUT",
	}
	xendit.Opt.SecretKey = "xnd_development_cUiYsYw0nFqaykCMXpl3cqoxlIy7zciDRVaTHemLUUXhh3iKKILDJvbYKo8U9t"

	createData := disbursement.CreateParams{
		IdempotencyKey:    "disbursement" + time.Now().String(),
		ExternalID:        inputdata.ID_Transaction,
		BankCode:          inputdata.Bank_Provider,
		AccountHolderName: emoney.AN_Rekening,
		AccountNumber:     inputdata.Nomor,
		Description:       "Redeem Emoney" + " - " + inputdata.ID_Transaction,
		Amount:            float64(emoney.Amount),
	}
	fmt.Println(createData)
	resp, err := disbursement.Create(&createData)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	errdb := repo.db.Create(&inputdata).Error
	if errdb != nil {
		return nil, errdb
	}
	return emoney, nil
}

func randomstring() string {
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Int()
	angka := []rune(fmt.Sprintf("%d", random))
	b := make([]rune, 8)
	for i := range b {
		b[i] = angka[rand.Intn(len(angka))]
	}
	hasil := string(b)
	inthasil, _ := strconv.Atoi(hasil)
	strhasil := strconv.Itoa(inthasil)
	return strhasil
}
