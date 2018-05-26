package category

import (
	"https://github.com/eduardosbcabral/core_forum/config"
)

var controller = &CategoryController{}
var routes = config.Routes{{}}

func CreateCategoryRoutes() config.Routes {
	routes := config.Routes{
		config.Route{
			"Index - Category",
			"GET",
			"/category",
			controller.Index,
		},
		config.Route{
			"Create - Category",
			"POST",
			"/category",
			controller.Create,
		},
		config.Route{
			"Show - Category",
			"GET",
			"/category/{id}",
			controller.Show,
		},
		config.Route{
			"Update - Category",
			"PUT",
			"/category/{id}",
			controller.Update,
		},
		config.Route{
			"Destroy - Category",
			"DELETE",
			"/category/{id}",
			controller.Destroy,
		},
		config.Route{
			"Index(All) - Category",
			"GET",
			"/categoryAll",
			controller.IndexAll,
		},
	}

	return routes
}