package model

type Author struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Books     []Book
}
