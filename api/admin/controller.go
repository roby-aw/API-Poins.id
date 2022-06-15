package admin

import (
	"api-redeem-point/business/admin"
	adminBusiness "api-redeem-point/business/admin"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service adminBusiness.Service
}

func NewController(service adminBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Create admin
// @description create admin with data
// @tags Admin
// @Accept json
// @Produce json
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Router /admin [post]
func (Controller *Controller) Dashboard(c echo.Context) error {
	result, err := Controller.service.Dashboard()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"Error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get dashboard",
		"result":   result,
	})
}

// Create godoc
// @Summary Create admin
// @description create admin with data
// @tags Admin
// @Accept json
// @Produce json
// @Param admin body admin.RegisterAdmin true "admin"
// @Success 201 {object} admin.RegisterAdmin
// @Failure 400 {object} map[string]interface{}
// @Router /admin [post]
func (Controller *Controller) CreateAdmin(c echo.Context) error {
	admin := adminBusiness.RegisterAdmin{}
	c.Bind(&admin)
	admins, err := Controller.service.CreateAdmin(&admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"Error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success create data",
		"data":     admins,
	})
}

// Create godoc
// @Summary Login admin
// @description Login admin
// @tags Admin
// @Accept json
// @Produce json
// @Param admin body admin.AuthLogin true "admin"
// @Success 200 {object} map[string]interface{}
// @Failure 400
// @Router /admin/login [post]
func (Controller *Controller) LoginAdmin(c echo.Context) error {
	var request adminBusiness.AuthLogin
	c.Bind(&request)
	result, err := Controller.service.LoginAdmin(&request)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success login",
		"result":  result,
	})
}

func (Controller *Controller) UpdateAdmin(c echo.Context) error {
	var admin *admin.Admin
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&admin)
	admin, err := Controller.service.UpdateAdmin(id, admin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success update data admin",
		"data":     admin,
	})
}

// Create godoc
// @Summary Approve Transaction
// @description Approve Transaction
// @tags Admin
// @Accept json
// @Produce json
// @Param transactionid path string true "transaction_id"
// @Success 200
// @Router /admin/approve/{transactionid} [post]
func (Controller *Controller) ApproveTransaction(c echo.Context) error {
	transactionid := c.Param("idtransaction")
	err := Controller.service.ApproveTransaction(transactionid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success approve transaction",
	})
}

func (Controller *Controller) FindCustomers(c echo.Context) error {
	result, err := Controller.service.FindCustomers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get customers",
		"result":   result,
	})
}
