package browser

// THIS ENTIRE PACKAGE IS A WIP AND NEEDS MAJOR REFACTORING

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"github.com/andrebenedetti/go-web-browser/dom"
	"github.com/andrebenedetti/go-web-browser/http_client"
)

func handleAddressBarSearch(url string, r Renderer) {
	res, _ := http_client.RetrieveUrl(url)
	tree := dom.NewTree()
	fmt.Printf("%s", url)
	dom.ParseBody(res.Body, tree)
	dom.PrintTree(tree.Root, 2)

	r.RenderPage(*tree)
}

func Run() {
	a := app.New()
	window := a.NewWindow("Go Web Browser")
	r := Renderer{w: window}
	addressBarHandler := AddressBarHandler{r}
	window.SetContent(addressBarHandler.createAddressBar((handleAddressBarSearch)))
	window.ShowAndRun()
}
