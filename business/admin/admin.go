package admin

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type AdminSwagger struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResponseLogin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	No_hp    string `json:"no_hp"`
	Token    string `json:"token"`
}

type InputAdminToken struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterAdmin struct {
	gorm.Model
	Email    string `json:"email" gorm:"size:255"`
	Fullname string `json:"fullname" gorm:"size:255"`
	Password string `json:"password" gorm:"size:255"`
	No_hp    string `json:"no_hp" gorm:"size:80"`
}

type Admin struct {
	ID       int    `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required"`
	Password string `json:"password" validate:"required"`
	No_hp    string `json:"no_hp" validate:"required"`
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
	ID    int
	Email string
	jwt.StandardClaims
}
