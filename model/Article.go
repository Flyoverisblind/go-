package model

type Article struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string
}
