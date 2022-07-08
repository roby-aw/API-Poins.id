package api

import (
	"api-redeem-point/api/admin"
	"api-redeem-point/api/customer"
	"api-redeem-point/api/middleware"
	"api-redeem-point/api/store"

	//auth "api-redeem-point/api/middleware"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	AdminControlller   *admin.Controller
	CustomerController *customer.Controller
	StoreController    *store.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	c := e.Group("/v1")
	c.POST("/customer/register", controller.CustomerController.Register)
	c.POST("/customer", controller.CustomerController.Login)
	c.PUT("/customer", controller.CustomerController.UpdateCustomer, middleware.CustomerSetupAuthenticationJWT())
	c.GET("/customer/:id", controller.CustomerController.FindCustomersByID, middleware.CustomerSetupAuthenticationJWT())
	c.GET("/history", controller.CustomerController.HistoryCustomer, middleware.CustomerSetupAuthenticationJWT())
	c.GET("/dethistory/:idtransaction", controller.CustomerController.DetailHistoryCustomer, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/pulsa", controller.CustomerController.OrderPulsa, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/paketdata", controller.CustomerController.OrderPaketData, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/cashout", controller.CustomerController.OrderCashout, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/emoney", controller.CustomerController.OrderEmoney, middleware.CustomerSetupAuthenticationJWT())
	c.POST("/callback", controller.CustomerController.CallbackXendit)
	//admin
	g := c.Group("/admin")
	g.POST("/login", controller.AdminControlller.LoginAdmin)
	g.POST("", controller.AdminControlller.CreateAdmin)
	g.GET("", controller.AdminControlller.Dashboard)
	g.GET("/:id", controller.AdminControlller.FindAdminByID)
	g.GET("/transaction/pending", controller.AdminControlller.TransactionPending)
	g.POST("/approve/:idtransaction", controller.AdminControlller.ApproveTransaction)
	g.GET("/history", controller.AdminControlller.FindHistoryCustomers)
	g.GET("/customer", controller.AdminControlller.FindCustomers)
	g.PUT("/customer", controller.AdminControlller.UpdateCustomer)
	g.DELETE("/customer", controller.AdminControlller.DeleteCustomer)
	g.PUT("/customer/point", controller.AdminControlller.UpdateCustomerPoint)
	g.GET("/stock", controller.AdminControlller.StockProduct)
	g.PUT("/stock", controller.AdminControlller.UpdateStock)
	g.GET("/historystore", controller.AdminControlller.HistoryStore)
	g.DELETE("/store", controller.AdminControlller.DeleteStore)
	g.GET("/store", controller.AdminControlller.GetStore)
	g.PUT("/store", controller.AdminControlller.UpdateStore)
	s := c.Group("/store")
	s.POST("", controller.CustomerController.RegisterStore)
	s.POST("/login", controller.StoreController.LoginStore)
	s.POST("/poin", controller.StoreController.InputPoinStore, middleware.StoreSetupAuthenticationJWT())
}
