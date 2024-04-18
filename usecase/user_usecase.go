package usecase

import (
	"github.com/yaqubmw/web-sales-app-golang/model"
	"github.com/yaqubmw/web-sales-app-golang/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Create(payload model.User) error
	GetByEmailPassword(email string, password string) (model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func (u *userUsecase) Create(payload model.User) error {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	payload.Password = string(bytes)
	err := u.userRepo.Create(payload)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) GetByEmailPassword(email string, password string) (model.User, error) {
	return u.userRepo.GetByEmailPassword(email, password)
}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
