package core

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type Systems struct {
	Sprites *SpriteSystem

	MainScreen MainButtons

	level      *Level
	simulation *SimulationState
	input      *PlayerInput
}

func InitSystems() (*Systems, error) {
	result := new(Systems)

	sprites, err := LoadSprites()
	if err != nil {
		return nil, err
	}
	result.Sprites = sprites

	result.MainScreen.Atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII)
	result.MainScreen.UpperButtonBounds = pixel.R(860, 45, 1000, 70)
	result.MainScreen.LowerButtonBounds = pixel.R(860, 10, 1000, 40)
	result.MainScreen.UpperButtonText = text.New(
		pixel.V(result.MainScreen.UpperButtonBounds.Min.X+10, result.MainScreen.UpperButtonBounds.Min.Y+10),
		result.MainScreen.Atlas)
	result.MainScreen.LowerButtonText = text.New(
		pixel.V(result.MainScreen.LowerButtonBounds.Min.X+10, result.MainScreen.LowerButtonBounds.Min.Y+10),
		result.MainScreen.Atlas)

	result.MainScreen.DescriptionText = text.New(
		pixel.V(870, 340),
		result.MainScreen.Atlas)
	result.MainScreen.DescriptionText.Color = colornames.Black
	return result, nil
}

func (systems *Systems) SetLevel(level *Level) {
	systems.level = level

	// reset the simulation state on level change as well
	systems.simulation = &SimulationState{}
	systems.input = &PlayerInput{
		RemainingTools: append(([]Tile)(nil), level.Toolbox...),
	}
}

type MainButtons struct {
	Atlas *text.Atlas

	UpperButtonText *text.Text
	LowerButtonText *text.Text

	UpperButtonBounds pixel.Rect
	LowerButtonBounds pixel.Rect

	DescriptionText *text.Text
}
