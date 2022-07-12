package store

import (
	storeBussiness "api-redeem-point/business/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service storeBussiness.Service
}

func NewController(service storeBussiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Login Store
// @description Register Store for Admin
// @tags Store
// @Accept json
// @Produce json
// @Param LoginStore body store.AuthStore true "LoginStore"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /store/login [post]
func (Controller *Controller) LoginStore(c echo.Context) error {
	var req storeBussiness.AuthStore
	c.Bind(&req)
	result, err := Controller.service.LoginStore(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success login store",
		"result":   result,
	})
}

// Create godoc
// @Summary Input Poin Store
// @description Input Poin Customer for Store
// @tags Store
// @Accept json
// @Produce json
// @Param InputPoinStore body store.InputPoin true "InputPoinStore"
// @Success 200	{object} response.Result
// @Failure 400 {object} response.Error
// @Router /store/poin [post]
func (Controller *Controller) InputPoinStore(c echo.Context) error {
	var req storeBussiness.InputPoin
	c.Bind(&req)
	result, err := Controller.service.InputPoin(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success input poin",
		"Add poin": result,
	})
}
