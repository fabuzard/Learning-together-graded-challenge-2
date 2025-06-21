package handler

import (
	"gc2/model"
	"gc2/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoanHandler struct {
	LoanService service.LoanService
}

func NewLoanHandler(ls service.LoanService) *LoanHandler {
	return &LoanHandler{ls}
}

func (h *LoanHandler) CreateLoan(c echo.Context) error {
	var req struct {
		BookID   uint `json:"book_id"`
		Duration int  `json:"duration"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid input"})
	}

	user := c.Get("userData").(model.User)

	loan, err := h.LoanService.Create(user.ID, req.BookID, req.Duration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to create loan"})
	}
	return c.JSON(http.StatusCreated, loan)
}
