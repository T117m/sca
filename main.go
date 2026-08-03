package main

import (
	"fmt"
	"os"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	width, height := 100, 50
	b, s := "3", "23"

	if len(os.Args) > 1 {
		if newB, newS, ok := validateRule(os.Args[1]); ok {
			b, s = newB, newS
		} else {
			fmt.Println("Invalid rule given")
			return
		}
	}

	g := NewGrid(uint(width), uint(height))
	if width >= 3 && height >= 3 {
		defaultGlider(g)
	}

	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(int32(width*SCALE), int32(height*SCALE), "sca (Simple Cellular Automata)")
	defer rl.CloseWindow()

	rl.SetTargetFPS(144)

	go func() {
		for !rl.WindowShouldClose() {
			time.Sleep(117 * time.Millisecond)
			g.Update(b, s)
		}
	}()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		DrawGrid(g)
		rl.EndDrawing()
	}
}

func defaultGlider(g *Grid) {
	g.Insert(1, 0)
	g.Insert(2, 1)
	g.Insert(2, 2)
	g.Insert(1, 2)
	g.Insert(0, 2)
}
