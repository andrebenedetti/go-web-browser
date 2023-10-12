package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/andrebenedetti/go-web-browser/dom"
	"github.com/andrebenedetti/go-web-browser/http_client"
)

func main() {
	// browser.Start()

	// tree := dom.NewTree()
	// tree.AddTag("html")
	// tree.AddTag("p")
	// tree.AddText("Lorem ipsum")
	// tree.AddTag("/p")
	// tree.AddTag("/html")
	// dom.PrintTree(tree.Root, 0)
	tree := dom.NewTree()
	fmt.Println(tree)

	res, err := http_client.RetrieveUrl("example.org")
	if err != nil {
		log.Fatal(err)
	}

	body := strings.ReplaceAll(res.Body, "\n", "")
	fmt.Println(body)
	dom.ParseBody(body, tree)

	dom.PrintTree(tree.Root, 0)

}
