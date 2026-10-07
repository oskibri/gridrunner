package grid

import "fmt"

type Grid struct {
	Width, Height, CellSize 	int
	Rows, Cols					[]int
	cells						[][]bool
}

func New(w, h, cs int) *Grid {
	cells := make([][]bool, h)
	for i := range cells {
		cells[i] = make([]bool, w)
	}

	rows := make([]int, w / cs)
	cols := make([]int, w / cs)
	//y := make([]int, w / cs)

	add20(rows)
	add20(cols)
	fmt.Println(rows)

	return &Grid{Width: w, Height: h, CellSize: cs, cells: cells, Rows: rows, Cols: cols}
}

func add20(ss []int) {
	for i := range ss {
		if i == 0 {
			ss[i] = 0
			continue
		}
		ss[i] = ss[i-1] + 20
	}
}

func (g *Grid) Set(x, y int, v bool) {
	g.cells[y][x] = v
}

func (g *Grid) Get(x, y int) bool {
	return g.cells[y][x]
}

func (g *Grid) Getcells() [][]bool {
	return g.cells
}
