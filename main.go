/*
Program:     main
Author:      Oskar Voldbakken Hesle
Date:        2026-10-05
*/

package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	cellSize	= 20
	cols		= 32
	rows		= 32
)

type Point struct {
	x, y float32
}

var (
	Up		= Point{0, -1}
	Down	= Point{0, 1}
	Left	= Point{-1, 0}
	Right	= Point{1, 0}
)

type Game struct {
	canvas *ebiten.Image
	head Point
}

func (g *Game) Move(p Point) {
	g.head.x += p.x
	g.head.y += p.y
}

func newGame() *Game {
	return &Game {
		canvas: ebiten.NewImage(640, 640),
		head: Point{0, 0},
	}
}

func drawCell(screen *ebiten.Image, p Point, cellColor color.Color) {
	//cellColor := color.RGBA{255, 0, 0, 255}
	px := p.x * cellSize
	py := p.y * cellSize
	vector.FillRect(screen, px+1, py+1, cellSize-1, cellSize-1, cellColor, false)
}

func getCursorPos(cx, cy int) Point {
	return Point{float32(cx / cellSize), float32(cy / cellSize)}
}

func drawGrid(screen *ebiten.Image) {
	lineColor := color.RGBA{40, 40, 40, 255}
	for x := range cols + 1 {
		px := float32(x * cellSize)
		vector.StrokeLine(screen, px, 0, px, 640, 1, lineColor, false)
	}
	for y := range rows + 1 {
		py := float32(y * cellSize)
		vector.StrokeLine(screen, 0, py, 640, py, 1, lineColor, false)
	}
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		drawCell(g.canvas, getCursorPos(ebiten.CursorPosition()), color.White)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight){
		g.Move(Right)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft){
		g.Move(Left)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp){
		g.Move(Up)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Move(Down)
	}
	g.head.x = max(0, min(cols-1, g.head.x))
	g.head.y = max(0, min(rows-1, g.head.y))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.canvas, nil)
	drawGrid(screen)
	drawCell(screen, g.head, color.RGBA{255, 0, 0, 255})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", g.head))
	x, y := ebiten.CursorPosition()
	x = x / 20
	y = y / 20
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d, %d", x, y), 600, 0)

}

func (g *Game) Layout(outsideW, outsideH int) (int, int) {
	return 640, 640
}

func main() {
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Gridrunner")
	ebiten.SetTPS(30)
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
