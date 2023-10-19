package browser

// THIS ENTIRE PACKAGE IS A WIP AND NEEDS MAJOR REFACTORING

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/andrebenedetti/go-web-browser/dom"
)

type Renderer struct {
	w fyne.Window
}

func (r *Renderer) NewRenderer(window fyne.Window) *Renderer {
	return &Renderer{w: window}
}

func (r *Renderer) RenderPage(tree dom.Tree) {

	texts := make([]*canvas.Text, 1)

	unvisited := []*dom.Element{tree.Root}
	for len(unvisited) > 0 {
		current := unvisited[0]
		unvisited = unvisited[1:]
		if current == nil {
			continue
		}
		for _, child := range current.Children {
			if child != nil {

				unvisited = append([]*dom.Element{child}, unvisited...)
			}
		}
		if current.ElementType == "text" {
			texts = append(texts, canvas.NewText(current.Text, color.Black))
		}
	}

	// grid := container.New(layout.NewGridLayout(len(texts)), texts[0], texts[1], texts[2])
	// r.w.SetContent(grid)

}
