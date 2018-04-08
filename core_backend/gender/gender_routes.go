package gender

import (
	"core_backend/config"
)

var controller = &GenderController{GenderRepository: GenderRepository{}}
var routes = config.Routes{{}}

func CreateGenderRoutes() config.Routes {
	routes := config.Routes{
		config.Route{
			"Index - Gender",
			"GET",
			"/gender",
			controller.Index,
		},
		config.Route{
			"Create - Gender",
			"POST",
			"/gender",
			controller.Create,
		},
		config.Route{
			"Show - Gender",
			"GET",
			"/gender/{id}",
			controller.Show,
		},
		config.Route{
			"Update - Gender",
			"PUT",
			"/gender/{id}",
			controller.Update,
		},
		config.Route{
			"Destroy - Gender",
			"DELETE",
			"/gender/{id}",
			controller.Destroy,
		},
	}

	return routes
}