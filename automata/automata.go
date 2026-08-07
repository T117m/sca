package automata

import "math/rand"

type Automata struct {
	wrapped       bool
	width, height uint
	cells         []bool
	b, s          []uint8
}

func NewAutomata(width, height uint, b, s []uint8, isWrapped bool) *Automata {
	cells := make([]bool, width*height)

	return &Automata{isWrapped, width, height, cells, b, s}
}

func (a *Automata) At(x, y int) bool {
	w, h := int(a.width), int(a.height)

	if y*w+x >= len(a.cells) || x >= w || y >= h || x < 0 || y < 0 {
		if !a.wrapped {
			return false
		}

		if x >= w {
			x = 0
		}
		if y >= h {
			y = 0
		}

		if x < 0 {
			x = w-1
		}
		if y < 0 {
			y = h-1
		}
	}

	return a.cells[y*w+x]
}

func (a *Automata) Insert(x, y uint) {
	coords := y*a.width + x

	if x >= a.width || y >= a.height || coords >= uint(len(a.cells)) {
		return
	}

	a.cells[coords] = true
}

func (a *Automata) Kill(x, y uint) {
	coords := y*a.width + x

	if x >= a.width || y >= a.height || coords >= uint(len(a.cells)) {
		return
	}

	a.cells[coords] = false
}

func (a *Automata) Update() {
	newCells := make([]bool, len(a.cells))

	for y := range a.height {
		for x := range a.width {
			c, n := y*a.width+x, a.checkNeighbors(x, y)
			newCells[c] = applyRule(a.cells[c], n, a.b, a.s)
		}
	}

	a.cells = newCells
}

func (a *Automata) checkNeighbors(x, y uint) uint8 {
	var (
		n uint8
		sX, sY = int(x), int(y)
	)

	if a.At(sX-1, sY+1) {
		n++
	}
	if a.At(sX, sY+1) {
		n++
	}
	if a.At(sX+1, sY+1) {
		n++
	}
	if a.At(sX-1, sY) {
		n++
	}
	if a.At(sX+1, sY) {
		n++
	}
	if a.At(sX-1, sY-1) {
		n++
	}
	if a.At(sX, sY-1) {
		n++
	}
	if a.At(sX+1, sY-1) {
		n++
	}

	return n
}

func (a *Automata) RandomFill() {
	for y := range a.height {
		for x := range a.width {
			a.cells[y*a.width+x] = rand.Intn(2) > 0
		}
	}
}

func (a *Automata) RandomCenter() {
	for y := a.height * 2 / 5; y < a.height*3/5; y++ {
		for x := a.width * 2 / 5; x < a.width*3/5; x++ {
			a.cells[y*a.width+x] = rand.Intn(2) > 0
		}
	}
}
