package core

import (
	"math"

	"github.com/faiface/pixel"
)

var _ Tile = (*TileGhost)(nil)

type GhostID uint

const (
	GhostBlue = iota
	GhostOrange
)

var ghostToColorMask = map[GhostID]pixel.RGBA{
	GhostBlue:   pixel.RGB(0.2, 0.2, 0.9),
	GhostOrange: pixel.RGB(0.8, 0.4, 0),
}

type TileGhost struct {
	ghostID   GhostID
	inventory []MemoryType
}

func NewTileGhost(ghostID GhostID) *TileGhost {
	return &TileGhost{ghostID: ghostID}
}

func (tile *TileGhost) SetInventory(newInventory []MemoryType) {
	tile.inventory = newInventory
}

var specialSlots = []float64{math.Pi / 2, math.Pi / 4, math.Pi * 3 / 4}

func (tile *TileGhost) Draw(sprites *SpriteSystem, target pixel.Target, position pixel.Matrix) {
	ghostSprite := sprites.tileSprites[LargeTileGhost]
	ghostSprite.DrawColorMask(target, position, ghostToColorMask[tile.ghostID])

	for idx, invMemory := range tile.inventory {
		rotation := math.Pi + math.Pi/4.0*float64(idx-len(specialSlots))
		if idx < len(specialSlots) {
			rotation = specialSlots[idx]
		}
		memorySprite := sprites.tileSprites[memoryTypeToTile[invMemory]]
		memorySprite.DrawColorMask(
			target,
			pixel.IM.Scaled(pixel.ZV, 0.2).Rotated(pixel.ZV, -rotation).Moved(pixel.V(30, 0)).Rotated(pixel.ZV, rotation).Chained(position),
			ghostToColorMask[tile.ghostID])

	}
}

func (tile *TileGhost) ModifyGhostPosition(position *GhostPosition) {
}
