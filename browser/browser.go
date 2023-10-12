package browser

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Browser struct {
	Document string
	Font     font.Face
}

func (b *Browser) Update() error {
	return nil
}

func (b *Browser) Draw(screen *ebiten.Image) {

	text.Draw(screen, b.Document, b.Font, 0, 0, color.White)
}

func (b *Browser) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// func Start() {
// 	// Initial tab
// 	url := "example.org"
// 	document, err := http_client.RetrieveUrl(url)
// 	if err != nil {
// 		log.Fatalf("Error visting %s: %s", url, err.Error())
// 	}

// 	// Load font
// 	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	const dpi = 72
// 	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
// 		Size:    24,
// 		DPI:     dpi,
// 		Hinting: font.HintingVertical,
// 	})

// 	if err != nil {
// 		log.Fatal("Failed to initialize fonts")
// 	}

// 	ebiten.SetWindowTitle("Go Web Browser")
// 	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

// 	if err := ebiten.RunGame(&Browser{Document: dom.ParseText(document.Body), Font: mplusNormalFont}); err != nil {
// 		log.Fatal(err)
// 	}
// }
