// main.go
package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/oskibri/gridrunner/grid"
)

type Game struct {
	screenW, screenH int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	grid := grid.New(640, 640, 20)
	grid.Set(10, 10, true)

	//vector.StrokeLine(screen, 0, 20, 640, 20, 1, color.RGBA{40, 40, 40, 255}, false)
	//vector.StrokeLine(screen, 0, 40, 640, 40, 1, color.RGBA{40, 40, 40, 255}, false)
	for y := range grid.Rows {
		vector.StrokeLine(screen, 0, float32(y), 640, float32(y), 1, color.RGBA{40, 40, 40, 255}, false)
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return g.screenW, g.screenH
}

func main() {
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Gridrunner")
	if err := ebiten.RunGame(&Game{
		screenW: 640,
		screenH: 640,
	}); err != nil {
		log.Fatal(err)
	}
}
