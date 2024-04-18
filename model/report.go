package model

type Report struct {
	User                   string  `gorm:"column:user" json:"user"`
	JumlahHariKerja        int     `gorm:"column:jumlah_hari_kerja" json:"jumlah_hari_kerja"`
	JumlahTransaksiBarang  int     `gorm:"column:jumlah_transaksi_barang" json:"jumlah_transaksi_barang"`
	JumlahTransaksiJasa    int     `gorm:"column:jumlah_transaksi_jasa" json:"jumlah_transaksi_jasa"`
	NominalTransaksiBarang float64 `gorm:"column:nominal_transaksi_barang" json:"nominal_transaksi_barang"`
	NominalTransaksiJasa   float64 `gorm:"column:nominal_transaksi_jasa" json:"nominal_transaksi_jasa"`
}
