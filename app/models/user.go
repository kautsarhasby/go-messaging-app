package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `json:"username" gorm:"type:varchar(20);unique;" validate:"required,min=6,max=32"`
	Fullname  string `json:"fullname" gorm:"type:varchar(255);" validate:"required,min=password"`
	Password  string `json:"password" gorm:"type:varchar(255);" validate:"required,min=6"`
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
