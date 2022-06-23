package api

import (
	"api-redeem-point/api/admin"
	"api-redeem-point/api/customermitra"
	"api-redeem-point/api/middleware"

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
	c.GET("/history/:idcustomer", controller.CustomerMitraController.HistoryCustomer, middleware.CustomerSetupAuthenticationJWT())
	c.GET("/dethistory/:idtransaction", controller.CustomerMitraController.DetailHistoryCustomer, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/pulsa", controller.CustomerMitraController.OrderPulsa, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/paketdata", controller.CustomerMitraController.OrderPaketData, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/cashout", controller.CustomerMitraController.OrderCashout, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/emoney", controller.CustomerMitraController.OrderEmoney, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/callback", controller.CustomerMitraController.CallbackXendit)
	c.GET("/test", controller.AdminControlller.Test)
	//admin
	g := c.Group("/admin")
	g.POST("/login", controller.AdminControlller.LoginAdmin)
	g.POST("", controller.AdminControlller.CreateAdmin)
	g.GET("", controller.AdminControlller.TransactionPending, middleware.AdminSetupAuthenticationJWT())
	g.POST("/approve/:idtransaction", controller.AdminControlller.ApproveTransaction, middleware.AdminSetupAuthenticationJWT())
	g.GET("/customer", controller.AdminControlller.FindCustomers, middleware.AdminSetupAuthenticationJWT())
	g.GET("/history", controller.AdminControlller.FindHistoryCustomers, middleware.AdminSetupAuthenticationJWT())
	g.GET("/transaction", controller.AdminControlller.TransactionByDate)
	g.PUT("/customer", controller.AdminControlller.UpdateCustomer)
	g.PUT("/customer/point", controller.AdminControlller.UpdateCustomerPoint)
	g.GET("/stock", controller.AdminControlller.StockProduct)
	g.PUT("/stock", controller.AdminControlller.UpdateStock)
}
