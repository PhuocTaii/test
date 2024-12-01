package service

import (
	"context"
	"ia-03-backend/application/domains/db/user/entity"
	"ia-03-backend/application/domains/db/user/models"
	"ia-03-backend/application/domains/db/user/repository"
	commonmodels "ia-03-backend/application/models"
)

type service struct {
	lib  *commonmodels.Lib
	repo repository.IRepository
}

func InitService(lib *commonmodels.Lib, repo repository.IRepository) IService {
	return &service{
		lib:  lib,
		repo: repo,
	}
}

// Create implements IService.
func (s *service) Create(ctx context.Context, params *models.SaveRequest) (*models.Response, error) {
	obj := &entity.User{}
	obj.ParseFromSaveRequest(params)
	res, err := s.repo.Create(ctx, obj)
	if err != nil {
		return nil, err
	}
	return res.Export(), nil
}

// GetByEmail implements IService.
func (s *service) GetByEmail(ctx context.Context, email string) (*models.Response, error) {
	res, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return res.Export(), nil
}
