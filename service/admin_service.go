package service

import (
	"gc2/model"
	"gc2/repository"
)

type AdminService interface {
	GetAuthorsWithBookCount() ([]model.AuthorAggResponse, error)
	GetGenresWithLoanCount() ([]model.GenreAggResponse, error)
	GetTopUsersByLoanCount() ([]model.TopUserResponse, error)
}

type adminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return &adminService{adminRepo: repo}
}

func (s *adminService) GetAuthorsWithBookCount() ([]model.AuthorAggResponse, error) {
	return s.adminRepo.GetAuthorsWithBookCount()
}

func (s *adminService) GetGenresWithLoanCount() ([]model.GenreAggResponse, error) {
	return s.adminRepo.GetGenresWithLoanCount()
}

func (s *adminService) GetTopUsersByLoanCount() ([]model.TopUserResponse, error) {
	return s.adminRepo.GetTopUsersByLoanCount()
}
