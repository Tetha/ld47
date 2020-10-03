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
}
