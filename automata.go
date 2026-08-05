package main

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var SCALE int32 = 15

type Grid struct {
	width, height uint
	lock          sync.RWMutex
	cells         []bool
}

func NewGrid(w, h uint) *Grid {
	cells := make([]bool, h*w)

	return &Grid{w, h, sync.RWMutex{}, cells}
}

func (g *Grid) At(x, y uint) bool {
	if x >= g.width || y >= g.height {
		return false
	}

	return g.cells[y*g.width+x]
}

func (g *Grid) Insert(x, y uint) {
	if x > g.width || y > g.height {
		return
	}

	g.cells[y*g.width+x] = true
}

func (g *Grid) Update(b, s string) {
	g.lock.Lock()
	defer g.lock.Unlock()

	newCells := make([]bool, len(g.cells))

	for y := range g.height {
		for x := range g.width {
			n := g.checkNeighbors(x, y)
			newCells[y*g.width+x] = applyRule(g.cells[y*g.width+x], n, b, s)
		}
	}

	g.cells = newCells
}

func (g *Grid) checkNeighbors(x, y uint) uint8 {
	var n uint8

	if g.At(x-1, y+1) {
		n++
	}
	if g.At(x, y+1) {
		n++
	}
	if g.At(x+1, y+1) {
		n++
	}
	if g.At(x-1, y) {
		n++
	}
	if g.At(x+1, y) {
		n++
	}
	if g.At(x-1, y-1) {
		n++
	}
	if g.At(x, y-1) {
		n++
	}
	if g.At(x+1, y-1) {
		n++
	}

	return n
}

func DrawGrid(g *Grid) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	for x := range g.width {
		for y := range g.height {
			if g.cells[y*g.width+x] {
				xPos, yPos := int32(x)*SCALE, int32(y)*SCALE
				rl.DrawRectangle(xPos, yPos, SCALE, SCALE, rl.White)
			}
		}
	}
}
