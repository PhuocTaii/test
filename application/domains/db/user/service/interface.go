package service

import (
	"context"
	"ia-03-backend/application/domains/db/user/models"
)

type IService interface {
	Create(ctx context.Context, params *models.SaveRequest) (*models.Response, error)
	GetByEmail(ctx context.Context, email string) (*models.Response, error)
}
