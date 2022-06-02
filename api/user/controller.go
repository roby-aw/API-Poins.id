package user

import (
	userBussiness "api-redeem-point/business/user"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service userBussiness.Service
}

func NewController(service userBussiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Login
// @description Login user
// @tags User
// @Accept json
// @Produce json
// @Param user body user.AuthLogin true "user"
// @Success 200 {object} user.Login
// @Router /v1/login [post]
func (Controller *Controller) Login(c echo.Context) error {
	result := &userBussiness.Login{
		ID:       1,
		Email:    "test@gmail.com",
		Password: "testpassword",
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTE2ODczODh9.dw2WuBIDJcb5dVT8iF_63POdhZFYOq4D1-kZTiCyo7c",
		Poin:     500000,
		Pin:      1234,
	}
	var req userBussiness.AuthLogin
	var err error
	c.Bind(&req)
	if (req.Email != result.Email) || (req.Password != result.Password) {
		err = fmt.Errorf("Email atau password salah")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success login",
		"result":   result,
	})
}

// Create godoc
// @Summary History
// @description History User
// @tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param iduser path int true "id user"
// @Success 200 {object} user.History
// @Router /v1/history/{iduser} [get]
func (Controller *Controller) History(c echo.Context) error {
	History1 := &userBussiness.History{
		ID:             1,
		Tipe_transaksi: "Redeem CashOut",
		Tanggal:        time.Date(2022, 5, 16, 156, 24, 34, 534, time.UTC),
		Status:         "Sukses",
	}
	History2 := &userBussiness.History{
		ID:             5,
		Tipe_transaksi: "Redeem paket data",
		Tanggal:        time.Date(2022, 5, 17, 156, 24, 34, 534, time.UTC),
		Status:         "Sukses",
	}
	History3 := &userBussiness.History{
		ID:             7,
		Tipe_transaksi: "Redeem CashOut",
		Tanggal:        time.Date(2022, 5, 18, 156, 24, 34, 534, time.UTC),
		Status:         "Pending",
	}
	var arr []userBussiness.History
	arr = append(arr, *History1)
	arr = append(arr, *History2)
	arr = append(arr, *History3)
	var err error
	iduser, _ := strconv.Atoi(c.Param("iduser"))
	if iduser != 1 {
		err = fmt.Errorf("iduser tidak ditemukan")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get history",
		"result":   arr,
	})
}

// Create godoc
// @Summary Detail History/transaction
// @description History/transaction User
// @tags User
// @Accept json
// @Produce json
// @Param id path int true "id detail history"
// @Success 200 {object} user.DetailTransaction
// @Router /v1/detailhistory/{id} [get]
func (Controller *Controller) DetailTransaction(c echo.Context) error {
	var err error
	detailtransaction := userBussiness.DetailTransaction{
		ID:                1,
		Jenis_transaction: "Redeem Cashout",
		Nama_bank:         "BNI",
		No_rekening:       12354665,
		Poin_account:      500000,
		Poin_redeem:       100000,
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if id != 1 {
		err = fmt.Errorf("ID tidak ditemukan")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get detail transaction",
		"result":   detailtransaction,
	})
}

// Create godoc
// @Summary Register
// @description Register user
// @tags User
// @Accept json
// @Produce json
// @Param RegisterUser body user.Register true "Register"
// @Success 200 {object} user.Register
// @Router /v1/register [post]
func (Controller *Controller) Register(c echo.Context) error {
	result := &userBussiness.Register{
		Name:     "ininamatest",
		Email:    "test@gmail.com",
		No_hp:    "063542251",
		Password: "testpassword",
		Pin:      1234,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
		"result":   result,
	})
}

// Create godoc
// @Summary Order Cashout
// @description Order Cashout
// @tags UserOrder
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param InputDataCashout body user.Bank true "inputdatacashout"
// @Success 200 {object} user.Bank
// @Router /v1/order/cashout [post]
func (Controller *Controller) OrderCashout(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}

// Create godoc
// @Summary Order Emoney
// @description Emoney user
// @tags UserOrder
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/emoney [post]
func (Controller *Controller) OrderEmoney(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}

// Create godoc
// @Summary Order Pulsa
// @description Pulsa user
// @tags UserOrder
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/pulsa [post]
func (Controller *Controller) OrderPulsa(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}

// Create godoc
// @Summary Order PaketData
// @description PaketData user
// @tags UserOrder
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/paketdata [post]
func (Controller *Controller) OrderPaketData(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}

// Create godoc
// @Summary UpdateUser
// @description UpdateUser
// @tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} user.Register
// @Router /v1/account [put]
func (Controller *Controller) UpdateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}

func (Controller *Controller) CallbackXendit(c echo.Context) error {
	req := c.Request()
	decoder := json.NewDecoder(req.Body)
	disbursermentData := userBussiness.Disbursement{}
	err := decoder.Decode(&disbursermentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	defer req.Body.Close()
	disbursement, _ := json.Marshal(disbursermentData)
	var resbank userBussiness.Disbursement
	json.Unmarshal(disbursement, &resbank)
	responseWriter := c.Response().Writer
	responseWriter.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	responseWriter.WriteHeader(200)
	databank, err := Controller.service.GetCallback(&resbank)
	fmt.Println(resbank)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": databank,
	})
}
