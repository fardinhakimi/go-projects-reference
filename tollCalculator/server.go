package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func helloHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Printf("The request: %v", request.URL.Query())
	responseWriter.WriteHeader(200)
	responseWriter.Header().Set("content-type", "application/json")
	json.NewEncoder(responseWriter).Encode(`{ "hello": "world! hakimi"}`)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Fatal("Failed to get port")
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server running on port %d \n", port)
}
