package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	URL           string `gorm:"not null" json:"url"`
	ShortURL      string `gorm:"uniqueIndex;not null" json:"short_url"`
	CreatorUserID uint   `gorm:"index" json:"creator_id"`
	Clicks        int    `gorm:"default:0" json:"clicks"`
}

type LinkRequest struct {
	URL string `json:"url"`
}

type LinkUserRequest struct {
	ShortURL string `json:"short_url"`
}
