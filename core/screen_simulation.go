package core

import (
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
	result := &SimulationScreen{
		systems: systems,
		batch:   pixel.NewBatch(&pixel.TrianglesData{}, systems.sprites.Sheet),
	}
	result.levelGrid = NewLevelGridComponent(systems)

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

func (screen *SimulationScreen) Run(target pixel.Target, dt float64) {
	screen.timeSinceLastSimulation += dt

	if screen.timeSinceLastSimulation >= timeBetweenSimulationTicks {
		screen.DoPhysicsStep()
		screen.timeSinceLastSimulation = 0
	}

	simulationTickPercentage := screen.timeSinceLastSimulation / timeBetweenSimulationTicks
	screen.batch.Clear()
	screen.levelGrid.DrawLevelGrid(screen.batch, simulationTickPercentage)
	screen.batch.Draw(target)
}

func (screen *SimulationScreen) DoPhysicsStep() {
	state := screen.systems.simulation

	state.CurrentGhostPositions = state.NextGhostPositions
	state.NextGhostPositions = nil

	for idx := range state.CurrentGhostPositions {
		ghostPosition := &state.CurrentGhostPositions[idx]
		for _, tile := range screen.systems.level.PresetTiles {
			if tile.x == ghostPosition.x && tile.y == ghostPosition.y {
				tile.content.ModifyGhostPosition(ghostPosition)
			}
		}

		newNextPosition := GhostPosition{
			x:         ghostPosition.x,
			y:         ghostPosition.y,
			direction: ghostPosition.direction,
			id:        ghostPosition.id,
			inventory: append([]MemoryType(nil), ghostPosition.inventory...),
		}

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
}
