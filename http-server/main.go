package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Test")

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/welcome", welcHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
func welcHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Welcome to my website")
}
