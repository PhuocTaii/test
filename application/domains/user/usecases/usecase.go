package usecases

import (
	"context"
	"errors"
	initdb "ia-03-backend/application/domains/db/init"
	usermodels "ia-03-backend/application/domains/db/user/models"
	userSv "ia-03-backend/application/domains/db/user/service"
	"ia-03-backend/application/domains/user/models"
	commonmodels "ia-03-backend/application/models"
	"ia-03-backend/pkg/utils"
)

type useCase struct {
	lib    *commonmodels.Lib
	userSv userSv.IService
}

func InitUseCase(lib *commonmodels.Lib, initDb *initdb.DbService) IUseCase {
	return &useCase{
		lib:    lib,
		userSv: initDb.User.Service,
	}
}

// Login implements IUseCase.
func (u *useCase) Login(ctx context.Context, req *models.LoginReq) (*models.LoginRes, error) {
	// Validate request
	if err := req.ValidateRequest(); err != nil {
		return nil, err
	}

	// Check if email already exists
	user, err := u.userSv.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("email or password is incorrect")
	}

	if !utils.VerifyPassword(user.Password, user.Salt+req.Password) {
		return nil, errors.New("email or password is incorrect")
	}

	return &models.LoginRes{IsSuccess: true}, nil
}

// Register implements IUseCase.
func (u *useCase) Register(ctx context.Context, req *models.RegisterReq) (*models.RegisterRes, error) {
	// Validate request
	if err := req.ValidateRequest(); err != nil {
		return nil, err
	}

	// Check if email already exists
	user, err := u.userSv.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("email already exists")
	}

	// Random salt and hash password
	salt := utils.GenerateRandomString(20)
	password := salt + req.Password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create user
	userReq := &usermodels.SaveRequest{
		FullName: req.FullName,
		Email:    req.Email,
		Password: *hashedPassword,
		Salt:     salt,
	}
	res, err := u.userSv.Create(ctx, userReq)
	if err != nil {
		return nil, err
	}

	return &models.RegisterRes{Id: res.UserId}, nil
}
