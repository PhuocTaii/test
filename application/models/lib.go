package models

import (
	"ia-03-backend/config/models"

	"gorm.io/gorm"
)

type Lib struct {
	Db  *gorm.DB
	Cfg *models.Config
}
