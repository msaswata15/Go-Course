package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/bar")
	if err != nil {
		// handle error
		log.Fatalf("http get failed %s\n", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	// ...
	if err != nil {
		// handle error
		log.Fatalf("readall failure %s\n", err)
	}
	fmt.Printf("Body: %s\n", body)

}
