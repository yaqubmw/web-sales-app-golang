package usecase

import (
	"github.com/google/uuid"
	"github.com/yaqubmw/web-sales-app-golang/model"
	"github.com/yaqubmw/web-sales-app-golang/repository"
)

type TokenUsecase interface {
	Create(payload model.Token) error
	Delete(id uuid.UUID) error
}

type tokenUsecase struct {
	tokenRepo repository.TokenRepo
}

func (t *tokenUsecase) Create(payload model.Token) error {
	return t.tokenRepo.Create(payload)
}

func (t *tokenUsecase) Delete(id uuid.UUID) error {
	return t.tokenRepo.Delete(id)
}

func NewTokenUsecase(tokenRepo repository.TokenRepo) TokenUsecase {
	return &tokenUsecase{
		tokenRepo: tokenRepo,
	}
}
