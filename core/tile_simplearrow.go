package core

import "github.com/faiface/pixel"

var _ Tile = (*TileSimpleArrow)(nil)

type ArrowDirection uint

const (
	ArrowDirectionStraightUp = iota
	ArrowDirectionStraightDown
	ArrowDirectionStraightRight
	ArrowDirectionStraightLeft

	ArrowDirectionLeftTurnFromBottom
	ArrowDirectionLeftTurnFromTop
	ArrowDirectionLeftTurnFromLeft
	ArrowDirectionLeftTurnFromRight

	ArrowDirectionRightTurnFromBottom
	ArrowDirectionRightTurnFromTop
	ArrowDirectionRightTurnFromLeft
	ArrowDirectionRightTurnFromRight
)

type TileSimpleArrow struct {
	direction ArrowDirection
	Marked    bool
}

func NewTileSimpleArrow(direction ArrowDirection) *TileSimpleArrow {
	return &TileSimpleArrow{direction: direction}
}

var directionToSprite = map[ArrowDirection]LargeTileID{
	ArrowDirectionStraightUp:    LargeTileArrowStraightUp,
	ArrowDirectionStraightDown:  LargeTileArrowStraightDown,
	ArrowDirectionStraightRight: LargeTileArrowStraightRight,
	ArrowDirectionStraightLeft:  LargeTileArrowStraightLeft,

	ArrowDirectionLeftTurnFromBottom: LargeTileArrowLeftTurnFromBottom,
	ArrowDirectionLeftTurnFromTop:    LargeTileArrowLeftTurnFromTop,
	ArrowDirectionLeftTurnFromLeft:   LargeTileArrowLeftTurnFromLeft,
	ArrowDirectionLeftTurnFromRight:  LargeTileArrowLeftTurnFromRight,

	ArrowDirectionRightTurnFromBottom: LargeTileArrowRightTurnFromBottom,
	ArrowDirectionRightTurnFromTop:    LargeTileArrowRightTurnFromTop,
	ArrowDirectionRightTurnFromLeft:   LargeTileArrowRightTurnFromLeft,
	ArrowDirectionRightTurnFromRight:  LargeTileArrowRightTurnFromRight,
}

func (tile *TileSimpleArrow) Draw(sprites *SpriteSystem, target pixel.Target, position pixel.Matrix) {
	sprite := sprites.tileSprites[directionToSprite[tile.direction]]
	sprite.Draw(target, position)

	if tile.Marked {
		marker := sprites.tileSprites[LargeTileMarker]
		marker.Draw(target, position)
	}
}

func (tile *TileSimpleArrow) Mark(marked bool) {
	tile.Marked = marked
}
func (tile *TileSimpleArrow) Reset() {
	tile.Marked = false
}

func (tile *TileSimpleArrow) ModifyGhostPosition(position *GhostPosition) {
	switch tile.direction {
	case ArrowDirectionStraightDown:
		position.direction = GhostDirectionDown
	case ArrowDirectionStraightUp:
		position.direction = GhostDirectionUp
	case ArrowDirectionStraightRight:
		position.direction = GhostDirectionRight
	case ArrowDirectionStraightLeft:
		position.direction = GhostDirectionLeft

	case ArrowDirectionLeftTurnFromBottom:
		if position.direction == GhostDirectionUp {
			position.direction = TurnLeft(position.direction)
		}
	case ArrowDirectionLeftTurnFromTop:
		if position.direction == GhostDirectionDown {
			position.direction = TurnLeft(position.direction)
		}
	case ArrowDirectionLeftTurnFromLeft:
		if position.direction == GhostDirectionRight {
			position.direction = TurnLeft(position.direction)
		}
	case ArrowDirectionLeftTurnFromRight:
		if position.direction == GhostDirectionLeft {
			position.direction = TurnLeft(position.direction)
		}

	case ArrowDirectionRightTurnFromBottom:
		if position.direction == GhostDirectionUp {
			position.direction = TurnRight(position.direction)
		}
	case ArrowDirectionRightTurnFromTop:
		if position.direction == GhostDirectionDown {
			position.direction = TurnRight(position.direction)
		}
	case ArrowDirectionRightTurnFromLeft:
		if position.direction == GhostDirectionRight {
			position.direction = TurnRight(position.direction)
		}
	case ArrowDirectionRightTurnFromRight:
		if position.direction == GhostDirectionLeft {
			position.direction = TurnRight(position.direction)
		}
	}
}
