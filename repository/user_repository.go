package repository

import (
	"gc2/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	FindByID(id uint) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.Preload("Loans").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
