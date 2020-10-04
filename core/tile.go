package core

import "github.com/faiface/pixel"

type Tile interface {
	Draw(sprites *SpriteSystem, target pixel.Target, position pixel.Matrix)
	ModifyGhostPosition(position *GhostPosition)
}
