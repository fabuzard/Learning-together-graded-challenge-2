package service

import (
	"gc2/model"
	"gc2/repository"
)

type LoanService interface {
	Create(userID, bookID uint, duration int) (*model.Loan, error)
}

type loanService struct {
	repo repository.LoanRepository
}

func NewLoanService(repo repository.LoanRepository) LoanService {
	return &loanService{repo}
}

func (s *loanService) Create(userID, bookID uint, duration int) (*model.Loan, error) {
	return s.repo.CreateLoan(userID, bookID, duration)
}
