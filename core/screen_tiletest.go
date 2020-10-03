package core

import "github.com/faiface/pixel"

var _ (Screen) = (*TileTestScreen)(nil)

type TileTestScreen struct {
	systems     *Systems
	screenBatch *pixel.Batch

	ghostTile              Tile
	ghostTileWithInventory *TileGhost

	simpleArrows  []Tile
	memoryBubbles []Tile
	tJunctions    []Tile

	goalTile Tile
}

func NewTileTestScreen(systems *Systems) *TileTestScreen {
	result := &TileTestScreen{
		systems:                systems,
		screenBatch:            pixel.NewBatch(&pixel.TrianglesData{}, systems.sprites.Sheet),
		ghostTile:              NewTileGhost(GhostBlue),
		ghostTileWithInventory: NewTileGhost(GhostBlue),
		goalTile:               NewTileGoal(),

		simpleArrows: []Tile{
			NewTileSimpleArrow(ArrowDirectionStraightUp),
			NewTileSimpleArrow(ArrowDirectionStraightDown),
			NewTileSimpleArrow(ArrowDirectionStraightRight),
			NewTileSimpleArrow(ArrowDirectionStraightLeft),
			NewTileSimpleArrow(ArrowDirectionLeftTurnFromBottom),
			NewTileSimpleArrow(ArrowDirectionRightTurnFromBottom),
			NewTileSimpleArrow(ArrowDirectionLeftTurnFromTop),
			NewTileSimpleArrow(ArrowDirectionRightTurnFromTop),
			NewTileSimpleArrow(ArrowDirectionLeftTurnFromRight),
			NewTileSimpleArrow(ArrowDirectionRightTurnFromRight),
			NewTileSimpleArrow(ArrowDirectionLeftTurnFromLeft),
			NewTileSimpleArrow(ArrowDirectionRightTurnFromLeft),
		},
		memoryBubbles: []Tile{
			NewTileMemory(GhostBlue, MemoryLove),
			NewTileMemory(GhostBlue, MemoryPet),
			NewTileMemory(GhostBlue, MemorySun),
		},
		tJunctions: []Tile{
			NewTileTJunction(GhostBlue, MemoryLove, TJunctionUp, TJunctionLeftTurn, TJunctionMemoryOnStraight),
			NewTileTJunction(GhostBlue, MemoryLove, TJunctionUp, TJunctionLeftTurn, TJunctionMemoryOnTurn),
			NewTileTJunction(GhostBlue, MemoryLove, TJunctionDown, TJunctionLeftTurn, TJunctionMemoryOnStraight),
			NewTileTJunction(GhostBlue, MemoryLove, TJunctionDown, TJunctionLeftTurn, TJunctionMemoryOnTurn),
			NewTileTJunction(GhostBlue, MemoryPet, TJunctionLeft, TJunctionLeftTurn, TJunctionMemoryOnStraight),
			NewTileTJunction(GhostBlue, MemoryPet, TJunctionLeft, TJunctionLeftTurn, TJunctionMemoryOnTurn),
			NewTileTJunction(GhostBlue, MemorySun, TJunctionRight, TJunctionLeftTurn, TJunctionMemoryOnStraight),
			NewTileTJunction(GhostBlue, MemorySun, TJunctionRight, TJunctionLeftTurn, TJunctionMemoryOnTurn),
			NewTileTJunction(GhostBlue, MemoryLove, TJunctionUp, TJunctionRightTurn, TJunctionMemoryOnStraight),
			NewTileTJunction(GhostBlue, MemoryLove, TJunctionUp, TJunctionRightTurn, TJunctionMemoryOnTurn),
			NewTileTJunction(GhostBlue, MemoryLove, TJunctionDown, TJunctionRightTurn, TJunctionMemoryOnStraight),
			NewTileTJunction(GhostBlue, MemoryLove, TJunctionDown, TJunctionRightTurn, TJunctionMemoryOnTurn),
			NewTileTJunction(GhostBlue, MemoryPet, TJunctionLeft, TJunctionRightTurn, TJunctionMemoryOnStraight),
			NewTileTJunction(GhostBlue, MemoryPet, TJunctionLeft, TJunctionRightTurn, TJunctionMemoryOnTurn),
			NewTileTJunction(GhostBlue, MemorySun, TJunctionRight, TJunctionRightTurn, TJunctionMemoryOnStraight),
			NewTileTJunction(GhostBlue, MemorySun, TJunctionRight, TJunctionRightTurn, TJunctionMemoryOnTurn),
		},
	}
	result.ghostTileWithInventory.SetInventory([]MemoryType{MemoryLove, MemoryPet, MemorySun})
	return result
}

func (screen *TileTestScreen) Run(target pixel.Target, _ float64) {
	screen.screenBatch.Clear()
	screen.ghostTile.Draw(screen.systems.sprites, screen.screenBatch, pixel.IM.Moved(pixel.V(100, 100)))
	screen.ghostTileWithInventory.Draw(screen.systems.sprites, screen.screenBatch, pixel.IM.Moved(pixel.V(200, 100)))

	for idx, tile := range screen.simpleArrows {
		tile.Draw(screen.systems.sprites, screen.screenBatch, pixel.IM.Moved(pixel.V(100+float64(50*idx), 200)))
	}

	for idx, tile := range screen.memoryBubbles {
		tile.Draw(screen.systems.sprites, screen.screenBatch, pixel.IM.Moved(pixel.V(100+float64(50*idx), 300)))
	}

	for idx, tile := range screen.tJunctions {
		tile.Draw(screen.systems.sprites, screen.screenBatch, pixel.IM.Moved(pixel.V(100+float64(50*idx), 400)))
	}

	screen.goalTile.Draw(screen.systems.sprites, screen.screenBatch, pixel.IM.Moved(pixel.V(100, 500)))
	screen.screenBatch.Draw(target)
}
