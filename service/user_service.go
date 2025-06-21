package service

import (
	"errors"
	"gc2/model"
	"gc2/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(first_name, last_name, address, email, password, date_of_birth string) (*model.User, error)
	Login(email, password string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) Register(first_name, last_name, address, email, password, date_of_birth string) (*model.User, error) {
	_, err := s.userRepo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("email sudah terdaftar")
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &model.User{
		FirstName:   first_name,
		LastName:    last_name,
		Address:     address,
		Email:       email,
		Password:    string(hash),
		DateOfBirth: date_of_birth,
	}
	err = s.userRepo.Create(user)
	return user, err
}

func (s *userService) FindByID(id uint) (*model.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) Login(email, password string) (*model.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("password salah")
	}
	return user, nil
}
