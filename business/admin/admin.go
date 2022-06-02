package admin

import "github.com/golang-jwt/jwt/v4"

type AdminSwagger struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type InputAdminToken struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Admin struct {
	ID       int    `gorm:"primaryKey"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type Kota struct {
	ID                 int    `json:"id"`
	Rajaongkir_city_id int    `json:"rajaongkir_city_id" validate:"required"`
	Kota_Nama          string `json:"kota_nama" validate:"required"`
	Postal_code        int    `json:"postal_code" validate:"required"`
	Tipe               string `json:"tipe" validate:"required"`
	Province_ID        int    `json:"province_id" validate:"required"`
}

type TestAdmin struct {
	ID       int    `gorm:"primaryKey"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Auth struct {
	Token string
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

type GetCityById struct {
	Nama_kota string `json:"nama_kota"`
	Tipe_kota string `json:"tipe_kota"`
}
