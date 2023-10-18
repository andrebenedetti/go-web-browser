package browser

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func onSubmitAddressBar() {
}

func Run() {
	a := app.New()
	window := a.NewWindow("Go Web Browser")
	entry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Go to:", Widget: entry}},
		OnSubmit: onSubmitAddressBar,
	}
	window.SetContent(form)
	window.ShowAndRun()
}
