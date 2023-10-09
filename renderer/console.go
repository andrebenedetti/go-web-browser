package renderer

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// func Render(r http_client.Response) {
// 	inAngle := false
// 	for _, c := range r.Body {
// 		if c == '<' {
// 			inAngle = true
// 		} else if c == '>' {
// 			inAngle = false
// 		} else if !inAngle {
// 			fmt.Printf("%c", c)
// 		}
// 	}
// }

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func Start() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go Web Browser!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
