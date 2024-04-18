package repository

import (
	"github.com/yaqubmw/web-sales-app-golang/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(payload model.User) error
	GetByEmailPassword(email string, password string) (model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func (u *userRepo) Create(payload model.User) error {
	return u.db.Create(&payload).Error
}

func (u *userRepo) GetByEmailPassword(email string, password string) (model.User, error) {
	user := model.User{}
	err := u.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
