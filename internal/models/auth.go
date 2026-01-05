package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
