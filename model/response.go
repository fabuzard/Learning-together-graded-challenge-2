package model

import "time"

// Users
type UserMeLoanResponse struct {
	BookTitle string    `json:"book_title"`
	StartDate time.Time `json:"start_date"`
	DueDate   time.Time `json:"due_date"`
}

type UserMeResponse struct {
	ID          uint                 `json:"id"`
	FirstName   string               `json:"first_name"`
	LastName    string               `json:"last_name"`
	Email       string               `json:"email"`
	Address     string               `json:"address"`
	DateOfBirth string               `json:"date_of_birth"`
	Loans       []UserMeLoanResponse `json:"loan history"`
}

// Loan

type LoanResponse struct {
	StartDate time.Time `json:"start_date"`
	DueDate   time.Time `json:"due_date"`
	BookName  string    `json:"book_name"`
}

// Book
type BookResponse struct {
	ID                uint         `json:"id"`
	Title             string       `json:"title"`
	Description       string       `json:"description"`
	MinAgeRestriction int          `json:"min_age_restriction"`
	CoverUrl          string       `json:"cover_url"`
	Author            AuthorBasic  `json:"author"`
	Genres            []GenreBasic `json:"genres"`
}

type AuthorBasic struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GenreBasic struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
