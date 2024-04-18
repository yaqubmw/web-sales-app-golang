package manager

import "github.com/yaqubmw/web-sales-app-golang/usecase"

type UsecaseManager interface {
	UserUsecase() usecase.UserUsecase
	SaleUsecase() usecase.SaleUsecase
	TokenUsecase() usecase.TokenUsecase
	ReportUsecase() usecase.ReportUsecase
	AuthUsecase() usecase.AuthUsecase
}

type usecaseManager struct {
	repoManager RepoManager
}

func (u *usecaseManager) AuthUsecase() usecase.AuthUsecase {
	return usecase.NewAuthUsecase(u.repoManager.TokenRepo(), u.UserUsecase())
}

func (u *usecaseManager) ReportUsecase() usecase.ReportUsecase {
	return usecase.NewReportUsecase(u.repoManager.ReportRepo())
}

func (u *usecaseManager) SaleUsecase() usecase.SaleUsecase {
	return usecase.NewSaleUsecase(u.repoManager.SaleRepo())
}

func (u *usecaseManager) TokenUsecase() usecase.TokenUsecase {
	return usecase.NewTokenUsecase(u.repoManager.TokenRepo())
}

func (u *usecaseManager) UserUsecase() usecase.UserUsecase {
	return usecase.NewUserUsecase(u.repoManager.UserRepo())
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
