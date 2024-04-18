package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	Nama     string    `json:"nama"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (User) TableName() string {
	return "users"
}
