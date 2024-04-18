package usecase

import (
	"github.com/yaqubmw/web-sales-app-golang/model"
	"github.com/yaqubmw/web-sales-app-golang/repository"
)

type SaleUsecase interface {
	Create(payload model.Sale) error
}

type saleUsecase struct {
	saleRepo repository.SaleRepo
}

func (s *saleUsecase) Create(payload model.Sale) error {
	return s.saleRepo.Create(payload)
}

func NewSaleUsecase(saleRepo repository.SaleRepo) SaleUsecase {
	return &saleUsecase{
		saleRepo: saleRepo,
	}
}
