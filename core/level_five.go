package core

var LevelFive = Level{
	Name: "New Experiences",
	Description: `Somtimes, it
helps to see some
new experiences`,
	PresetTiles: []TileDefinition{
		// Heart Hook Left
		{
			x:       3,
			y:       7,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromRight),
		},
		{
			x:       3,
			y:       5,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromTop),
		},
		{
			x:       4,
			y:       7,
			content: NewTileMemory(GhostBlue, MemoryLove),
		},
		// Heart Hook right
		{
			x:       14,
			y:       7,
			content: NewTileSimpleArrow(ArrowDirectionRightTurnFromLeft),
		},
		{
			x:       14,
			y:       5,
			content: NewTileSimpleArrow(ArrowDirectionRightTurnFromTop),
		},
		{
			x:       13,
			y:       7,
			content: NewTileMemory(GhostBlue, MemoryLove),
		},
		// Central loop
		{
			x:       6,
			y:       9,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromRight),
		},
		{
			x:       6,
			y:       7,
			content: NewTileTJunction(GhostBlue, MemorySun, TJunctionDown, TJunctionRightTurn, TJunctionMemoryOnTurn),
		},
		{
			x:       6,
			y:       4,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromTop),
		},
		{
			x:       10,
			y:       4,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromLeft),
		},
		{
			x:       10,
			y:       7,
			content: NewTileTJunction(GhostBlue, MemoryPet, TJunctionUp, TJunctionRightTurn, TJunctionMemoryOnTurn),
		},
		{
			x:       10,
			y:       9,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromBottom),
		},
		//Top exit
		{
			x:       9,
			y:       9,
			content: NewTileSimpleArrow(ArrowDirectionStraightLeft),
		},
		{
			x:       7,
			y:       9,
			content: NewTileTJunction(GhostBlue, MemoryLove, TJunctionLeft, TJunctionRightTurn, TJunctionMemoryOnTurn),
		},
		{
			x:       7,
			y:       10,
			content: NewTileTJunction(GhostBlue, MemoryLove, TJunctionUp, TJunctionRightTurn, TJunctionMemoryOnStraight),
		},
		{
			x:       9,
			y:       10,
			content: NewTileSimpleArrow(ArrowDirectionRightTurnFromLeft),
		},
		{
			x:       7,
			y:       12,
			content: NewTileGoal(),
		},
	},
	InitialGhostPositions: []GhostPosition{
		{
			x:         8,
			y:         4,
			id:        GhostBlue,
			direction: GhostDirectionRight,
		},
	},
	Toolbox: []Tile{
		NewTileSimpleArrow(ArrowDirectionStraightUp),
		NewTileSimpleArrow(ArrowDirectionStraightDown),
		NewTileMemory(GhostBlue, MemoryPet),
		NewTileMemory(GhostBlue, MemorySun),
	},
}
