package auth

import (
	"net/http"

	"github.com/eduardosbcabral/core_forum/config"
    "github.com/codegangsta/negroni"
)

var routes = config.Routes{{}}

func CreateAuthRoutes() config.Routes {
	routes := config.Routes{

	}

	return routes
}


func SetAuthenticatedMiddleware(r func(http.ResponseWriter, *http.Request)) (n *negroni.Negroni) {
    n = negroni.New(negroni.HandlerFunc(ValidateToken), negroni.Wrap(http.HandlerFunc(r)))
    return
}