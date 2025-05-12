package model

import (
	"time"
)

type Article struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string
	Date    time.Time `gorm:"autoCreateTime"`
	Views   uint
}
