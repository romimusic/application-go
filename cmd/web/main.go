package main

import (
	"fmt"
	"net/http"

	"github.com/romimusic/application-go/pkg/handlers"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
