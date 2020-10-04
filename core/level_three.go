package core

var LevelThree = Level{
	Name: "Astray",
	Description: `Sometimes these
memories lead us
astray. Help
the ghost avoid
this.`,
	PresetTiles: []TileDefinition{
		{
			x:       2,
			y:       5,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromTop),
		},
		{
			x:       2,
			y:       6,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromRight),
		},
		{
			x:       3,
			y:       5,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromLeft),
		},
		{
			x:       3,
			y:       6,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromBottom),
		},

		{
			x:       6,
			y:       6,
			content: NewTileTJunction(GhostBlue, MemoryLove, TJunctionUp, TJunctionLeftTurn, TJunctionMemoryOnTurn),
		},
		{
			x:       6,
			y:       3,
			content: NewTileMemory(GhostBlue, MemoryLove),
		},
		{
			x:       6,
			y:       9,
			content: NewTileGoal(),
		},
	},
	InitialGhostPositions: []GhostPosition{
		{
			x:         6,
			y:         2,
			id:        GhostBlue,
			direction: GhostDirectionUp,
		},
	},
	Toolbox: []Tile{
		NewTileSimpleArrow(ArrowDirectionRightTurnFromBottom),
		NewTileSimpleArrow(ArrowDirectionLeftTurnFromLeft),
		NewTileSimpleArrow(ArrowDirectionLeftTurnFromBottom),
		NewTileSimpleArrow(ArrowDirectionRightTurnFromRight),
	},
}
