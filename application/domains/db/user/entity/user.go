package entity

import (
	"ia-03-backend/application/domains/db/user/models"
	"time"
)

type User struct {
	UserId    int64     `gorm:"primaryKey;column:user_id"`
	FullName  string    `gorm:"column:full_name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Salt      string    `gorm:"column:salt"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}

func (User) TableName() string {
	return "tbl_users"
}

func (e *User) ParseFromSaveRequest(req *models.SaveRequest) {
	if req != nil {
		e.FullName = req.FullName
		e.Email = req.Email
		e.Password = req.Password
		e.Salt = req.Salt
		e.CreatedAt = time.Now().UTC()
	}
}

func (e *User) Export() *models.Response {
	if e != nil {
		return &models.Response{
			UserId:    e.UserId,
			FullName:  e.FullName,
			Email:     e.Email,
			Password:  e.Password,
			Salt:      e.Salt,
			CreatedAt: e.CreatedAt,
		}
	}
	return nil
}
