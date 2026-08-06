package main

import (
	"flag"
	"time"

	atmt "github.com/T117m/sca/automata"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	FSCRN  bool
	WIDTH  uint
	HEIGHT uint
	SCALE  int32
	TICK   time.Duration
	B, S   = []uint8{3}, []uint8{2, 3}
)

func init() {
	var (
		w, h  = flag.Int("w", 100, "Width"), flag.Int("h", 50, "Height")
		ms    = flag.Int("ms", 117, "Milliseconds per tick")
		scale = flag.Int("scale", 15, "Cell scale in pixels")
		fscrn = flag.Bool("f", false, "Toggle fullscreen (overrides w and h)")
	)

	flag.Parse()

	WIDTH, HEIGHT, TICK, FSCRN = uint(*w), uint(*h), time.Duration(*ms), *fscrn

	if *scale > 0 {
		SCALE = int32(*scale)
	} else {
		panic("Scale must be bigger than 0")
	}

	if len(flag.Args()) >= 1 {
		if newB, newS, ok := atmt.ValidateRule(flag.Arg(0)); ok {
			B, S = newB, newS
		} else {
			panic("Invalid rule given. Consult the README.md")
		}
	}
}

func main() {
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(int32(WIDTH)*SCALE, int32(HEIGHT)*SCALE, "sca")
	defer rl.CloseWindow()

	if FSCRN {
		rl.ToggleBorderlessWindowed()
		WIDTH = uint(rl.GetScreenWidth() / int(SCALE))
		HEIGHT = uint(rl.GetScreenHeight() / int(SCALE))
	}

	a := atmt.NewAutomata(uint(WIDTH), uint(HEIGHT), B, S)

	if WIDTH >= 3 && HEIGHT >= 3 {
		defaultGlider(a)
	}

	rl.SetTargetFPS(144)

	go func() {
		for !rl.WindowShouldClose() {
			time.Sleep(TICK * time.Millisecond)
			a.Update()
		}
	}()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		for x := range WIDTH {
			for y := range HEIGHT {
				if a.At(x, y) {
					xPos, yPos := int32(x)*SCALE, int32(y)*SCALE
					rl.DrawRectangle(xPos, yPos, SCALE, SCALE, rl.White)
				}
			}
		}

		rl.EndDrawing()
	}
}

func defaultGlider(a *atmt.Automata) {
	a.Insert(1, 0)
	a.Insert(2, 1)
	a.Insert(2, 2)
	a.Insert(1, 2)
	a.Insert(0, 2)
}
