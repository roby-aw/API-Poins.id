package user

import (
	"api-redeem-point/business/user"
	"fmt"
	"math/rand"
	"strconv"
	"time"

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

func (repo *PosgresRepository) FindFoodByName(name string) (foods []user.Food, err error) {
	fmt.Println("repo jalan")
	result := repo.db.Where("name LIKE ?", "%"+name+"%").Find(&foods)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(foods) == 0 {
		return nil, fmt.Errorf("data tidak ditemukan")
	}
	return foods, nil
}

func (repo *PosgresRepository) TakeCallback(data *user.Disbursement) (*user.Disbursement, error) {
	var TransactionBank *user.TransactionBank
	if data.Status == "COMPLETED" {
		TransactionBank.Status = data.Status
	}

	err := repo.db.Model(TransactionBank).Where("ID_Transaction = ?", data.ExternalID).Updates(TransactionBank).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *PosgresRepository) GetOrderEmoney(emoney *user.InputTransactionBank) (*user.InputTransactionBank, error) {
	var err error
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Int()
	fmt.Println(random)
	angka := []rune(fmt.Sprintf("%d", random))
	b := make([]rune, 8)
	for i := range b {
		b[i] = angka[rand.Intn(len(angka))]
	}
	hasil := string(b)
	inthasil, _ := strconv.Atoi(hasil)
	strhasil := strconv.Itoa(inthasil)
	var databank user.TransactionBank

	repo.db.Where("ID_Transaction = ?", hasil).First(databank)
	if databank.ID_Transaction != "" {
		inthasil = inthasil + 1
		strhasil = strconv.Itoa(inthasil)
	}
	inputdata := user.TransactionBank{
		ID_Transaction:    "EM" + strhasil,
		ID_User:           emoney.ID_User,
		Jenis_transaction: emoney.Jenis_transaction,
		Nama_bank:         emoney.Nama_bank,
		AN_Bank:           emoney.AN_Bank,
		No_rekening:       emoney.No_rekening,
		Amount:            emoney.Amount,
		Status:            "PENDING",
	}
	defer repo.db.Create(&inputdata)
	xendit.Opt.SecretKey = "xnd_development_cUiYsYw0nFqaykCMXpl3cqoxlIy7zciDRVaTHemLUUXhh3iKKILDJvbYKo8U9t"

	createData := disbursement.CreateParams{
		IdempotencyKey:    "disbursement" + time.Now().String(),
		ExternalID:        "EM" + strhasil,
		BankCode:          emoney.Nama_bank,
		AccountHolderName: emoney.AN_Bank,
		AccountNumber:     emoney.No_rekening,
		Description:       "ID" + strhasil,
		Amount:            float64(emoney.Amount),
	}

	resp, err := disbursement.Create(&createData)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	return emoney, nil
}
