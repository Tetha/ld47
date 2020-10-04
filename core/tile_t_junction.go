package core

import (
	"math"

	"github.com/faiface/pixel"
)

var _ Tile = (*TileTJunction)(nil)

type TJunctionDirection uint

const (
	TJunctionLeftTurn         = true
	TJunctionRightTurn        = false
	TJunctionMemoryOnStraight = true
	TJunctionMemoryOnTurn     = false
)
const (
	TJunctionUp = iota
	TJunctionDown
	TJunctionLeft
	TJunctionRight
)

type TileTJunction struct {
	Direction        TJunctionDirection
	JunctionLeft     bool
	MemoryOnStraight bool

	RequiredMemory MemoryType
	Ghost          GhostID
}

func NewTileTJunction(ghost GhostID, requiredMemory MemoryType, direction TJunctionDirection, junctionLeft bool, memoryOnStraight bool) *TileTJunction {
	return &TileTJunction{
		Direction:        direction,
		JunctionLeft:     junctionLeft,
		MemoryOnStraight: memoryOnStraight,

		RequiredMemory: requiredMemory,
		Ghost:          ghost,
	}
}

func (tile *TileTJunction) Draw(sprites *SpriteSystem, target pixel.Target, position pixel.Matrix) {
	baseTile := sprites.tileSprites[LargeTileArrowT]

	baseTransformation := pixel.IM

	if tile.JunctionLeft {
		baseTransformation = baseTransformation.ScaledXY(pixel.ZV, pixel.V(-1, 1))
	}
	if tile.Direction == TJunctionDown {
		baseTransformation = baseTransformation.Scaled(pixel.ZV, -1)
	}
	if tile.Direction == TJunctionLeft {
		baseTransformation = baseTransformation.Rotated(pixel.ZV, -math.Pi/2)
	}
	if tile.Direction == TJunctionRight {
		baseTransformation = baseTransformation.Rotated(pixel.ZV, math.Pi/2)
	}

	baseTransformation = baseTransformation.Chained(position)
	baseTile.Draw(target, baseTransformation)

	memoryTile := sprites.tileSprites[memoryTypeToTile[tile.RequiredMemory]]
	noneTile := sprites.tileSprites[LargeTileMemoryNone]

	memoryOffset := pixel.ZV
	noneOffset := pixel.ZV

	if tile.MemoryOnStraight {
		memoryOffset = memoryOffset.Add(pixel.V(-11, 0))
		noneOffset = noneOffset.Add(pixel.V(15, -15))
	} else {
		memoryOffset = memoryOffset.Add(pixel.V(15, -15))
		noneOffset = noneOffset.Add(pixel.V(-11, 0))
	}
	memoryTile.DrawColorMask(target, pixel.IM.Scaled(pixel.ZV, 0.25).Moved(memoryOffset).Chained(baseTransformation), ghostToColorMask[tile.Ghost])
	noneTile.Draw(target, pixel.IM.Scaled(pixel.ZV, 0.25).Moved(noneOffset).Chained(baseTransformation))

}

func (tile *TileTJunction) ModifyGhostPosition(position *GhostPosition) {
	switch tile.Direction {
	case TJunctionDown:
		if position.direction != GhostDirectionDown {
			return
		}

		if position.HasMemory(tile.RequiredMemory) {
			position.RemoveFirstMemory(tile.RequiredMemory)
			// Memory is on the turn
			if !tile.MemoryOnStraight {
				if tile.JunctionLeft {
					position.direction = TurnLeft(position.direction)
				} else {
					position.direction = TurnRight(position.direction)
				}
			}
			return
		} else {
			// Memory is required straight ahead but we dont have it
			if tile.MemoryOnStraight {
				if tile.JunctionLeft {
					position.direction = TurnLeft(position.direction)
				} else {
					position.direction = TurnRight(position.direction)
				}
			}
		}

	case TJunctionUp:
		if position.direction != GhostDirectionUp {
			return
		}

		if position.HasMemory(tile.RequiredMemory) {
			position.RemoveFirstMemory(tile.RequiredMemory)
			// Memory is on the turn
			if !tile.MemoryOnStraight {
				if tile.JunctionLeft {
					position.direction = TurnLeft(position.direction)
				} else {
					position.direction = TurnRight(position.direction)
				}
			}
			return
		} else {
			// Memory is required straight ahead but we dont have it
			if tile.MemoryOnStraight {
				if tile.JunctionLeft {
					position.direction = TurnLeft(position.direction)
				} else {
					position.direction = TurnRight(position.direction)
				}
			}
		}

	case TJunctionLeft:
		if position.direction != GhostDirectionLeft {
			return
		}

		if position.HasMemory(tile.RequiredMemory) {
			position.RemoveFirstMemory(tile.RequiredMemory)
			// Memory is on the turn
			if !tile.MemoryOnStraight {
				if tile.JunctionLeft {
					position.direction = TurnLeft(position.direction)
				} else {
					position.direction = TurnRight(position.direction)
				}
			}
			return
		} else {
			// Memory is required straight ahead but we dont have it
			if tile.MemoryOnStraight {
				if tile.JunctionLeft {
					position.direction = TurnLeft(position.direction)
				} else {
					position.direction = TurnRight(position.direction)
				}
			}
		}

	case TJunctionRight:
		if position.direction != GhostDirectionRight {
			return
		}

		if position.HasMemory(tile.RequiredMemory) {
			position.RemoveFirstMemory(tile.RequiredMemory)
			// Memory is on the turn
			if !tile.MemoryOnStraight {
				if tile.JunctionLeft {
					position.direction = TurnLeft(position.direction)
				} else {
					position.direction = TurnRight(position.direction)
				}
			}
			return
		} else {
			// Memory is required straight ahead but we dont have it
			if tile.MemoryOnStraight {
				if tile.JunctionLeft {
					position.direction = TurnLeft(position.direction)
				} else {
					position.direction = TurnRight(position.direction)
				}
			}
		}
	}
}
