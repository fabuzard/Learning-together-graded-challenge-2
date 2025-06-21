package model

type BookGenre struct {
	BookID  uint `gorm:"primaryKey"`
	GenreID uint `gorm:"primaryKey"`
}
