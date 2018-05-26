package thread

import (
	"github.com/eduardosbcabral/core_forum/config"
)

var controllerT = &ThreadController{ThreadRepository: ThreadRepository{}}
var routesT = config.Routes{{}}

func CreateThreadRoutes() config.Routes {
	routesT := config.Routes{
		config.Route{
			"Index - Thread",
			"GET",
			"/{categoryId}/thread",
			controllerT.Index,
		},
		config.Route{
			"Create - Thread",
			"POST",
			"/{categoryId}/thread",
			controllerT.Create,
		},
		config.Route{
			"Show - Thread",
			"GET",
			"/{categoryId}/thread/{threadId}",
			controllerT.Show,
		},
		config.Route{
			"Update - Thread",
			"PUT",
			"/{categoryId}/thread/{threadId}",
			controllerT.Update,
		},
		config.Route{
			"Destroy - Thread",
			"DELETE",
			"/{categoryId}/thread/{threadId}",
			controllerT.Destroy,
		},
		config.Route{
			"Index(All) - Thread",
			"GET",
			"/threadAll",
			controllerT.IndexAll,
		},
	}

	return routesT
}