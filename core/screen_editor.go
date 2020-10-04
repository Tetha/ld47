package core

import (
	"fmt"

	"github.com/faiface/pixel"
)

var _ (Screen) = (*EditLevelScreen)(nil)

type EditLevelScreen struct {
	systems *Systems
	batch   *pixel.Batch

	toolbox   *ToolboxComponent
	levelGrid *LevelGridComponent
}

func NewEditLevelScreen(systems *Systems) *EditLevelScreen {
	systems.MainScreen.UpperButtonText.Clear()
	fmt.Fprint(systems.MainScreen.UpperButtonText, ">> Run! >>")

	systems.MainScreen.DescriptionText.Clear()
	fmt.Fprint(systems.MainScreen.DescriptionText, systems.level.Description)

	result := &EditLevelScreen{
		systems: systems,
		batch:   pixel.NewBatch(&pixel.TrianglesData{}, systems.Sprites.Sheet),

		toolbox:   NewToolboxComponent(systems),
		levelGrid: NewLevelGridComponent(systems),
	}

	result.toolbox.ResetBaseMatrix(pixel.IM.Moved(pixel.V(875, 400)))
	result.levelGrid.ResetBaseMatrix(pixel.IM.Moved(pixel.V(25, 25)))
	result.levelGrid.ShowGhostDirections = true

	systems.simulation = &SimulationState{}
	for _, initialPosition := range systems.level.InitialGhostPositions {
		systems.simulation.CurrentGhostPositions = append(systems.simulation.CurrentGhostPositions, initialPosition.Clone())
	}
	return result
}

func (screen *EditLevelScreen) Run(target pixel.Target, _ float64) GameState {
	screen.batch.Clear()
	screen.toolbox.Draw(screen.batch)
	screen.levelGrid.DrawLevelGrid(screen.batch, 0)
	screen.batch.Draw(target)
	return GameStateKeep
}

func (screen *EditLevelScreen) Click(pos pixel.Vec) GameState {
	inputSystem := screen.systems.input
	if pixel.R(875, 400, 1000, 800).Contains(pos) {
		// toolbox clicked
		screen.toolbox.SelectItem(pos)
	}

	if pixel.R(0, 0, 800, 800).Contains(pos) {
		// screen
		tileX, tileY, preset, tile := screen.levelGrid.GetTile(pos)
		if inputSystem.SelectedTile == nil {
			if tile != nil && !preset {
				for idx, placedTile := range inputSystem.PlacedTools {
					if placedTile.x == tileX && placedTile.y == tileY {
						placedTile.content.Mark(false)
						inputSystem.RemainingTools = append(inputSystem.RemainingTools, placedTile.content)
						inputSystem.PlacedTools = append(inputSystem.PlacedTools[:idx], inputSystem.PlacedTools[idx+1:]...)
					}
				}
			}
		} else {

			if tile == nil {
				for idx, toolboxTile := range inputSystem.RemainingTools {
					if toolboxTile == inputSystem.SelectedTile {
						inputSystem.RemainingTools = append(inputSystem.RemainingTools[:idx], inputSystem.RemainingTools[idx+1:]...)
					}
				}
				inputSystem.PlacedTools = append(inputSystem.PlacedTools, TileDefinition{
					x:       tileX,
					y:       tileY,
					content: inputSystem.SelectedTile,
				})
				inputSystem.SelectedTile.Mark(true)
				inputSystem.SelectedTile = nil
			}
		}
	}
	screen.levelGrid.HighlightEmpty = inputSystem.SelectedTile != nil

	if screen.systems.MainScreen.UpperButtonBounds.Contains(pos) {
		return GameStateSimulation
	}
	return GameStateKeep
}
