package webserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"golang.org/x/net/context"
)

func Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	apiRoutes(ctx)
	webRoutes(ctx)

	http.Handle("/", router)

	fmt.Println("Webserver is started on :8081")
	httpErr := http.ListenAndServe(":8081", handlers.LoggingHandler(os.Stdout, router))
	if httpErr != nil {
		fmt.Println("Could not initialise HTTP listener:", httpErr)
		return httpErr
	}

	return nil
}
