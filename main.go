package main

import (
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/eric-lab-star/ebitGame/res/run"
	ebi "github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  int     = 320
	screenHeight int     = 240
	imgWidth     float64 = 40
	imgHeight    float64 = 64
)

var runnerImg []*ebi.Image

// Game implements ebi.Game interface.
type Game struct {
	count int
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.count++
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebi.Image) {
	op := &ebi.DrawImageOptions{}
	op.GeoM.Translate(float64(screenWidth)/2, float64(screenHeight)/2)
	op.GeoM.Translate(-imgWidth/2, -imgHeight/2)
	screen.Fill(color.White)
	speed := g.count / 8
	index := (speed) % 6
	screen.DrawImage(runnerImg[index], op)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	run.InitImgs()
	for _, b := range run.Imgs {
		img, _, err := image.Decode(b)
		if err != nil {
			log.Printf("err %v\n ", err)
		}

		ebiImg := ebi.NewImageFromImage(img)
		runnerImg = append(runnerImg, ebiImg)
	}

	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.

	ebi.SetWindowTitle("animation")
	// Call ebiten.RunGame to start your game loop.
	if err := ebi.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
