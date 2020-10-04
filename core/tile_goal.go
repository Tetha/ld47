package core

import "github.com/faiface/pixel"

var _ Tile = (*TileGoal)(nil)

type TileGoal struct {
}

func NewTileGoal() *TileGoal {
	return &TileGoal{}
}

func (tile *TileGoal) Draw(sprites *SpriteSystem, target pixel.Target, position pixel.Matrix) {
	goalSprite := sprites.tileSprites[LargeTileTarget]
	goalSprite.Draw(target, position)
}

func (tile *TileGoal) ModifyGhostPosition(position *GhostPosition) {
}
