# sca - Simple Cellular Automata

A simple program for launching 2D cellular automata with custom rulestrings and
dimensions that can be paused and edited. The default rulestring is b3/s23 with
dimensions 100 by 50 cells, 117 miliseconds per tick and 15 pixels cell scale.

> [!Warning]
> This is WIP

#### TODO:

- [x] Add width and height flags
- [x] Add fullscreen
- [ ] Add ability to launch multiple automata with different rulestrings
- [ ] Add some kind of random noise generation
- [ ] Add pause/resume
- [ ] Add editing
- [ ] Add no-borders flag (wrap around)
- [ ] Add bg and cell colors flags
- [ ] Make a man page

### Building from source

#### Requirements

[Go 1.26.5 or later](https://go.dev/dl/)\
[raylib-go](https://github.com/gen2brain/raylib-go)\
GNU Make

#### Building

```sh
make sca
```
The compiled binary should be located in ./build/

### Usage

```sh
sca [flags] <rulestring>
```

The list of available flags can be accessed through ```sca --help```.
The rulestring is expected to be in a format of "b#/s#", where:

\# - digits 0-8. The b and s fields do not necessarly require one digit, they
can accept multiple as well as none at all;\
b - amount(s) of living neighbors needed for a dead cell to become alive;\
/ - a separator between birth and survival;\
s - amount(s) of living neighbors needed for a living cell to survive to the
next tick.

Here are some examples of valid rulestrings:

- b3/s23 - famous Conway's Game of Life;
- b3/s234 - carykh's "Ant Colony";
- b/s - Yes, this is still valid. Every cell will die tho;
- B345/S2 - [Adam P. Goucher's](https://catagolue.hatsya.com/home) "Blinkers".
Capital letters are also supported.

---

This project is heavily inspired by
[this video](https://www.youtube.com/watch?v=QK_KZv-YyOc) by
[carykh](https://www.youtube.com/@carykh).
