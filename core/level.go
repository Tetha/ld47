package core

type TileDefinition struct {
	x       uint
	y       uint
	content Tile
}

type Level struct {
	Toolbox               []Tile
	PresetTiles           []TileDefinition
	InitialGhostPositions []GhostPosition
}

type PlayerInput struct {
	RemainingTools []Tile
	PlacedTools    []TileDefinition

	SelectedTile Tile
}
