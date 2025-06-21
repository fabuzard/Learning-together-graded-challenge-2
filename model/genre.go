package model

type Genre struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Books []Book `gorm:"many2many:book_genres;" json:"-"`
}
