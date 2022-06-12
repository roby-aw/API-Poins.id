package api

import (
	"api-redeem-point/api/admin"
	"api-redeem-point/api/customermitra"

	//auth "api-redeem-point/api/middleware"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	AdminControlller        *admin.Controller
	CustomerMitraController *customermitra.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	c := e.Group("/v1")
	c.POST("/customer/register", controller.CustomerMitraController.Register)
	c.POST("/customer", controller.CustomerMitraController.Login)
	c.PUT("/customer", controller.CustomerMitraController.UpdateCustomer)
	c.POST("/pulsa", controller.CustomerMitraController.OrderPulsa)
	c.POST("/paketdata", controller.CustomerMitraController.OrderPaketData)
	c.POST("/emoney", controller.CustomerMitraController.OrderEmoney)
	//admin
	// g := e.Group("/admin")
	// g.POST("", controller.AdminControlller.CreateAdmin)
	// g.POST("/token", controller.AdminControlller.GetToken)
	// g.PUT("/:id", controller.AdminControlller.UpdateAdmin)
	// g.GET("/:id", controller.AdminControlller.GetAdminByID)
	// // admin using jwt
	// g.DELETE("/:id", controller.AdminControlller.DeleteAdmin, auth.SetupAuthenticationJWT())
	// g.GET("", controller.AdminControlller.GetAdmins, auth.SetupAuthenticationJWT())
	//callback
	// c.POST("/callback", controller.CustomerMitraController.CallbackXendit)
	// //order
	// c.POST("/order/emoney", controller.CustomerMitraController.OrderEmoney)
}
