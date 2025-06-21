package model

import "time"

type Loan struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	BookID    uint      `gorm:"not null" json:"book_id"`
	StartDate time.Time `json:"start_date"`
	DueDate   time.Time `json:"due_date"`
	Book      Book
	User      User `gorm:"foreignKey:UserID"`
}
