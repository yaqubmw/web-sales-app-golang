package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	Id               uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	UserId           uuid.UUID `json:"user_id"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
	Jenis            string    `json:"jenis_pembelian"`
	Nominal          int       `json:"nominal"`
	User             User      `gorm:"foreignKey:UserId"`
}

func (Sale) TableName() string {
	return "sales"
}
