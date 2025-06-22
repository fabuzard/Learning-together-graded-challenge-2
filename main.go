package main

import (
	"fmt"
	"gc2/config"
	"gc2/handler"
	"gc2/middleware"
	"gc2/model"
	"gc2/repository"
	"gc2/service"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	config.LoadEnv()

	db := config.DBInit()

	e := echo.New()

	//User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Book
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	// Loan
	loanRepo := repository.NewLoanRepository(db)
	loanService := service.NewLoanService(loanRepo)
	loanHandler := handler.NewLoanHandler(loanService)

	// Testing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "🚀 Server running and DB connected!")
	})

	// Group user
	userGroup := e.Group("/users")
	userGroup.POST("/register", userHandler.Register)
	userGroup.POST("/login", userHandler.Login)

	userGroup.GET("/me", userHandler.Me, middleware.JWTMiddleware((os.Getenv("JWT_SECRET"))))

	// Group book
	bookGroup := e.Group("/books")
	bookGroup.GET("", bookHandler.GetBooks)

	// Group loan
	loanGroup := e.Group("/loans")
	loanGroup.POST("", loanHandler.CreateLoan, middleware.JWTMiddleware((os.Getenv("JWT_SECRET"))))

	db.AutoMigrate(
		&model.Author{},
		&model.Book{},
		&model.BookGenre{},
		&model.Genre{},
		&model.Loan{},
		&model.User{},
	)
	fmt.Println("✅ Connected to PostgreSQL database!")
	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}
