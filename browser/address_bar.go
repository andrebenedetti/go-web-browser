package browser

// THIS ENTIRE PACKAGE IS A WIP AND NEEDS MAJOR REFACTORING

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type AddressBarHandler struct {
	r Renderer
}

func (a *AddressBarHandler) createAddressBar(onSubmit func(url string, r Renderer)) fyne.CanvasObject {
	entry := widget.NewEntry()
	return &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Go to:", Widget: entry}},
		OnSubmit: func() {
			onSubmit(entry.Text, a.r)
		},
	}
}
