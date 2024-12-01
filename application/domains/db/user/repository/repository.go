package repository

import (
	"context"
	"errors"
	"ia-03-backend/application/domains/db/user/entity"
	commonmodels "ia-03-backend/application/models"

	"gorm.io/gorm"
)

type repository struct {
	lib       *commonmodels.Lib
	tableName string
}

func InitRepo(lib *commonmodels.Lib) IRepository {
	return &repository{
		lib:       lib,
		tableName: (&entity.User{}).TableName(),
	}
}

func (r *repository) dbWithContext(ctx context.Context) *gorm.DB {
	return r.lib.Db.WithContext(ctx)
}

// Create implements IRepository.
func (r *repository) Create(ctx context.Context, obj *entity.User) (*entity.User, error) {
	result := r.dbWithContext(ctx).Create(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

// GetByEmail implements IRepository.
func (r *repository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	obj := &entity.User{}
	result := r.dbWithContext(ctx).Where("email = ?", email).First(obj)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return obj, nil
}
