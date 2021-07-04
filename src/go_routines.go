package main

import (
	"fmt"
	"net/http"
	"github.com/pkg/errors"
	"sync"
)

func returnType(url string) (string, error){

	res, err := http.Get(url)

	if err != nil {
		return "", errors.Wrap(err, "failed to get content-type")
	}

	defer res.Body.Close()
	contentType := res.Header.Get("content-type")

	return contentType, nil
}

func main() {

	urls := []string{
		"https://fardinhakimi.com",
		"https://api.github.com",
		"https://google.com",
	}


	var wg sync.WaitGroup

	for _, url := range urls {

		// The go keyword makes it a go-routine

		wg.Add(1)

		go func (url string) {
			contentType, err := returnType(url)
			wg.Done()
			if err != nil {
				fmt.Printf("error %s \n", err)
			} else {
				fmt.Printf("%s -> %s \n", url, contentType)
			}
		}(url)	
	}

	wg.Wait()
}