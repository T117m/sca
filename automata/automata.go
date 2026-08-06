package automata

import "math/rand"

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
			c, n := y*a.width+x, a.checkNeighbors(x, y)
			newCells[c] = applyRule(a.cells[c], n, a.b, a.s)
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

func (a *Automata) RandomFill() {
	for y := range a.height {
		for x := range a.width {
			if rand.Intn(2) > 0 {
				a.cells[y*a.width+x] = true
			}
		}
	}
}

func (a *Automata) RandomCenter() {
	for y := a.height * 2 / 5; y < a.height*3/5; y++ {
		for x := a.width * 2 / 5; x < a.width*3/5; x++ {
			if rand.Intn(2) > 0 {
				a.cells[y*a.width+x] = true
			}
		}
	}
}
