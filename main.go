package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

func AddValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(":8080", nil)
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}
