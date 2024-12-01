package handler

import (
	"ia-03-backend/application/domains/user/models"
	"ia-03-backend/application/domains/user/usecases"
	commonmodels "ia-03-backend/application/models"
	"ia-03-backend/pkg/http/resp"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	lib     *commonmodels.Lib
	useCase usecases.IUseCase
}

func NewHandler(lib *commonmodels.Lib, useCase usecases.IUseCase) Handler {
	return Handler{
		lib:     lib,
		useCase: useCase,
	}
}

// Login godoc
// @Summary      Login a user
// @Description  Authenticate a user and return a token
// @Tags         UserService
// @Accept       json
// @Produce      json
// @Param        req body models.LoginReq true "Login request"
// @Success      200 {object} resp.SuccessResp "Login success"
// @Failure      400 {object} resp.ErrorResp "Invalid request"
// @Failure      500 {object} resp.ErrorResp "Internal server error"
// @Router       /users/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	req := new(models.LoginReq)
	if err := ctx.ShouldBind(&req); err != nil {
		panic(resp.ErrInvalidRequest(err))
	}

	result, err := h.useCase.Login(ctx, req)
	if err != nil {
		panic(resp.ErrInternalServer(err))
	}

	resp.Response(ctx, resp.NewSuccessResp(200, "Success", result))
}

// Register godoc
// @Summary      Register a new user
// @Description  Create a new user account
// @Tags         UserService
// @Accept       json
// @Produce      json
// @Param        req body models.RegisterReq true "Register request"
// @Success      200 {object} resp.SuccessResp "Register success"
// @Failure      400 {object} resp.ErrorResp "Invalid request"
// @Failure      500 {object} resp.ErrorResp "Internal server error"
// @Router       /users/register [post]
func (h *Handler) Register(ctx *gin.Context) {
	req := new(models.RegisterReq)
	if err := ctx.ShouldBind(&req); err != nil {
		panic(resp.ErrInvalidRequest(err))
	}

	result, err := h.useCase.Register(ctx, req)
	if err != nil {
		panic(resp.ErrInternalServer(err))
	}

	resp.Response(ctx, resp.NewSuccessResp(200, "Success", result))
}
