package webserver

import (
	"log"
	"net/http"
	"os"
	"weatherapi/cache"
	"weatherapi/webserver/services"

	"github.com/gorilla/handlers"
	"golang.org/x/net/context"
)

// Start is responsible to initalize cache storage, register routes and start the server
func Start() error {
	services.Storage = cache.NewStorage()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	apiRoutes(ctx)
	webRoutes(ctx)

	http.Handle("/", router)

	port := os.Getenv("PORT")

	log.Println("Webserver is started on :" + port)
	httpErr := http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, router))
	if httpErr != nil {
		log.Fatal("Could not initialise HTTP listener:", httpErr)
		return httpErr
	}

	return nil
}
