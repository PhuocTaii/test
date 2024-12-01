package routes

import (
	dbsv "ia-03-backend/application/domains/db/init"
	userhandler "ia-03-backend/application/domains/user/handler"
	commonmodels "ia-03-backend/application/models"
	"ia-03-backend/pkg/http/route"
)

type IHandler interface {
	NewGroupRoutes() []route.GroupRoute
}

type handler struct {
	lib         *commonmodels.Lib
	dbSv        *dbsv.DbService
	userHandler userhandler.Handler
}

func InitHandler(
	lib *commonmodels.Lib,
	dbSv *dbsv.DbService,
	userHandler userhandler.Handler,
) IHandler {
	return &handler{
		lib:         lib,
		dbSv:        dbSv,
		userHandler: userHandler,
	}
}

func (h *handler) NewGroupRoutes() []route.GroupRoute {
	groupRoute := make([]route.GroupRoute, 0)
	groupRoute = append(
		groupRoute,
		h.userGroupRoute(),
	)

	return groupRoute
}
