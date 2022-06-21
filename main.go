package main

import (
	"api-redeem-point/api"
	"api-redeem-point/app/modules"
	"api-redeem-point/config"
	"api-redeem-point/repository"
	"api-redeem-point/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "api-redeem-point/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

// @title API Jasa Pengiriman
// @version 1.0
// @description Berikut API Jasa Pengiriman
// @host api-poins-id.herokuapp.com/v1
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	port := os.Getenv("PORT")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := config.GetConfig()
	dbCon := utils.NewConnectionDatabase(config)

	defer dbCon.CloseConnection()

	controllers := modules.RegistrationModules(dbCon, config)
	var db gorm.DB
	db.AutoMigrate(&repository.History_Transaction{})
	db.AutoMigrate(&repository.Store{})
	db.AutoMigrate(&repository.Customer{})
	db.AutoMigrate(&repository.StockProduct{})
	db.AutoMigrate(&repository.Admin{})

	e := echo.New()
	handleSwagger := echoSwagger.WrapHandler
	e.GET("/swagger/*", handleSwagger)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})
	api.RegistrationPath(e, controllers)

	go func() {
		if port == "" {
			port = "8080"
		}
		address := fmt.Sprintf(":%s", port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	<-quit
}
