package webserver

import (
	"weatherapi/webserver/api"

	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

var (
	router    = mux.NewRouter()
	apiRouter = router.PathPrefix("/v1").Subrouter()
)

// apiRoutes is a collection of API routes
func apiRoutes(ctx context.Context) {
	apiRouter.Handle("/weather", cached("3s", true, api.HandleWeather)).Methods("GET")
}

// webRoutes is a collection of web routes
func webRoutes(ctx context.Context) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
}
