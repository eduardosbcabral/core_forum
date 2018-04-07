package user

import (
	"core_backend/config"
)

var controller = &UserController{UserRepository: UserRepository{}}
var routes = config.Routes{{}}

func CreateUserRoutes() config.Routes {
	routes := config.Routes{
		config.Route{
			"Index - User",
			"GET",
			"/user",
			controller.Index,
		},
		config.Route{
			"Create - User",
			"POST",
			"/user",
			controller.Create,
		},
	}

	return routes
}