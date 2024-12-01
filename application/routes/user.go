package routes

import (
	"ia-03-backend/pkg/http/method"
	"ia-03-backend/pkg/http/route"
)

func (h *handler) userGroupRoute() route.GroupRoute {
	return route.GroupRoute{
		Prefix: "/api/users",
		Routes: []route.Route{
			{
				Path:    "/login",
				Method:  method.POST,
				Handler: h.userHandler.Login,
			},
			{
				Path:    "/register",
				Method:  method.POST,
				Handler: h.userHandler.Register,
			},
		},
	}
}
