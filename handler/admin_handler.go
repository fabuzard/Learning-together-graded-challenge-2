package handler

import (
	"gc2/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService service.AdminService
}

func NewAdminHandler(as service.AdminService) *AdminHandler {
	return &AdminHandler{AdminService: as}
}

// GET /admin/authors - Get authors with book count
func (h *AdminHandler) GetAuthors(c echo.Context) error {
	authors, err := h.AdminService.GetAuthorsWithBookCount()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch authors"})
	}
	return c.JSON(http.StatusOK, authors)
}

// GET /admin/genres - Get genres with loan count
func (h *AdminHandler) GetGenres(c echo.Context) error {
	genres, err := h.AdminService.GetGenresWithLoanCount()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch genres"})
	}
	return c.JSON(http.StatusOK, genres)
}

// GET /admin/users - Get top 5 users by loan count
func (h *AdminHandler) GetTopUsers(c echo.Context) error {
	users, err := h.AdminService.GetTopUsersByLoanCount()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch users"})
	}
	return c.JSON(http.StatusOK, users)
}
