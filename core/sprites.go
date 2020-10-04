package core

import (
	"image"
	_ "image/png" // to load PNG files

	"os"

	"github.com/faiface/pixel"
)

type LargeTileID uint

const (
	LargeTileGhost LargeTileID = iota
	LargeTileArrowStraightUp
	LargeTileArrowStraightRight
	LargeTileArrowStraightDown
	LargeTileArrowStraightLeft

	LargeTileArrowLeftTurnFromBottom
	LargeTileArrowLeftTurnFromLeft
	LargeTileArrowLeftTurnFromTop
	LargeTileArrowLeftTurnFromRight

	LargeTileArrowRightTurnFromBottom
	LargeTileArrowRightTurnFromLeft
	LargeTileArrowRightTurnFromTop
	LargeTileArrowRightTurnFromRight

	LargeTileArrowT

	LargeTileTarget

	LargeTileMemoryBubble

	LargeTileMemoryLove
	LargeTileMemoryPet
	LargeTileMemorySun
	LargeTileMemoryNone

	LargeTileMarker
)

var largeTiles = []LargeTileID{
	LargeTileGhost,
	LargeTileArrowStraightUp,
	LargeTileArrowStraightRight,
	LargeTileArrowStraightDown,
	LargeTileArrowStraightLeft,

	LargeTileArrowLeftTurnFromBottom,
	LargeTileArrowLeftTurnFromLeft,
	LargeTileArrowLeftTurnFromTop,
	LargeTileArrowLeftTurnFromRight,

	LargeTileArrowRightTurnFromBottom,
	LargeTileArrowRightTurnFromLeft,
	LargeTileArrowRightTurnFromTop,
	LargeTileArrowRightTurnFromRight,

	LargeTileArrowT,

	LargeTileTarget,

	LargeTileMemoryBubble,

	LargeTileMemoryLove,
	LargeTileMemoryPet,
	LargeTileMemorySun,
	LargeTileMemoryNone,

	LargeTileMarker,
}

type MaskID uint

const (
	MaskGeneric MaskID = iota
)

var masks = []MaskID{
	MaskGeneric,
}

type SpriteSystem struct {
	Sheet     *pixel.PictureData
	MaskSheet *pixel.PictureData

	tileSprites map[LargeTileID]*pixel.Sprite
	MaskSprites map[MaskID]*pixel.Sprite
}

const assetPath = "assets/tiles.png"

func LoadSprites() (*SpriteSystem, error) {
	file, err := os.Open(assetPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	result := new(SpriteSystem)
	result.Sheet = pixel.PictureDataFromImage(img)
	result.tileSprites = make(map[LargeTileID]*pixel.Sprite)

	for _, id := range largeTiles {
		bounds := pixel.R(0, result.Sheet.Rect.H()-float64(48*id), 48, result.Sheet.Rect.H()-float64(48.0*(id+1)))
		result.tileSprites[id] = pixel.NewSprite(result.Sheet, bounds)
	}

	if err := LoadMasks(result); err != nil {
		return nil, err
	}
	return result, nil
}

func LoadMasks(sprites *SpriteSystem) error {
	file, err := os.Open("assets/overall_masks.png")
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	sprites.MaskSheet = pixel.PictureDataFromImage(img)
	sprites.MaskSprites = make(map[MaskID]*pixel.Sprite, len(masks))

	for _, id := range masks {
		bounds := pixel.R(0, sprites.MaskSheet.Rect.H()-float64(768*id), 1024, sprites.MaskSheet.Rect.H()-float64(768*(id+1)))
		sprites.MaskSprites[id] = pixel.NewSprite(sprites.MaskSheet, bounds)
	}
	return nil
}
