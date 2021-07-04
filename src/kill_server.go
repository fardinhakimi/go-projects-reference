package main 

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strconv"
)

func killServer(fileUrl string) (int, error) {

	fileData, err := ioutil.ReadFile(fileUrl)

	if err != nil {
		return 0, errors.Wrap(err, "can not open configuration file")
	}

	pidid, err := strconv.Atoi(string(fileData))

	if err != nil {
		return 0, errors.Wrap(err, "Error thrown when converting string to int")
	}
	
	return pidid, nil
}


func main() {

	pid, err := killServer("./pidid.txt")

	if err != nil {
		fmt.Printf("Failed to kill the server: Error: %v\n", err )
	} else {
		fmt.Println("Killing process with <pid> %v\n", pid)
	}
}