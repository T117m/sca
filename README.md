# sca - Simple Cellular Automata

A simple program for launching 2D cellular automata with custom rulestrings and
dimensions that can be paused and edited. The default rulestring is b3/s23 with
dimensions 100 by 50 cells, 117 miliseconds per tick and 15 pixels cell scale.

### Building from source

#### Requirements

[Go 1.26.5 or later](https://go.dev/dl/)\
[raylib-go](https://github.com/gen2brain/raylib-go)

#### Building

```sh
go build ./cmd/sca.go
```

### Installaion

#### GNU/Linux

Run as root (with ```doas```, ```sudo```, ```su``` or any other preferred way)
```./install.sh```.

#### WIndows

Start PowersShell as administrator. Can be done with
```ps1
Start-Process powershell -Verb RunAs
```

Navigate to project direcory and run ```.\install.ps1```

### Usage

#### CLI

```sh
sca [flags] [rulestring]
```

The list of available flags and their defaults can be accessed through 
```sca --help``` or ```man sca``` if you're on GNU/Linux. The rulestring is 
expected to be in a format of "b#/s#", where:

**\#** - digits 0-8. The b and s fields do not necessarly require one digit, 
they can accept multiple as well as none at all;\
**b** - amount(s) of living neighbors needed for a dead cell to become alive;\
**/** - a separator between birth and survival. While technically unnecessary, 
made mandatory for aesthetical reasons;\
**s** - amount(s) of living neighbors needed for a living cell to survive to 
the next tick.

Here are some examples of valid rulestrings:

- b3/s23 - famous Conway's Game of Life;
- b3/s234 - carykh's "Ant Colony";
- b/s - Yes, this is still valid. Every cell will die tho;
- B345/S2 - [Adam P. Goucher's](https://catagolue.hatsya.com/home) "Blinkers".
Capital letters are also supported.

Color flags accept either a hex color string (#00000 - #ffffff) or one of these
predefined raylib color options: "beige", "black", "blue", "brown", "dark blue"
, "dark brown", "dark gray", "dark green", "dark purple", "gold", "gray",
"green", "light gray", "lime", "magenta", "maroon", "orange", "pink", "purple",
"red", "sky blue", "violet", "white", "yellow". Capital lettes are supported.

#### Mouse/Keyboard

Use mouse1 to create living cells. Use mouse2 to kill. Press space to 
pause/resume.

---

This project is heavily inspired by
[this video](https://www.youtube.com/watch?v=QK_KZv-YyOc) by
[carykh](https://www.youtube.com/@carykh).
