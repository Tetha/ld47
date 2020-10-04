package core

var LevelTwo = Level{
	Name: "Memories",
	Description: `Sometimes a ghost
is stuck in a loop
and needs to remember
somthing happy.

Use the memory from
the toolbox to help`,
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
			x:       7,
			y:       3,
			content: NewTileGoal(),
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
	Toolbox: []Tile{
		NewTileMemory(GhostBlue, MemoryLove),
	},
}
