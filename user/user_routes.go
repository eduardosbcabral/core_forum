package user

import (
	"github.com/eduardosbcabral/core_forum/config"
)

var Controller = &UserController{UserRepository: UserRepository{}}
var routes = config.Routes{{}}

func CreateUserRoutes() config.Routes {
	routes := config.Routes{
		config.Route{
			"Index - User",
			"GET",
			"/user",
			Controller.Index,
		},
		config.Route{
			"Create - User",
			"POST",
			"/user",
			Controller.Create,
		},
		config.Route{
			"Login - User",
			"POST",
			"/login",
			Controller.Login,
		},
		/*config.Route{
			"Show - User",
			"GET",
			"/user/{username}",
			controller.Show,
		},
		config.Route{
			"Update - User",
			"PUT",
			"/user/{username}",
			Controller.Update,
		},
		config.Route{
			"Destroy - User",
			"DELETE",
			"/user/{username}",
			Controller.Destroy,
		},*/
		config.Route{
			"Index(All) - User",
			"GET",
			"/userAll",
			Controller.IndexAll,
		},
	}

	return routes
}