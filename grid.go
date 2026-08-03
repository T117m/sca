package main

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const SCALE = 15

type (
	coords struct {
		x, y uint
	}

	Grid struct {
		width, height uint
		lock          sync.RWMutex
		cells         map[coords]bool
	}
)

func NewGrid(w, h uint) *Grid {
	cells := make(map[coords]bool, h*w)
	for x := range w {
		for y := range h {
			crd := coords{x, y}

			cells[crd] = false
		}
	}

	return &Grid{w, h, sync.RWMutex{}, cells}
}

func (g *Grid) At(x, y uint) bool {
	if x >= g.width || y >= g.height {
		return false
	}

	return g.cells[coords{x, y}]
}

func (g *Grid) Insert(x, y uint) {
	if x > g.width || y > g.height {
		return
	}

	g.cells[coords{x, y}] = true
}

func (g *Grid) Update(b, s string) {
	g.lock.Lock()
	defer g.lock.Unlock()

	newCells := make(map[coords]bool, len(g.cells))

	for crd := range g.cells {
		newCells[crd] = applyRule(g.cells[crd], g.checkNeighbors(crd), b, s)
	}

	g.cells = newCells
}

func (g *Grid) checkNeighbors(crd coords) uint8 {
	var n uint8

	if g.At(crd.x-1, crd.y+1) {
		n++
	}
	if g.At(crd.x, crd.y+1) {
		n++
	}
	if g.At(crd.x+1, crd.y+1) {
		n++
	}
	if g.At(crd.x-1, crd.y) {
		n++
	}
	if g.At(crd.x+1, crd.y) {
		n++
	}
	if g.At(crd.x-1, crd.y-1) {
		n++
	}
	if g.At(crd.x, crd.y-1) {
		n++
	}
	if g.At(crd.x+1, crd.y-1) {
		n++
	}

	return n
}

func DrawGrid(g *Grid) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	for x := range g.width {
		for y := range g.height {
			if g.cells[coords{x, y}] {
				rl.DrawRectangle(int32(x)*SCALE, int32(y)*SCALE, SCALE, SCALE, rl.White)
			}
		}
	}
}
