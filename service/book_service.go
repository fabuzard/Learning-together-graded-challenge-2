package service

import (
	"gc2/model"
	"gc2/repository"
)

type BookService interface {
	GetAllBooks() ([]model.Book, error)
	GetBooksByGenre(genreID int) ([]model.Book, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBooksByGenre(genreID int) ([]model.Book, error) {
	return s.repo.FindByGenreID(genreID)
}
