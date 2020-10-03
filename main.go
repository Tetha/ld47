package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/tetha/ld47/core"
	"golang.org/x/image/colornames"
)

type GameState uint

const (
	GameStateNone = iota
	GameStateTileTest
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Yay pixel",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	systems, err := core.InitSystems()
	if err != nil {
		panic(err)
	}
	state := GameStateTileTest
	lastState := GameStateNone
	var screen core.Screen
	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		stateChanged := (state != lastState)
		lastState = state

		win.Clear(colornames.Black)
		switch state {
		case GameStateTileTest:
			if stateChanged {
				screen = core.NewTileTestScreen(systems)
			}
			screen.Run(win, dt)
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
