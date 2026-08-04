package main

import (
	"flag"
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var (
		w, h  = flag.Int("w", 100, "Width"), flag.Int("h", 50, "Height")
		ms    = flag.Int("ms", 117, "Milliseconds per tick")
		scale = flag.Int("scale", 15, "Cell scale in pixels")
		fscrn = flag.Bool("f", false, "Toggle fullscreen (overrides w and h)")
		b, s  = "3", "23"
	)

	flag.Parse()
	width, height, tick := *w, *h, time.Duration(*ms)
	if *scale > 0 {
		SCALE = int32(*scale)
	} else {
		fmt.Println("Scale must be bigger than 0")
		return
	}

	if len(flag.Args()) >= 1 {
		if newB, newS, ok := validateRule(flag.Arg(0)); ok {
			b, s = newB, newS
		} else {
			fmt.Println("Invalid rule given. Consult the README.md")
			return
		}
	}

	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(int32(width)*SCALE, int32(height)*SCALE, "sca")
	defer rl.CloseWindow()

	if *fscrn {
		rl.ToggleBorderlessWindowed()
		width = rl.GetScreenWidth() / *scale
		height = rl.GetScreenHeight() / *scale
	}

	g := NewGrid(uint(width), uint(height))
	if width >= 3 && height >= 3 {
		defaultGlider(g)
	}

	rl.SetTargetFPS(144)

	go func() {
		for !rl.WindowShouldClose() {
			time.Sleep(tick * time.Millisecond)
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
