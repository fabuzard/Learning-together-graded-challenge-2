package repository

import (
	"gc2/model"
	"time"

	"gorm.io/gorm"
)

type LoanRepository interface {
	CreateLoan(userID, bookID uint, duration int) (*model.Loan, error)
}

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{db}
}

func (r *loanRepository) CreateLoan(userID, bookID uint, duration int) (*model.Loan, error) {
	loan := model.Loan{
		UserID:    userID,
		BookID:    bookID,
		StartDate: time.Now(),
		DueDate:   time.Now().AddDate(0, 0, duration),
	}

	// Create loan
	if err := r.db.Create(&loan).Error; err != nil {
		return nil, err
	}

	// Preload Book for response
	if err := r.db.Preload("Book").First(&loan, loan.ID).Error; err != nil {
		return nil, err
	}

	return &loan, nil
}
