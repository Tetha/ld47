package core

import (
	"math"

	"github.com/faiface/pixel"
)

type LevelGridComponent struct {
	systems             *Systems
	baseMatrix          pixel.Matrix
	ShowGhostDirections bool
	HighlightEmpty      bool
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
	usedTiles := make([]TileDefinition, 0)
	for _, tile := range levelGrid.systems.level.PresetTiles {
		tile.content.Draw(levelGrid.systems.Sprites, target, levelGrid.baseMatrix.Moved(pixel.V(float64(tile.x*48), float64(tile.y*48))))
		usedTiles = append(usedTiles, tile)
	}

	for _, tile := range levelGrid.systems.input.PlacedTools {
		tile.content.Draw(levelGrid.systems.Sprites, target, levelGrid.baseMatrix.Moved(pixel.V(float64(tile.x*48), float64(tile.y*48))))
		usedTiles = append(usedTiles, tile)
	}

	for idx := range levelGrid.systems.simulation.CurrentGhostPositions {
		currentPosition := levelGrid.systems.simulation.CurrentGhostPositions[idx]
		if currentPosition.ascended {
			continue
		}
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
		ghostTile.SetDirection(currentPosition.direction)
		ghostTile.displayDirection = levelGrid.ShowGhostDirections

		ghostTile.Draw(levelGrid.systems.Sprites, target, levelGrid.baseMatrix.Moved(drawPosition))
		usedTiles = append(usedTiles, TileDefinition{x: currentPosition.x, y: currentPosition.y})
	}

	if levelGrid.HighlightEmpty {
		emptyTile := TileEmpty{Highlight: true}
		for x := uint(0); x < 16; x++ {
			for y := uint(0); y < 16; y++ {
				used := false
				for _, tile := range usedTiles {
					if tile.x == x && tile.y == y {
						used = true
						break
					}
				}
				if !used {
					emptyTile.Draw(levelGrid.systems.Sprites, target, levelGrid.baseMatrix.Moved(pixel.V(float64(48*x), float64(y*48))))
				}
			}
		}
	}
}

func (grid *LevelGridComponent) GetTile(position pixel.Vec) (uint, uint, bool, Tile) {
	componentPos := grid.baseMatrix.Unproject(position)

	tileX := uint(math.Floor((componentPos.X + 24) / 48))
	tileY := uint(math.Floor((componentPos.Y + 24) / 48))

	for _, tile := range grid.systems.level.PresetTiles {
		if tile.x == tileX && tile.y == tileY {
			return tileX, tileY, true, tile.content
		}
	}

	for _, tile := range grid.systems.level.InitialGhostPositions {
		if tile.x == tileX && tile.y == tileY {
			return tileX, tileY, true, NewTileGhost(tile.id)
		}
	}

	for _, tile := range grid.systems.input.PlacedTools {
		if tile.x == tileX && tile.y == tileY {
			return tileX, tileY, false, tile.content
		}
	}
	return tileX, tileY, false, nil
}
