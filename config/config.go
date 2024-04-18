package config

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yaqubmw/web-sales-app-golang/utils/checker"
	"github.com/yaqubmw/web-sales-app-golang/utils/common"
	"github.com/yaqubmw/web-sales-app-golang/utils/convert"
)

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	Driver   string
}

type TokenConfig struct {
	ApplicationName   string
	JwtSignatureKey   []byte
	JwtSigningMethod  *jwt.SigningMethodHMAC
	AccessTokenExpiry time.Duration
}

type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
}

func (c *Config) ReadConfig() error {
	err := common.LoadEnv()
	checker.CheckErr(err)

	c.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"),
		ApiPort: os.Getenv("API_PORT"),
	}

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	appTokenExpire, _ := convert.StrToInt(os.Getenv("APP_TOKEN_EXPIRE"))
	accessTokenExpiry := time.Duration(appTokenExpire) * time.Minute
	c.TokenConfig = TokenConfig{
		ApplicationName:   os.Getenv("APPLICATION_NAME"),
		JwtSignatureKey:   []byte(os.Getenv("JWT_SIGNATURE_KEY")),
		JwtSigningMethod:  jwt.SigningMethodHS256,
		AccessTokenExpiry: accessTokenExpiry,
	}
	
	if c.ApiConfig.ApiHost == "" || c.ApiConfig.ApiPort == "" || c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.DbName == "" || c.DbConfig.User == "" || c.DbConfig.Password == "" || c.DbConfig.Driver == "" || c.TokenConfig.ApplicationName == "" || c.TokenConfig.JwtSignatureKey == nil || c.TokenConfig.JwtSigningMethod == nil {
		return fmt.Errorf("missing required .env variables")
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
