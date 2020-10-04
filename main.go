package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/tetha/ld47/core"
	"golang.org/x/image/colornames"
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
	state := core.GameStateTileTest
	lastState := core.GameStateNone
	var screen core.Screen
	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		stateChanged := (state != lastState)
		lastState = state

		win.Clear(colornames.Black)
		systems.Sprites.MaskSprites[core.MaskGeneric].Draw(win, pixel.IM.Moved(pixel.V(1024/2, 768/2)))
		systems.MainScreen.UpperButtonText.Draw(win, pixel.IM)
		systems.MainScreen.LowerButtonText.Draw(win, pixel.IM)

		if win.Pressed(pixelgl.KeyEscape) {
			return
		}

		if win.Pressed(pixelgl.KeyF1) {
			state = core.GameStateTileTest
		}

		if win.Pressed(pixelgl.KeyF2) {
			state = core.GameStateSimulationTest
		}
		if win.Pressed(pixelgl.KeyF3) {
			state = core.GameStateSimulationTest2
		}
		if win.Pressed(pixelgl.KeyF4) {
			state = core.GameStateEditTest1
		}

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			newState := screen.Click(win.MousePosition())
			if newState != core.GameStateKeep {
				state = newState
			}
		}

		switch state {
		case core.GameStateTileTest:
			if stateChanged {
				screen = core.NewTileTestScreen(systems)
			}
		case core.GameStateSimulationTest:
			if stateChanged {
				systems.SetLevel(&core.TestLevelOne)
				screen = core.NewSimulationScreen(systems)
			}
		case core.GameStateSimulationTest2:
			if stateChanged {
				systems.SetLevel(&core.TestLevelTwo)
				screen = core.NewSimulationScreen(systems)
			}
		case core.GameStateEditTest1:
			if stateChanged {
				systems.SetLevel(&core.TestLevelOne)
				state = core.GameStateEdit
			}
		case core.GameStateEdit:
			if stateChanged {
				screen = core.NewEditLevelScreen(systems)
			}

		case core.GameStateSimulation:
			if stateChanged {
				screen = core.NewSimulationScreen(systems)
			}
		}
		newState := screen.Run(win, dt)
		if newState != core.GameStateKeep {
			state = newState
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
