package main

import (
	"fmt"
	"log"

	"github.com/andrebenedetti/go-web-browser/http_client"
)

func main() {
	url := "example.org"
	res, err := http_client.Request("GET", url)
	if err != nil {
		log.Fatalf("Error visting %s: %s", url, err.Error())
	}

	fmt.Println(res)
}
