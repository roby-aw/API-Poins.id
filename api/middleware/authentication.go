package middleware

import (
	"api-redeem-point/business/admin"
	"api-redeem-point/business/customermitra"
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
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

func StoreSetupAuthenticationJWT() echo.MiddlewareFunc {
	SECRET_KEY := os.Getenv("SECRET_JWT")
	config := middleware.JWTConfig{
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return SECRET_KEY, nil
			}

			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
			token, err := jwt.Parse(auth, keyFunc)
			claims, _ := token.Claims.(jwt.MapClaims)
			fmt.Println(claims["Customer"])
			if claims["Store"] != nil {
				return nil, errors.New("Role not store")
			}
			fmt.Println(claims["Store"])
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}
	return middleware.JWTWithConfig(config)
}
