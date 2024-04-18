package manager

import (
	"fmt"

	"github.com/yaqubmw/web-sales-app-golang/config"
	"github.com/yaqubmw/web-sales-app-golang/utils/checker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	Conn() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
	cfg *config.Config
}

func (i *infraManager) initDB() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", i.cfg.Host, i.cfg.Port, i.cfg.User, i.cfg.Password, i.cfg.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checker.CheckErr(err)
	i.db = db
	return nil
}

func (i *infraManager) Conn() *gorm.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{
		cfg: cfg,
	}
	err := conn.initDB()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
