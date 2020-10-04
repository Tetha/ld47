package core

import (
	"fmt"

	"github.com/faiface/pixel"
)

var _ (Screen) = (*SimulationScreen)(nil)

type SimulationScreen struct {
	systems *Systems
	batch   *pixel.Batch

	levelGrid *LevelGridComponent

	timeSinceLastSimulation float64
}

type SimulationState struct {
	CurrentGhostPositions []GhostPosition
	NextGhostPositions    []GhostPosition
}

func NewSimulationScreen(systems *Systems) *SimulationScreen {
	systems.MainScreen.UpperButtonText.Clear()
	fmt.Fprint(systems.MainScreen.UpperButtonText, "<< Edit! Go back <<")
	result := &SimulationScreen{
		systems: systems,
		batch:   pixel.NewBatch(&pixel.TrianglesData{}, systems.Sprites.Sheet),
	}
	result.levelGrid = NewLevelGridComponent(systems)
	result.levelGrid.ResetBaseMatrix(pixel.IM.Moved(pixel.V(25, 25)))

	for _, initial := range systems.level.InitialGhostPositions {
		systems.simulation.CurrentGhostPositions = append(systems.simulation.CurrentGhostPositions, GhostPosition{
			x:         initial.x,
			y:         initial.y,
			direction: initial.direction,
			id:        initial.id,
		})
		systems.simulation.NextGhostPositions = append(systems.simulation.NextGhostPositions, GhostPosition{
			x:         initial.x,
			y:         initial.y,
			direction: initial.direction,
			id:        initial.id,
		})
	}

	return result
}

const timeBetweenSimulationTicks = 0.5

func (screen *SimulationScreen) Run(target pixel.Target, dt float64) GameState {

	screen.timeSinceLastSimulation += dt

	if screen.timeSinceLastSimulation >= timeBetweenSimulationTicks {
		allAscended := screen.DoPhysicsStep()
		if allAscended {
			screen.systems.level.Won = true

			screen.systems.MainScreen.DescriptionText.Clear()
			fmt.Fprintln(screen.systems.MainScreen.DescriptionText, `Ayy you win!
			
No there is no
fancy victory
screen.

But the level
name just became
green. That's 
something!`)

			for _, tile := range screen.systems.level.PresetTiles {
				tile.content.Reset()
			}
			for _, tile := range screen.systems.level.Toolbox {
				tile.Reset()
			}
			return GameStateLevelSelect
		}
		screen.timeSinceLastSimulation = 0
	}

	simulationTickPercentage := screen.timeSinceLastSimulation / timeBetweenSimulationTicks
	screen.batch.Clear()
	screen.levelGrid.DrawLevelGrid(screen.batch, simulationTickPercentage)
	screen.batch.Draw(target)
	return GameStateKeep
}

func (screen *SimulationScreen) DoPhysicsStep() bool {
	state := screen.systems.simulation

	state.CurrentGhostPositions = state.NextGhostPositions
	state.NextGhostPositions = nil

	allAscended := true
	for idx := range state.CurrentGhostPositions {
		ghostPosition := &state.CurrentGhostPositions[idx]
		for _, tile := range screen.systems.level.PresetTiles {
			if tile.x == ghostPosition.x && tile.y == ghostPosition.y {
				tile.content.ModifyGhostPosition(ghostPosition)
			}
		}

		for _, tile := range screen.systems.input.PlacedTools {
			if tile.x == ghostPosition.x && tile.y == ghostPosition.y {
				tile.content.ModifyGhostPosition(ghostPosition)
			}
		}

		if ghostPosition.ascended {
			newNextPosition := ghostPosition.Clone()
			state.NextGhostPositions = append(state.NextGhostPositions, newNextPosition)
			continue
		}
		allAscended = false

		newNextPosition := ghostPosition.Clone()

		switch ghostPosition.direction {
		case GhostDirectionDown:
			newNextPosition.y--
		case GhostDirectionUp:
			newNextPosition.y++
		case GhostDirectionLeft:
			newNextPosition.x--
		case GhostDirectionRight:
			newNextPosition.x++
		}

		//fmt.Println("-----------")
		//spew.Dump(ghostPosition, newNextPosition)
		state.NextGhostPositions = append(state.NextGhostPositions, newNextPosition)
	}
	return allAscended
}

func (screen *SimulationScreen) Click(pos pixel.Vec) GameState {
	if screen.systems.MainScreen.UpperButtonBounds.Contains(pos) {
		for _, tile := range screen.systems.level.PresetTiles {
			tile.content.Reset()
		}
		for _, tile := range screen.systems.level.Toolbox {
			tile.Reset()
		}
		return GameStateEdit
	}
	return GameStateKeep
}
