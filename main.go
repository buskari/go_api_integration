package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const BASE_URL string = "http://localhost:3000"

func main() {
	// fetch ALL TODOS using the GET method
	resp, err := http.Get(BASE_URL + "/todos")
	if err != nil {
		log.Fatalf("Failed to get todos: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println("Raw respnse:")
	fmt.Println(string(body))
}
