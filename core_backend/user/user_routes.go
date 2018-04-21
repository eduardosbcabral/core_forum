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
		config.Route{
			"Show - User",
			"GET",
			"/user/{username}",
			controller.Show,
		},
		config.Route{
			"Login - User",
			"POST",
			"/login",
			controller.Login,
		},
		config.Route{
			"Update - User",
			"PUT",
			"/user/{username}",
			controller.Update,
		},
		config.Route{
			"Destroy - User",
			"DELETE",
			"/user/{username}",
			controller.Destroy,
		},
		config.Route{
			"Index(All) - User",
			"GET",
			"/userAll",
			controller.IndexAll,
		},
	}

	return routes
}