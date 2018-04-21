package thread

import (
	"core_backend/config"
)

var controller = &ThreadController{ThreadRepository: ThreadRepository{}}
var routes = config.Routes{{}}

func CreateThreadRoutes() config.Routes {
	routes := config.Routes{
		config.Route{
			"Index - Thread",
			"GET",
			"/{categoryId}/thread",
			controller.Index,
		},
		config.Route{
			"Create - Thread",
			"POST",
			"/{categoryId}/thread",
			controller.Create,
		},
		config.Route{
			"Show - Thread",
			"GET",
			"/{categoryId}/thread/{threadId}",
			controller.Show,
		},
		config.Route{
			"Update - Thread",
			"PUT",
			"/{categoryId}/thread/{threadId}",
			controller.Update,
		},
		config.Route{
			"Destroy - Thread",
			"DELETE",
			"/{categoryId}/thread/{threadId}",
			controller.Destroy,
		},
		config.Route{
			"Index(All) - Thread",
			"GET",
			"/threadAll",
			controller.IndexAll,
		},
	}

	return routes
}