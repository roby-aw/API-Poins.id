package customermitra

import (
	customermitraBussiness "api-redeem-point/business/customermitra"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service customermitraBussiness.Service
}

func NewController(service customermitraBussiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Login
// @description Login Customer
// @tags Customer
// @Accept json
// @Produce json
// @Param Customer body customermitra.AuthLogin true "Customer"
// @Success 200 {object} customermitra.Login
// @Router /v1/login [post]
func (Controller *Controller) Login(c echo.Context) error {
	var req customermitraBussiness.AuthLogin
	var err error
	c.Bind(&req)
	result, err := Controller.service.LoginCustomer(&req)
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
// @Summary Register
// @description Register customer
// @tags customer
// @Accept json
// @Produce json
// @Param Registercustomer body customer.Register true "Register"
// @Success 200 {object} customermitra.Register
// @Router /register [post]
func (Controller *Controller) Register(c echo.Context) error {
	var req customermitraBussiness.RegisterCustomer
	c.Bind(&req)
	result, err := Controller.service.CreateCustomer(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
		"result":   result,
	})
}

// // Create godoc
// // @Summary Updatecustomer
// // @description Updatecustomer
// // @tags customer
// // @Accept json
// // @Produce json
// // @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// // @Success 200 {object} customermitra.Register
// // @Router /v1/account [put]
func (Controller *Controller) UpdateCustomer(c echo.Context) error {
	var req customermitraBussiness.UpdateCustomer
	c.Bind(&req)
	result, err := Controller.service.UpdateCustomer(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success update customer",
		"result":   result,
	})
}

// Create godoc
// @Summary History
// @description History customer
// @tags customer
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param idcustomer path int true "id customer"
// @Success 200 {object} customermitra.History
// @Router /v1/history/{idcustomer} [get]
// func (Controller *Controller) History(c echo.Context) error {
// 	var err error
// 	//idcustomer, _ := strconv.Atoi(c.Param("idcustomer"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"code":     400,
// 			"messages": err.Error(),
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"code":     200,
// 		"messages": "success get history",
// 		//"result":   arr,
// 	})
// }

// Create godoc
// @Summary Detail History/transaction
// @description History/transaction customer
// @tags customer
// @Accept json
// @Produce json
// @Param id path int true "id detail history"
// @Success 200 {object} customer.DetailTransaction
// @Router /v1/detailhistory/{id} [get]
// func (Controller *Controller) DetailTransaction(c echo.Context) error {
// 	var err error
// 	//id, _ := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"code":     400,
// 			"messages": err.Error(),
// 		})
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"code":     200,
// 		"messages": "success get detail transaction",
// 		// "result":   detailtransaction,
// 	})
// }

// Create godoc
// @Summary Order Emoney/Cashout
// @description Emoney customer
// @tags customerOrder
// @Accept json
// @Produce json
// @Param InputDataCashout body customermitra.InputTransactionBank true "inputdataemoney"
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/emoney [post]
func (Controller *Controller) OrderEmoney(c echo.Context) error {
	emoney := customermitraBussiness.InputTransactionBank{}
	c.Bind(&emoney)
	result, err := Controller.service.ToOrderEmoney(&emoney)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"Error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success order emoney",
		"result":   result,
	})
}

// Create godoc
// @Summary Order Emoney/Cashout
// @description Emoney customer
// @tags customerOrder
// @Accept json
// @Produce json
// @Param InputDataCashout body customermitra.InputTransactionBank true "inputdataemoney"
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/emoney [post]
func (Controller *Controller) OrderCashout(c echo.Context) error {
	req := customermitraBussiness.InputTransactionBank{}
	c.Bind(&req)
	result, err := Controller.service.RedeemBank(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"Error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success order cashout",
		"result":   result,
	})
}

// Create godoc
// @Summary Order Pulsa
// @description Pulsa customer
// @tags customerOrder
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/pulsa [post]
func (Controller *Controller) OrderPulsa(c echo.Context) error {
	var req customermitraBussiness.RedeemPulsaData
	c.Bind(&req)
	err := Controller.service.RedeemPulsa(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success redeem pulsa",
	})
}

// Create godoc
// @Summary Order PaketData
// @description PaketData customer
// @tags customerOrder
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/paketdata [post]
func (Controller *Controller) OrderPaketData(c echo.Context) error {
	var req customermitraBussiness.RedeemPulsaData
	c.Bind(&req)
	fmt.Println(req)
	err := Controller.service.RedeemPaketData(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}

func (Controller *Controller) CallbackXendit(c echo.Context) error {
	req := c.Request()
	decoder := json.NewDecoder(req.Body)
	disbursermentData := customermitraBussiness.Disbursement{}
	err := decoder.Decode(&disbursermentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	defer req.Body.Close()
	disbursement, _ := json.Marshal(disbursermentData)
	var resbank customermitraBussiness.Disbursement
	json.Unmarshal(disbursement, &resbank)
	responseWriter := c.Response().Writer
	responseWriter.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	responseWriter.WriteHeader(200)
	fmt.Println(resbank)
	databank, err := Controller.service.GetCallback(&resbank)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"Error":   err.Error(),
		})
	}
	fmt.Println(resbank)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": databank,
	})
}
