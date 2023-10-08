package renderer

import (
	"fmt"

	"github.com/andrebenedetti/go-web-browser/http_client"
)

func Render(r http_client.Response) {
	inAngle := false
	for _, c := range r.Body {
		if c == '<' {
			inAngle = true
		} else if c == '>' {
			inAngle = false
		} else if !inAngle {
			fmt.Printf("%c", c)
		}
	}
}
