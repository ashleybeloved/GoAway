package models

import "time"

type User struct {
	UserID    int
	Login     string
	Password  string
	CreatedAt time.Time
}

type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
