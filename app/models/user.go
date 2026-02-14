package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Users
type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `json:"username" gorm:"type:varchar(20);unique;" validate:"required,min=6,max=32"`
	Fullname  string `json:"fullname" gorm:"type:varchar(255);" validate:"required"`
	Password  string `json:"-" gorm:"type:varchar(255);" validate:"required,min=6"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserSession struct {
	ID                  uint      `gorm:"primarykey"`
	UserID              int       `json:"user_id" gorm:"type:int;" validate:"required"`
	Token               string    `json:"token" gorm:"type:varchar(255);" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:varchar(255);" validate:"required"`
	TokenExpired        time.Time `json:"token_expired" validate:"required"`
	RefreshTokenExpired time.Time `json:"refresh_token_expired" validate:"required"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

// Register
type RegisterRequest struct {
	Username string `json:"username" gorm:"type:varchar(20);unique;" validate:"required,min=6,max=32"`
	Fullname string `json:"fullname" gorm:"type:varchar(255);" validate:"required"`
	Password string `json:"password" gorm:"type:varchar(255);" validate:"required,min=6"`
}

type RegisterResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

func (l RegisterRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

// Login

type LoginRequest struct {
	Username string `json:"username" gorm:"type:varchar(20);" validate:"required,min=6,max=32"`
	Password string `json:"password" gorm:"type:varchar(255);" validate:"required,min=6"`
}

type LoginResponse struct {
	Username     string `json:"username"`
	Fullname     string `json:"fullname"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (l LoginRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
