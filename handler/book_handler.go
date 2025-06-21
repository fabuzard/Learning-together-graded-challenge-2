package handler

import (
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
	genreParam := c.QueryParam("genre")
	if genreParam != "" {
		genreID, err := strconv.Atoi(genreParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid genre id"})
		}
		books, err := h.BookService.GetBooksByGenre(genreID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to fetch books"})
		}
		return c.JSON(http.StatusOK, books)
	}
	books, err := h.BookService.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to fetch books"})
	}
	return c.JSON(http.StatusOK, books)
}
