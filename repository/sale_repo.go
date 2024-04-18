package repository

import (
	"github.com/yaqubmw/web-sales-app-golang/model"
	"gorm.io/gorm"
)

type SaleRepo interface {
	Create(payload model.Sale) error
}

type saleRepo struct {
	db *gorm.DB
}

func (s *saleRepo) Create(payload model.Sale) error {
	return s.db.Create(&payload).Error
}

func NewSaleRepo(db *gorm.DB) SaleRepo {
	return &saleRepo{
		db: db,
	}
}
