package repository

import (
	"gc2/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAuthorsWithBookCount() ([]model.AuthorAggResponse, error)
	GetGenresWithLoanCount() ([]model.GenreAggResponse, error)
	GetTopUsersByLoanCount() ([]model.TopUserResponse, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) GetAuthorsWithBookCount() ([]model.AuthorAggResponse, error) {
	var authors []model.AuthorAggResponse
	err := r.db.Table("authors").
		Select("authors.id, authors.first_name, authors.last_name, COUNT(books.id) AS book_count").
		Joins("LEFT JOIN books ON books.author_id = authors.id").
		Group("authors.id").
		Scan(&authors).Error
	return authors, err
}

func (r *adminRepository) GetGenresWithLoanCount() ([]model.GenreAggResponse, error) {
	var genres []model.GenreAggResponse
	err := r.db.Table("genres").
		Select("genres.id, genres.name, COUNT(loans.id) AS loan_count").
		Joins("LEFT JOIN book_genres ON book_genres.genre_id = genres.id").
		Joins("LEFT JOIN loans ON loans.book_id = book_genres.book_id").
		Group("genres.id").
		Scan(&genres).Error
	return genres, err
}

func (r *adminRepository) GetTopUsersByLoanCount() ([]model.TopUserResponse, error) {
	var users []model.TopUserResponse
	err := r.db.Table("users").
		Select("users.id, users.first_name, users.last_name, COUNT(loans.id) AS loan_count").
		Joins("LEFT JOIN loans ON loans.user_id = users.id").
		Group("users.id").
		Order("loan_count DESC").
		Limit(5).
		Scan(&users).Error
	return users, err
}
