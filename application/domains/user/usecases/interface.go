package usecases

import (
	"context"
	"ia-03-backend/application/domains/user/models"
)

type IUseCase interface {
	Login(ctx context.Context, req *models.LoginReq) (*models.LoginRes, error)
	Register(ctx context.Context, req *models.RegisterReq) (*models.RegisterRes, error)
}
