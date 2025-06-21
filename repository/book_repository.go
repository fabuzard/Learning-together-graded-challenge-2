package repository

import (
	"gc2/config"
	"gc2/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]model.Book, error)
	FindByGenreID(genreID int) ([]model.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := config.DB.Preload("Genres").Preload("Author").Find(&books).Error
	return books, err
}

func (br *bookRepository) FindByGenreID(genreID int) ([]model.Book, error) {
	var books []model.Book
	err := config.DB.Preload("Genres").Preload("Author").
		Joins("JOIN book_genres ON book_genres.book_id = books.id").
		Where("book_genres.genre_id = ?", genreID).Find(&books).Error
	return books, err
}
