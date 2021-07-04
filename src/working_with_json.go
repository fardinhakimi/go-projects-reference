package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// our json data
var data string = `
{
	"user": "Fardin Hakimi",
	"type": "deposit",
	"amount": 3500.6
}
`

// go construct / object represting the data
type Request struct {
	Login  string  `json:"user"` // this is called "a field" tag
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {

	reader := bytes.NewBufferString(data) // Simulates a file or a socker
	// Decode the request

	fmt.Printf(" The reader %v", reader)

	decoder := json.NewDecoder(reader)

	req := &Request{}

	// Populate the req object with decoded json
	if err := decoder.Decode(req); err != nil {
		log.Fatalf("Error: can't decode - %s", err)
	}

	fmt.Printf("Transaction object: %v\n", req)

	// create a response part

	prevBalance := 5000.0
	// interface{} -> any type in go

	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	// Encode the response to Json

	var buf bytes.Buffer

	// encode into this buffer
	enc := json.NewEncoder(&buf)

	if err := enc.Encode(resp); err != nil {
		log.Fatalf("Error: can't encode - %s", err)
	}

	fmt.Printf("Ecoded json response: %v\n", &buf)
}
