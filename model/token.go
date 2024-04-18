package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	User      User      `gorm:"foreignKey:UserId"`
}

func (Token) TableName() string {
	return "tokens"
}
