package thread

import (
	"core_backend/config"
)

var controllerP = &PostController{PostRepository: PostRepository{}}
var routesP = config.Routes{{}}

func CreatePostRoutes() config.Routes {
	routesP := config.Routes{
		config.Route{
			"Index - Post",
			"GET",
			"/{categoryId}/thread/{threadId}/post",
			controllerP.Index,
		},
		config.Route{
			"Create - Post",
			"POST",
			"/{categoryId}/thread/{threadId}/post",
			controllerP.Create,
		},
		config.Route{
			"Show - Post",
			"GET",
			"/{categoryId}/thread/{threadId}/post/{postId}",
			controllerP.Show,
		},
		config.Route{
			"Update - Post",
			"PUT",
			"/{categoryId}/thread/{threadId}/post/{postId}",
			controllerP.Update,
		},
		config.Route{
			"Destroy - Post",
			"DELETE",
			"/{categoryId}/thread/{threadId}/post/{postId}",
			controllerP.Destroy,
		},
		config.Route{
			"Index(All) - Post",
			"GET",
			"/postAll",
			controllerP.IndexAll,
		},
	}

	return routesP
}