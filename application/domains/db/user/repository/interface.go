package repository

import (
	"context"
	"ia-03-backend/application/domains/db/user/entity"
)

type IRepository interface {
	Create(ctx context.Context, obj *entity.User) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}
