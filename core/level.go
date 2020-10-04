package core

type TileDefinition struct {
	x       uint
	y       uint
	content Tile
}

type Level struct {
	Name                  string
	Description           string
	Won                   bool
	Toolbox               []Tile
	PresetTiles           []TileDefinition
	InitialGhostPositions []GhostPosition
}

type PlayerInput struct {
	RemainingTools []Tile
	PlacedTools    []TileDefinition

	SelectedTile Tile
}
