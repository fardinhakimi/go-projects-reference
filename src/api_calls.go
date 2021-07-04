package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func makeGetRequest() {
	var url string = "https://httpbin.org/get"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("Error: can't call %s", url)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

func makePostRequest() {

	type Job struct {
		User   string `json:"user"`
		Action string `json:"action"`
		Count  int    `json:"count"`
	}

	// POST REQUEST BODY

	job := &Job{
		User:   "Fardin",
		Action: "punch",
		Count:  1,
	}

	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)

	if err := enc.Encode(job); err != nil {
		log.Fatalf("Error:= can't encode job - %s", err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", &buf)

	if err != nil {
		log.Fatalf("Error: can't call httpbin.org")
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

func main() {
	makeGetRequest()
	makePostRequest()
}
