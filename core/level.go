package core

type TileDefinition struct {
	x       uint
	y       uint
	content Tile
}

type Level struct {
	InitialToolbox        []Tile
	PresetTiles           []TileDefinition
	InitialGhostPositions []GhostPosition
}
