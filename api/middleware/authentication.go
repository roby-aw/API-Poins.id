package middleware

import (
	"api-redeem-point/business/admin"
	"api-redeem-point/business/customermitra"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CustomerSetupAuthenticationJWT() echo.MiddlewareFunc {
	SECRET_KEY := os.Getenv("SECRET_JWT")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		Claims:        &customermitra.Claims{},
		SigningKey:    []byte(SECRET_KEY),
	})
}

func AdminSetupAuthenticationJWT() echo.MiddlewareFunc {
	SECRET_KEY := os.Getenv("SECRET_JWT")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		Claims:        &admin.Claims{Role: "Admin"},
		SigningKey:    []byte(SECRET_KEY),
	})
}
