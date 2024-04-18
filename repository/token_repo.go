package repository

import (
	"github.com/google/uuid"
	"github.com/yaqubmw/web-sales-app-golang/model"
	"gorm.io/gorm"
)

type TokenRepo interface {
	Create(payload model.Token) error
	Delete(id uuid.UUID) error
}

type tokenRepo struct {
	db *gorm.DB
}

func (t *tokenRepo) Create(payload model.Token) error {
	return t.db.Create(&payload).Error
}

func (t *tokenRepo) Delete(id uuid.UUID) error {
	token := model.Token{}
	return t.db.Where("id = ?", id).Delete(&token).Error
}

func NewTokenRepo(db *gorm.DB) TokenRepo {
	return &tokenRepo{
		db: db,
	}
}
