package core

var TestLevelTwo = Level{
	PresetTiles: []TileDefinition{
		{
			x:       2,
			y:       2,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromTop),
		},
		{
			x:       2,
			y:       4,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromRight),
		},
		{
			x:       4,
			y:       2,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromLeft),
		},
		{
			x:       3,
			y:       2,
			content: NewTileMemory(GhostBlue, MemoryLove),
		},
		{
			x:       4,
			y:       4,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromBottom),
		},
		{
			x:       2,
			y:       3,
			content: NewTileTJunction(GhostBlue, MemoryLove, TJunctionDown, TJunctionLeftTurn, TJunctionMemoryOnTurn),
		},
		{
			x:       6,
			y:       2,
			content: NewTileSimpleArrow(ArrowDirectionRightTurnFromRight),
		},
		{
			x:       6,
			y:       3,
			content: NewTileSimpleArrow(ArrowDirectionStraightUp),
		},
		{
			x:       6,
			y:       4,
			content: NewTileSimpleArrow(ArrowDirectionRightTurnFromBottom),
		},
		{
			x:       8,
			y:       2,
			content: NewTileSimpleArrow(ArrowDirectionRightTurnFromTop),
		},
		{
			x:       7,
			y:       2,
			content: NewTileMemory(GhostBlue, MemoryLove),
		},
		{
			x:       8,
			y:       4,
			content: NewTileSimpleArrow(ArrowDirectionRightTurnFromLeft),
		},
	},
	InitialGhostPositions: []GhostPosition{
		{
			x:         3,
			y:         4,
			id:        GhostBlue,
			direction: GhostDirectionLeft,
		},
	},
}
