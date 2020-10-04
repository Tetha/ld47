package core

var LevelFour = Level{
	Name: "Habits",
	Description: `And sometimes,
we spend some time
atray and need 
something to bring
us back.

Just dont fall
into old traps`,
	PresetTiles: []TileDefinition{
		{
			x:       2,
			y:       2,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromTop),
		},
		{
			x:       6,
			y:       2,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromLeft),
		},
		{
			x:       6,
			y:       6,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromBottom),
		},
		{
			x:       2,
			y:       6,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromRight),
		},

		{
			x:       8,
			y:       2,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromTop),
		},
		{
			x:       12,
			y:       2,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromLeft),
		},
		{
			x:       12,
			y:       6,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromBottom),
		},
		{
			x:       8,
			y:       6,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromRight),
		},

		{
			x:       8,
			y:       4,
			content: NewTileSimpleArrow(ArrowDirectionStraightDown),
		},

		{
			x:       6,
			y:       4,
			content: NewTileTJunction(GhostBlue, MemoryLove, TJunctionUp, TJunctionRightTurn, TJunctionMemoryOnTurn),
		},
		{
			x:       4,
			y:       6,
			content: NewTileTJunction(GhostBlue, MemoryLove, TJunctionLeft, TJunctionRightTurn, TJunctionMemoryOnTurn),
		},

		{
			x:       4,
			y:       2,
			content: NewTileMemory(GhostBlue, MemoryLove),
		},

		{
			x:       10,
			y:       2,
			content: NewTileMemory(GhostBlue, MemoryLove),
		},
		{
			x:       9,
			y:       2,
			content: NewTileMemory(GhostBlue, MemorySun),
		},

		{
			x:       4,
			y:       8,
			content: NewTileGoal(),
		},
	},
	InitialGhostPositions: []GhostPosition{
		{
			x:         2,
			y:         4,
			id:        GhostBlue,
			direction: GhostDirectionDown,
		},
	},
	Toolbox: []Tile{
		NewTileSimpleArrow(ArrowDirectionStraightUp),
		NewTileTJunction(GhostBlue, MemorySun, TJunctionDown, TJunctionRightTurn, TJunctionMemoryOnTurn),
	},
}
