package usecase

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yaqubmw/web-sales-app-golang/config"
	"github.com/yaqubmw/web-sales-app-golang/model"
	"github.com/yaqubmw/web-sales-app-golang/repository"
	"github.com/yaqubmw/web-sales-app-golang/utils/security"
)

type AuthUsecase interface {
	Login(email, password string) (string, error)
	Logout(userId uuid.UUID) error
}

type authUsecase struct {
	tokenRepo   repository.TokenRepo
	userUsecase UserUsecase
}

func (a *authUsecase) Login(email, password string) (string, error) {
	cfg, _ := config.NewConfig()
	user, err := a.userUsecase.GetByEmailPassword(email, password)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}
	generatedToken, err := security.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("failed to create token")
	}

	expiredAt := time.Now().Add(cfg.AccessTokenExpiry)
	token := model.Token{
		UserId:    user.Id,
		Token:     generatedToken,
		ExpiredAt: expiredAt,
	}

	err = a.tokenRepo.Create(token)
	if err != nil {
		return "", fmt.Errorf("failed to create token")
	}

	return generatedToken, nil
}

func (a *authUsecase) Logout(userId uuid.UUID) error {
	return a.tokenRepo.Delete(userId)
}

func NewAuthUsecase(tokenRepo repository.TokenRepo, userUsecase UserUsecase) AuthUsecase {
	return &authUsecase{
		tokenRepo:   tokenRepo,
		userUsecase: userUsecase,
	}
}
