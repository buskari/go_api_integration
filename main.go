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

	fmt.Printf("Response Status Code: %d", int(resp.StatusCode))
	if resp.StatusCode == http.StatusOK {
		fmt.Println("\nRaw response:")
		fmt.Println(string(body))
	} else if resp.StatusCode == http.StatusNotFound {
		fmt.Println("\nThe required resource does not exists.")
		fmt.Printf("Error details: %v", body)
	} else if resp.StatusCode == http.StatusUnauthorized {
		fmt.Println("\nUnauthorized. Access denied.")
		fmt.Printf("Error details: %v", body)
	} else if resp.StatusCode == http.StatusBadRequest {
		fmt.Println("\nBad Request. The server does not understand the request due to syntax error.")
		fmt.Printf("Error details: %v", body)
	} else if resp.StatusCode == http.StatusInternalServerError {
		fmt.Println("\nInternal Server Error. The server could not handle the request. Try again later.")
		fmt.Printf("Error details: %v", body)
	} else {
		fmt.Println("\nUnexpected Status Code")
		fmt.Printf("Error details: %v", body)
	}
}
