package main 

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string , error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("Can not make a get call")
	} else {
		defer resp.Body.Close()
		return resp.Header.Get("content-type"), nil
	}
}

func main() {

	contentType, error :=contentType("https://fardinhakimi.com")

	if error != nil {
		fmt.Printf("Error: %v\n", error)
	} else {
		fmt.Printf("Content-Type: %v \n", contentType)
	}
}