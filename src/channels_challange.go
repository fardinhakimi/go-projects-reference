package main 

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan) {

	res, err := http.Get(url)

	if err != nil {
		out <- fmt.Printf("Failed to get content-type for url %s", url)
	}

	defer res.Body.Close()

	contentType := res.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s", url, contentType)
}

func main() {

	urls := []string{
		"https://fardinhakimi.com",
		"https://api.github.com",
		"https://google.com",
	}

	ch := make(chan string)

	for _, url := range urls {
		go returnType(url, ch)
	}

	for range urls {
		out := <-ch
		fmt.Println(out)
	}

}