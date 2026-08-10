package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	atmt "github.com/T117m/sca/automata"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	PAUSE, RF, RC bool
	FSCRN, WRAP   bool
	WIDTH, HEIGHT uint

	DEADCOLOR, ALIVECOLOR rl.Color

	SCALE int32
	TICK  time.Duration
	RLSTR = "b3/s23"
	B, S  = []uint8{3}, []uint8{2, 3}
)

func init() {
	var (
		w, h  = flag.Int("w", 100, "Width"), flag.Int("h", 50, "Height")
		ms    = flag.Int("ms", 117, "Milliseconds per tick")
		scale = flag.Int("scale", 15, "Cell scale in pixels")
		dc    = flag.String("color-dead", "black", "Color of a dead cell")
		ac    = flag.String("color-alive", "white", "Color of a living cell")
		fscrn = flag.Bool("f", false, "Toggle fullscreen (overrides w and h)")
		p     = flag.Bool("p", false, "Start paused")
		wrp   = flag.Bool("wrap", false, "Make the board wrap around")
		rf    = flag.Bool("rf", false, "Fill automata with random noise")
		rc    = flag.Bool("rc", false,
			"Fill 1/5 of the automata in the center with random noise")
		ok bool
	)

	flag.Parse()

	FSCRN, RF, RC, PAUSE, WRAP = *fscrn, *rf, *rc, *p, *wrp
	WIDTH, HEIGHT = uint(*w), uint(*h)
	TICK = time.Duration(*ms)

	if *w < 0 || *h < 0 {
		fmt.Fprintln(os.Stderr, "Dimensions must be bigger than 0")
		os.Exit(2)
	}

	if *scale > 0 {
		SCALE = int32(*scale)
	} else {
		fmt.Fprintln(os.Stderr, "Scale must be bigger than 0")
		os.Exit(2)
	}

	argsLen := len(flag.Args())
	if argsLen == 1 {
		RLSTR = flag.Arg(0)
		if newB, newS, ok := atmt.ValidateRule(RLSTR); ok {
			B, S = newB, newS
		} else {
			fmt.Fprintln(os.Stderr,
				"Invalid rule given; consult the README or the man page")
			os.Exit(2)
		}
	} else if argsLen > 1 {
		fmt.Fprintln(os.Stderr, "Too many arguments")
		os.Exit(2)
	}

	DEADCOLOR, ok = getColor(*dc)
	if !ok {
		fmt.Fprintln(os.Stderr, "Invalid dead cell color given")
		os.Exit(2)
	}

	ALIVECOLOR, ok = getColor(*ac)
	if !ok {
		fmt.Fprintln(os.Stderr, "Invalid living cell color given")
		os.Exit(2)
	}
}

func getColor(s string) (rl.Color, bool) {
	switch l := strings.ToLower(s); l {
	case "beige":
		return rl.Beige, true
	case "black":
		return rl.Black, true
	case "blue":
		return rl.Blue, true
	case "brown":
		return rl.Brown, true
	case "dark blue":
		return rl.DarkBlue, true
	case "dark brown":
		return rl.DarkBrown, true
	case "dark gray":
		return rl.DarkGray, true
	case "dark green":
		return rl.DarkGreen, true
	case "dark purple":
		return rl.DarkPurple, true
	case "gold":
		return rl.Gold, true
	case "gray":
		return rl.Gray, true
	case "green":
		return rl.Green, true
	case "light gray":
		return rl.LightGray, true
	case "lime":
		return rl.Lime, true
	case "magenta":
		return rl.Magenta, true
	case "maroon":
		return rl.Maroon, true
	case "orange":
		return rl.Orange, true
	case "pink":
		return rl.Pink, true
	case "purple":
		return rl.Purple, true
	case "red":
		return rl.Red, true
	case "sky blue":
		return rl.SkyBlue, true
	case "violet":
		return rl.Violet, true
	case "white":
		return rl.White, true
	case "yellow":
		return rl.Yellow, true
	default:
		if strings.HasPrefix(l, "#") && len(l) == 7 {
			red := l[1:3]
			grn := l[3:5]
			blu := l[5:]

			r, err := strconv.ParseUint(red, 16, 8)
			if err != nil {
				break
			}

			g, err := strconv.ParseUint(grn, 16, 8)
			if err != nil {
				break
			}

			b, err := strconv.ParseUint(blu, 16, 8)
			if err != nil {
				break
			}

			return rl.NewColor(uint8(r), uint8(g), uint8(b), 255), true
		}
	}

	return rl.Color{}, false
}

func main() {
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(int32(WIDTH)*SCALE, int32(HEIGHT)*SCALE, RLSTR)
	defer rl.CloseWindow()

	if FSCRN {
		rl.ToggleBorderlessWindowed()
		WIDTH = uint(rl.GetScreenWidth() / int(SCALE))
		HEIGHT = uint(rl.GetScreenHeight() / int(SCALE))
	}

	a := atmt.NewAutomata(uint(WIDTH), uint(HEIGHT), B, S, WRAP)

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

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			mX, mY := convertMouseToBoard(rl.GetMousePosition())
			if mX > WIDTH {
				mX = WIDTH
			}
			if mY > HEIGHT {
				mY = HEIGHT
			}
			a.Insert(mX, mY)
		}
		if rl.IsMouseButtonDown(rl.MouseButtonRight) {
			mX, mY := convertMouseToBoard(rl.GetMousePosition())
			if mX > WIDTH {
				mX = WIDTH
			}
			if mY > HEIGHT {
				mY = HEIGHT
			}
			a.Kill(mX, mY)
		}

		rl.BeginDrawing()
		rl.ClearBackground(DEADCOLOR)

		for x := range WIDTH {
			for y := range HEIGHT {
				if a.At(int(x), int(y)) {
					xPos, yPos := int32(x)*SCALE, int32(y)*SCALE
					rl.DrawRectangle(xPos, yPos, SCALE, SCALE, ALIVECOLOR)
				}
			}
		}

		rl.EndDrawing()
	}
}

func convertMouseToBoard(v rl.Vector2) (x, y uint) {
	x = uint(float64(v.X) / float64(SCALE))
	y = uint(float64(v.Y) / float64(SCALE))
	return x, y
}
