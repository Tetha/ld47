package core

import "github.com/faiface/pixel"

type MemoryType uint

var _ Tile = (*TileMemory)(nil)

const (
	MemoryLove = iota
	MemoryPet  = iota
	MemorySun  = iota
)

type TileMemory struct {
	memoryType MemoryType
	ghostID    GhostID
	collected  bool
}

func NewTileMemory(ghostID GhostID, memoryType MemoryType) *TileMemory {
	return &TileMemory{memoryType: memoryType, ghostID: ghostID}
}

var memoryTypeToTile = map[MemoryType]LargeTileID{
	MemoryLove: LargeTileMemoryLove,
	MemoryPet:  LargeTileMemoryPet,
	MemorySun:  LargeTileMemorySun,
}

func (tile *TileMemory) Draw(sprites *SpriteSystem, target pixel.Target, position pixel.Matrix) {
	memoryBubble := sprites.tileSprites[LargeTileMemoryBubble]
	memoryBubble.Draw(target, position)

	memoryIcon := sprites.tileSprites[memoryTypeToTile[tile.memoryType]]
	if !tile.collected {
		memoryIcon.DrawColorMask(target, pixel.IM.Scaled(pixel.ZV, 0.25).Chained(position).Moved(pixel.V(0, 5)), ghostToColorMask[tile.ghostID])
	} else {
		memoryIcon.DrawColorMask(target, pixel.IM.Scaled(pixel.ZV, 0.25).Chained(position).Moved(pixel.V(0, 5)), pixel.Alpha(0.7))
	}
}

func (tile *TileMemory) Mark(marked bool) {
}

func (tile *TileMemory) ModifyGhostPosition(position *GhostPosition) {
	if !tile.collected {
		position.inventory = append(position.inventory, tile.memoryType)
		tile.collected = true
	}
}
