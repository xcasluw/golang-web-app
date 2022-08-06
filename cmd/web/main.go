package main

import (
	"fmt"
	"net/http"

	"github.com/xcasluw/golang-basic-app/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(":8080", nil)
}
