package main

import (
	"os"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	width, height := 100, 50
	if len(os.Args) > 1 {
		w, err := strconv.Atoi(os.Args[1])
		if err == nil {
			width = w
		}
	}
	if len(os.Args) > 2 {
		h, err := strconv.Atoi(os.Args[2])
		if err == nil {
			height = h
		}
	}

	GRID := NewGrid(uint(width), uint(height))
	GRID.Insert(4, 2)
	GRID.Insert(5, 3)
	GRID.Insert(5, 4)
	GRID.Insert(4, 4)
	GRID.Insert(3, 4)

	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(int32(width*SCALE), int32(height*SCALE), "sca (Simple Cellular Automata)")
	defer rl.CloseWindow()

	rl.SetTargetFPS(144)

	go func() {
		for !rl.WindowShouldClose() {
			time.Sleep(150 * time.Millisecond)
			GRID.Update()
		}
	}()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)
		DrawGrid(GRID)
		//rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}
}
