package handler

import (
	"gc2/model"
	"gc2/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	BookService service.BookService
}

func NewBookHandler(bs service.BookService) *BookHandler {
	return &BookHandler{bs}
}
func (h *BookHandler) GetBooks(c echo.Context) error {
	genreQuery := c.QueryParam("genre")

	var books []model.Book
	var err error

	if genreQuery != "" {
		genreID, errConv := strconv.Atoi(genreQuery)
		if errConv != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid genre id"})
		}
		books, err = h.BookService.GetBooksByGenre(genreID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch books by genre"})
		}
	} else {
		books, err = h.BookService.GetAllBooks()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to fetch books"})
		}
	}

	var response []model.BookResponse
	for _, book := range books {
		var genres []model.GenreBasic
		for _, g := range book.Genres {
			genres = append(genres, model.GenreBasic{
				ID:   g.ID,
				Name: g.Name,
			})
		}

		response = append(response, model.BookResponse{
			ID:                book.ID,
			Title:             book.Title,
			Description:       book.Description,
			MinAgeRestriction: book.MinAgeRestriction,
			CoverUrl:          book.CoverUrl,
			Author: model.AuthorBasic{
				FirstName: book.Author.FirstName,
				LastName:  book.Author.LastName,
			},
			Genres: genres,
		})
	}

	return c.JSON(http.StatusOK, response)
}
