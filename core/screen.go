package core

import (
	"github.com/faiface/pixel"
)

type GameState uint

const (
	GameStateNone GameState = iota
	GameStateKeep
	GameStateTileTest
	GameStateSimulationTest
	GameStateSimulationTest2
	GameStateEditTest1
	GameStateSimulation
	GameStateEdit
	GameStateLevelSelect
	GameStateQuit
)

type Screen interface {
	Run(target pixel.Target, dt float64) GameState
	Click(position pixel.Vec) GameState
}
