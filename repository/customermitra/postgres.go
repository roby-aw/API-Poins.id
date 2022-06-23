package customermitra

import (
	"api-redeem-point/business/customermitra"
	"api-redeem-point/utils"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursement"
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

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (repo *PosgresRepository) GetCustomersByID(id int) (*customermitra.Customers, error) {
	var data customermitra.Customers
	err := repo.db.Model(&customermitra.Customers{}).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *PosgresRepository) SignCustomer(login *customermitra.AuthLogin) (*customermitra.ResponseLogin, error) {
	var Customer *customermitra.Customers
	err := repo.db.Where("email = ?", login.Email).First(&Customer).Error
	if Customer.Email == "" {
		err = errors.New("email salah")
		return nil, err
	}
	err = VerifyPassword(Customer.Password, login.Password)
	if err != nil {
		err = errors.New("Password salah")
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
	password, err := Hash(Data.Password)
	var Customer customermitra.Customers
	Customer.Email = Data.Email
	Customer.Fullname = Data.Fullname
	Customer.Password = string(password)
	Customer.No_hp = Data.No_hp
	Customer.Pin = Data.Pin
	err = repo.db.Where("email = ?", Data.Email).First(&Customer).Error
	if err == nil {
		err = errors.New("email sudah digunakan")
		return nil, err
	}
	fmt.Println(Customer)
	repo.db.Create(&Customer)

	return Data, nil
}

func (repo *PosgresRepository) UpdateCustomer(Data *customermitra.UpdateCustomer) (*customermitra.UpdateCustomer, error) {
	err := repo.db.Model(&customermitra.Customers{}).Where("ID = ?", Data.ID).Updates(customermitra.Customers{Email: Data.Email, Fullname: Data.Name, No_hp: Data.No_hp}).Error
	if err != nil {
		return nil, err
	}
	return Data, nil
}

func (repo *PosgresRepository) HistoryCustomer(id int) ([]customermitra.History, error) {
	var DataHistory []customermitra.History_Transaction
	err := repo.db.Where("Customer_id = ? AND Status_Poin = ?", id, "OUT").Order("created_at desc").Find(&DataHistory).Error
	if err != nil {
		return nil, err
	}
	var History []customermitra.History
	for _, v := range DataHistory {
		var tmpHistory customermitra.History
		tmpHistory.ID = int(v.ID)
		tmpHistory.ID_Transaction = v.ID_Transaction
		tmpHistory.CreatedAt = v.CreatedAt
		tmpHistory.Transaction_type = v.Transaction_type
		tmpHistory.Status_Transaction = v.Status_Transaction
		History = append(History, tmpHistory)
	}
	return History, nil
}

func (repo *PosgresRepository) DetailHistoryCustomer(idtransaction string) (*customermitra.DetailHistory, error) {
	var transaction customermitra.History_Transaction
	err := repo.db.Where("ID_Transaction = ?", idtransaction).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	DetailHistory := customermitra.DetailHistory{
		ID_Transaction:     transaction.ID_Transaction,
		Transaction_type:   transaction.Transaction_type,
		CreatedAt:          transaction.CreatedAt,
		Bank_Provider:      transaction.Bank_Provider,
		Nomor:              transaction.Nomor,
		Amount:             transaction.Amount,
		Poin_account:       transaction.Poin_Account,
		Poin_redeem:        transaction.Poin_Redeem,
		Description:        transaction.Description,
		Status_Transaction: transaction.Status_Transaction,
	}
	return &DetailHistory, nil
}

func (repo *PosgresRepository) ClaimPulsa(Data *customermitra.RedeemPulsaData) error {
	var Customers customermitra.Customers
	err := repo.db.Where("ID = ?", Data.Customer_id).First(&Customers).Error
	if err != nil {
		return err
	}
	fmt.Println(Customers)
	hasil := Customers.Poin - Data.Poin_redeem
	err = repo.db.Model(&customermitra.Customers{}).Where("ID = ?", Data.Customer_id).Updates(&customermitra.Customers{Poin: hasil}).Error
	if err != nil {
		return err
	}
	fmt.Println(Customers)
	random := utils.Randomstring()
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
	var tmpCustomer customermitra.Customers
	err := repo.db.Where("ID = ?", Data.Customer_id).First(&tmpCustomer).Error
	if err != nil {
		return err
	}
	random := utils.Randomstring()
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

func (repo *PosgresRepository) GetOrderEmoney(emoney *customermitra.InputTransactionBankEmoney) (*customermitra.InputTransactionBankEmoney, error) {
	random := utils.Randomstring()
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
		AccountHolderName: emoney.AN_Rekening,
		AccountNumber:     inputdata.Nomor,
		Description:       "Redeem Emoney" + " - " + inputdata.ID_Transaction,
		Amount:            float64(emoney.Amount),
	}
	_, err := disbursement.Create(&createData)
	if err != nil {
		return nil, err
	}
	errdb := repo.db.Create(&inputdata).Error
	if errdb != nil {
		return nil, errdb
	}
	return emoney, nil
}

func (repo *PosgresRepository) ClaimBank(emoney *customermitra.InputTransactionBankEmoney) (*customermitra.InputTransactionBankEmoney, error) {
	random := utils.Randomstring()
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
