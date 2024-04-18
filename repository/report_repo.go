package repository

import (
	"time"

	"github.com/yaqubmw/web-sales-app-golang/model"
	"gorm.io/gorm"
)

type ReportRepo interface {
	GetReport(startDate, endDate time.Time) ([]model.Report, error)
}

type reportRepo struct {
	db *gorm.DB
}

func (r *reportRepo) GetReport(startDate time.Time, endDate time.Time) ([]model.Report, error) {
	report := []model.Report{}
	err := r.db.Table("users u").Select("u.nama AS user, COUNT(DISTINCT s.tanggal_transaksi) AS jumlah_hari_kerja, SUM(CASE WHEN s.jenis = 'barang' THEN 1 ELSE 0 END) AS jumlah_transaksi_barang, SUM(CASE WHEN s.jenis = 'jasa' THEN 1 ELSE 0 END) AS jumlah_transaksi_jasa, SUM(CASE WHEN s.jenis = 'barang' THEN s.nominal ELSE 0 END) AS nominal_transaksi_barang, SUM(CASE WHEN s.jenis = 'jasa' THEN s.nominal ELSE 0 END) AS nominal_transaksi_jasa").Joins("JOIN sales s ON u.id = s.user_id").Where("s.tanggal_transaksi BETWEEN ? AND ?", startDate, endDate).Group("u.nama").Find(&report).Error
	if err != nil {
		return nil, err
	}
	return report, nil
}

func NewReportRepo(db *gorm.DB) ReportRepo {
	return &reportRepo{
		db: db,
	}
}
