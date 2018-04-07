package user

import ("core_back_end/config")

var controller = &UserController{UserRepository: UserRepository{}}

func CreateUserRoutes() config.Routes {
	var routes = config.Routes{
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
