package main

import (
	"flag"
	"time"

	atmt "github.com/T117m/sca/automata"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	FSCRN, RF, RC, PAUSE bool
	WIDTH, HEIGHT        uint

	SCALE int32
	TICK  time.Duration
	B, S  = []uint8{3}, []uint8{2, 3}
)

func init() {
	var (
		w, h  = flag.Int("w", 100, "Width"), flag.Int("h", 50, "Height")
		ms    = flag.Int("ms", 117, "Milliseconds per tick")
		scale = flag.Int("scale", 15, "Cell scale in pixels")
		fscrn = flag.Bool("f", false, "Toggle fullscreen (overrides w and h)")
		rc    = flag.Bool("rc", false, "Fill 1/5 of the automata in the center with random noise")
		rf    = flag.Bool("rf", false, "Fill automata with random noise")
		p     = flag.Bool("p", false, "Start paused")
	)

	flag.Parse()

	FSCRN, RF, RC, PAUSE = *fscrn, *rf, *rc, *p
	WIDTH, HEIGHT = uint(*w), uint(*h)
	TICK = time.Duration(*ms)

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

	if RC {
		a.RandomCenter()
	}
	if RF {
		a.RandomFill()
	}

	rl.SetTargetFPS(144)

	go func() {
		for !rl.WindowShouldClose() {
			time.Sleep(TICK * time.Millisecond)
			if !PAUSE {
				a.Update()
			}
		}
	}()

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeySpace) {
			PAUSE = !PAUSE
		}

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
