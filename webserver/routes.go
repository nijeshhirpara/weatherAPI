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

func apiRoutes(ctx context.Context) {
	apiRouter.HandleFunc("/weather", api.HandleWeather).Methods("GET")
}

func webRoutes(ctx context.Context) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
}
