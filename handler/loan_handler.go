package handler

import (
	"gc2/helper"
	"gc2/model"
	"gc2/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoanHandler struct {
	loanervice service.LoanService
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

	userID, err := helper.ExtractUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	loan, err := h.loanervice.Create(userID, req.BookID, req.Duration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to create loan"})
	}

	response := model.LoanResponse{
		StartDate: loan.StartDate,
		DueDate:   loan.DueDate,
		BookName:  loan.Book.Title,
	}
	return c.JSON(http.StatusCreated, response)
}
