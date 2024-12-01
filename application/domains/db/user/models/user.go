package models

import (
	"time"
)

type SaveRequest struct {
	FullName string
	Email    string
	Password string
	Salt     string
}

type Response struct {
	UserId    int64
	FullName  string
	Email     string
	Password  string
	Salt      string
	CreatedAt time.Time
}
