package handler

import (
	"gc2/service"
	"net/http"
	"os"

	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(us service.UserService) *UserHandler {
	return &UserHandler{us}
}

func (h *UserHandler) Register(c echo.Context) error {
	var req struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Address     string `json:"address"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		DateOfBirth string `json:"date_of_birth"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid input"})
	}

	// Validate missing fields
	if req.FirstName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "first_name is required"})
	}
	if req.LastName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "last_name is required"})
	}
	if req.Address == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "address is required"})
	}
	if req.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "email is required"})
	}
	if req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "password is required"})
	}
	if req.DateOfBirth == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "date_of_birth is required"})
	}

	user, err := h.UserService.Register(req.FirstName, req.LastName, req.Address, req.Email, req.Password, req.DateOfBirth)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email,
	})
}

func (h *UserHandler) Login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid input"})
	}
	user, err := h.UserService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Generate JWT token pakai data user

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "JWT_SECRET not set"})
	}

	log.Println("JWT_SECRET (user_handler.go):", jwtSecret) // Debug log
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}

func (h *UserHandler) Me(c echo.Context) error {
	userToken, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok || !userToken.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user_id in token"})
	}

	userID := uint(userIDFloat)
	user, err := h.UserService.FindByID(userID)
	log.Println("user_id from JWT:", userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	// Buat response custom biar password gak ikut ke-return
	response := map[string]interface{}{
		"id":            user.ID,
		"first_name":    user.FirstName,
		"last_name":     user.LastName,
		"email":         user.Email,
		"address":       user.Address,
		"date_of_birth": user.DateOfBirth,
		"loans":         user.Loans,
	}

	return c.JSON(http.StatusOK, response)
}
