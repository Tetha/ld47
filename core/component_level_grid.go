package core

import (
	"github.com/faiface/pixel"
)

type LevelGridComponent struct {
	systems    *Systems
	baseMatrix pixel.Matrix
}

func NewLevelGridComponent(systems *Systems) *LevelGridComponent {
	return &LevelGridComponent{
		systems:    systems,
		baseMatrix: pixel.IM,
	}
}

func (levelGrid *LevelGridComponent) ResetBaseMatrix(baseMatrix pixel.Matrix) {
	levelGrid.baseMatrix = baseMatrix
}

func (levelGrid *LevelGridComponent) DrawLevelGrid(target pixel.Target, tickPercentage float64) {
	for _, tile := range levelGrid.systems.level.PresetTiles {
		tile.content.Draw(levelGrid.systems.sprites, target, levelGrid.baseMatrix.Moved(pixel.V(float64(tile.x*48), float64(tile.y*48))))
	}

	for idx := range levelGrid.systems.simulation.CurrentGhostPositions {
		currentPosition := levelGrid.systems.simulation.CurrentGhostPositions[idx]
		//spew.Dump("In draw", currentPosition)
		nextPosition := currentPosition
		if idx < len(levelGrid.systems.simulation.NextGhostPositions) {
			nextPosition = levelGrid.systems.simulation.NextGhostPositions[idx]
		}

		deltaPosX := float64(nextPosition.x) - float64(currentPosition.x)
		deltaPosY := float64(nextPosition.y) - float64(currentPosition.y)

		drawPosition := pixel.V(
			48*(float64(currentPosition.x)+deltaPosX*tickPercentage),
			48*(float64(currentPosition.y)+deltaPosY*tickPercentage))

		ghostTile := NewTileGhost(currentPosition.id)
		ghostTile.SetInventory(currentPosition.inventory)
		ghostTile.Draw(levelGrid.systems.sprites, target, levelGrid.baseMatrix.Moved(drawPosition))

	}
}
