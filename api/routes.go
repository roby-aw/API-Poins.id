package api

import (
	"api-redeem-point/api/admin"
	auth "api-redeem-point/api/middleware"
	"api-redeem-point/api/user"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	AdminControlller *admin.Controller
	UserController   *user.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	c := e.Group("/v1")
	c.POST("/login", controller.UserController.Login)
	//admin
	g := e.Group("/admin")
	g.POST("", controller.AdminControlller.CreateAdmin)
	g.POST("/token", controller.AdminControlller.GetToken)
	g.PUT("/:id", controller.AdminControlller.UpdateAdmin)
	g.GET("/:id", controller.AdminControlller.GetAdminByID)
	// admin using jwt
	g.DELETE("/:id", controller.AdminControlller.DeleteAdmin, auth.SetupAuthenticationJWT())
	g.GET("", controller.AdminControlller.GetAdmins, auth.SetupAuthenticationJWT())
	//callback
	c.POST("/callback", controller.UserController.CallbackXendit)
	//order
	c.POST("/order/emoney", controller.UserController.OrderEmoney)
}
