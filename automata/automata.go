package automata

import rl "github.com/gen2brain/raylib-go/raylib"

var SCALE int32 = 15

type Automata struct {
	width, height uint
	cells         []bool
	b, s          []uint8
}

func NewAutomata(w, h uint, b, s []uint8) *Automata {
	cells := make([]bool, h*w)

	return &Automata{w, h, cells, b, s}
}

func (a *Automata) At(x, y uint) bool {
	if x >= a.width || y >= a.height {
		return false
	}

	return a.cells[y*a.width+x]
}

func (a *Automata) Insert(x, y uint) {
	if x > a.width || y > a.height {
		return
	}

	a.cells[y*a.width+x] = true
}

func (a *Automata) Update() {
	newCells := make([]bool, len(a.cells))

	for y := range a.height {
		for x := range a.width {
			n := a.checkNeighbors(x, y)
			newCells[y*a.width+x] = applyRule(a.cells[y*a.width+x], n, a.b, a.s)
		}
	}

	a.cells = newCells
}

func (a *Automata) checkNeighbors(x, y uint) uint8 {
	var n uint8

	if a.At(x-1, y+1) {
		n++
	}
	if a.At(x, y+1) {
		n++
	}
	if a.At(x+1, y+1) {
		n++
	}
	if a.At(x-1, y) {
		n++
	}
	if a.At(x+1, y) {
		n++
	}
	if a.At(x-1, y-1) {
		n++
	}
	if a.At(x, y-1) {
		n++
	}
	if a.At(x+1, y-1) {
		n++
	}

	return n
}

func DrawAutomata(a *Automata) {
	for x := range a.width {
		for y := range a.height {
			if a.cells[y*a.width+x] {
				xPos, yPos := int32(x)*SCALE, int32(y)*SCALE
				rl.DrawRectangle(xPos, yPos, SCALE, SCALE, rl.White)
			}
		}
	}
}
