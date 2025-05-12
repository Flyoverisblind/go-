package model

import (
	"time"
)

type Article struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Content   string
	Date      time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Views     uint
}
