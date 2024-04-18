package manager

import "github.com/yaqubmw/web-sales-app-golang/repository"

type RepoManager interface {
	UserRepo() repository.UserRepo
	SaleRepo() repository.SaleRepo
	TokenRepo() repository.TokenRepo
	ReportRepo() repository.ReportRepo
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) ReportRepo() repository.ReportRepo {
	return repository.NewReportRepo(r.infra.Conn())
}

func (r *repoManager) SaleRepo() repository.SaleRepo {
	return repository.NewSaleRepo(r.infra.Conn())
}

func (r *repoManager) TokenRepo() repository.TokenRepo {
	return repository.NewTokenRepo(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepo {
	return repository.NewUserRepo(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
