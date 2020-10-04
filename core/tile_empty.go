package core

import "github.com/faiface/pixel"

var _ Tile = (*TileEmpty)(nil)

type TileEmpty struct {
	Highlight bool
}

func (tile *TileEmpty) Draw(sprites *SpriteSystem, target pixel.Target, position pixel.Matrix) {
	if !tile.Highlight {
		return
	}
	markerSprite := sprites.tileSprites[LargeTileMarker]
	markerSprite.DrawColorMask(target, position, pixel.RGB(0, 0.8, 0))
}

func (tile *TileEmpty) Mark(marked bool) {
}

func (tile *TileEmpty) ModifyGhostPosition(position *GhostPosition) {
}
