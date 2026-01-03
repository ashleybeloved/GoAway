package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	URL      string `gorm:"not null"`
	ShortURL string `gorm:"uniqueIndex;not null"`
}
