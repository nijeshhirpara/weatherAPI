package main

import (
	"fmt"
	"log"
	"weatherapi/webserver"
)

func main() {
	fmt.Println("Trying to start server")
	webErr := webserver.Start()
	if webErr != nil {
		fmt.Println("Server cannot be able to start", webErr)
		log.Fatal(webErr)
	}
}
