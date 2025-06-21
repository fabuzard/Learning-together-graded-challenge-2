package model

type Book struct {
	ID                uint   `gorm:"primaryKey" json:"id"`
	Title             string `gorm:"not null" json:"title"`
	Description       string `gorm:"not null" json:"description"`
	MinAgeRestriction int    `gorm:"not null" json:"min_age_restriction"`
	CoverUrl          string `gorm:"not null" json:"cover_url"`
	AuthorID          uint   `gorm:"not null"`
	Author            Author
	Genres            []Genre `gorm:"many2many:book_genres;" json:"genres"`
	Loans             []Loan
}
