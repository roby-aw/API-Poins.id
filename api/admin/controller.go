package admin

import (
	"api-redeem-point/business/admin"
	adminBusiness "api-redeem-point/business/admin"
	"api-redeem-point/business/customermitra"
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
// @tags admin
// @Accept json
// @Produce json
// @Param admin body admin.AdminSwagger true "admin"
// @Success 201 {object} admin.Admin
// @Failure 400 {object} map[string]interface{}
// @Router /admin [post]
func (Controller *Controller) CreateAdmin(c echo.Context) error {
	admin := customermitra.Admin{}
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
// @Summary Get Token
// @description Get token for admin
// @tags admin
// @Accept json
// @Produce json
// @Param admin body admin.InputAdminToken true "admin"
// @Success 200 {object} map[string]interface{}
// @Failure 400
// @Router /admin/token [post]
func (Controller *Controller) GetToken(c echo.Context) error {
	var request adminBusiness.Admin

	c.Bind(&request)
	token, err := Controller.service.GetToken(&request)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success login",
		"token":   token,
	})
}

// Create godoc
// @Summary Delete Admin
// @description delete data admin
// @tags admin
// @Accept json
// @Produce json
// @Param id path int true "id admin"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]interface{}
// @Failure 400
// @Router /admin/{id} [delete]
func (Controller *Controller) DeleteAdmin(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := Controller.service.DeleteAdmin(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success delete admin",
		"data id ": id,
	})
}

// Create godoc
// @Summary Update Admin
// @description update data admin
// @tags Admin using Token JWT
// @Accept json
// @Produce json
// @Param id path int true "id admin"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param admin body admin.AdminSwagger true "admin"
// @Success 200 {object} map[string]interface{}
// @Failure 400
// @Router /admin/{id} [PUT]
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
