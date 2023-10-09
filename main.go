package main

import (
	"log"

	"github.com/andrebenedetti/go-web-browser/http_client"
	"github.com/andrebenedetti/go-web-browser/renderer"
)

func main() {
	url := "file://go-web-browser"
	_, err := http_client.RetrieveUrl(url)
	if err != nil {
		log.Fatalf("Error visting %s: %s", url, err.Error())
	}

	renderer.Start()
}
