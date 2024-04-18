package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yaqubmw/web-sales-app-golang/config"
	"github.com/yaqubmw/web-sales-app-golang/model"
	"github.com/yaqubmw/web-sales-app-golang/utils/checker"
)

type Claims struct {
	UserId uuid.UUID
	Name   string
	Email  string
	jwt.RegisteredClaims
}

func GenerateToken(user model.User) (string, error) {
	cfg, err := config.NewConfig()
	checker.CheckErr(err)
	now := time.Now().UTC()
	end := now.Add(cfg.AccessTokenExpiry)

	claims := &Claims{
		UserId: user.Id,
		Name:   user.Nama,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
	}

	token := jwt.NewWithClaims(cfg.JwtSigningMethod, claims)
	tokenString, err := token.SignedString(cfg.JwtSignatureKey)
	if err != nil {
		return "", fmt.Errorf("failed to create token: %v", err)
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	cfg, err := config.NewConfig()
	checker.CheckErr(err)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return cfg.JwtSignatureKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
