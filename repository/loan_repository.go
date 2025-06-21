package repository

import (
	"time"
	"gc2/model"
	"gorm.io/gorm"
)

type LoanRepository interface {
	CreateLoan(userID, bookID uint, duration int) (*model.Loan, error)
	FindByUserID(userID uint) ([]model.Loan, error)
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
	err := r.db.Create(&loan).Error
	return &loan, err
}

func (r *loanRepository) FindByUserID(userID uint) ([]model.Loan, error) {
	var loans []model.Loan
	err := r.db.Where("user_id = ?", userID).Preload("Book").Find(&loans).Error
	return loans, err
}