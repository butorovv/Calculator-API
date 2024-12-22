package main

import (
	"Calculate-Service/internal/service"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", service.CalculateHandler)
	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
