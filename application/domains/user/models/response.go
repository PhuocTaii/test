package models

type RegisterRes struct {
	Id int64 `json:"id"`
}

type LoginRes struct {
	IsSuccess bool `json:"is_success"`
}
