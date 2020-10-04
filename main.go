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
	GameStateNone GameState = iota
	GameStateTileTest
	GameStateSimulationTest
	GameStateSimulationTest2
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

		if win.Pressed(pixelgl.KeyEscape) {
			return
		}

		if win.Pressed(pixelgl.KeyF1) {
			state = GameStateTileTest
		}

		if win.Pressed(pixelgl.KeyF2) {
			state = GameStateSimulationTest
		}
		if win.Pressed(pixelgl.KeyF3) {
			state = GameStateSimulationTest2
		}

		switch state {
		case GameStateTileTest:
			if stateChanged {
				screen = core.NewTileTestScreen(systems)
			}
		case GameStateSimulationTest:
			if stateChanged {
				systems.SetLevel(&core.TestLevelOne)
				screen = core.NewSimulationScreen(systems)
			}
		case GameStateSimulationTest2:
			if stateChanged {
				systems.SetLevel(&core.TestLevelTwo)
				screen = core.NewSimulationScreen(systems)
			}
		}
		screen.Run(win, dt)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
