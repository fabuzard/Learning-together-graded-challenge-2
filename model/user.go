package model

type User struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	FirstName   string `gorm:"not null" json:"first_name"`
	LastName    string `gorm:"not null" json:"last_name"`
	Email       string `gorm:"unique;not null" json:"email"`
	Password    string `gorm:"not null" json:"password"`
	Address     string `gorm:"not null" json:"address"`
	DateOfBirth string `gorm:"not null" json:"date_of_birth"`
	Role        string `gorm:"default:user" json:"role"`
	Loans       []Loan `gorm:"foreignKey:UserID"`
}

type RegisterRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
}
