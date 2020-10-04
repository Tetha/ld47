package core

var TestLevelOne = Level{
	Name: "Tutorial",
	Description: `Look at these
ghosts. They are
stuck in limbo.

Use these arrows 
and guide them to
the goal (that
circle).

Click on them in
the toolbox, then
the grid to place
them. To remove
them. click again.

Then click Run to
see what happens`,
	PresetTiles: []TileDefinition{
		{
			x:       4,
			y:       4,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromTop),
		},
		{
			x:       4,
			y:       12,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromRight),
		},
		{
			x:       12,
			y:       4,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromLeft),
		},
		{
			x:       12,
			y:       12,
			content: NewTileSimpleArrow(ArrowDirectionLeftTurnFromBottom),
		},
		{
			x:       6,
			y:       6,
			content: NewTileSimpleArrow(ArrowDirectionRightTurnFromRight),
		},
		{
			x:       6,
			y:       10,
			content: NewTileSimpleArrow(ArrowDirectionRightTurnFromBottom),
		},
		{
			x:       10,
			y:       6,
			content: NewTileSimpleArrow(ArrowDirectionRightTurnFromTop),
		},
		{
			x:       10,
			y:       10,
			content: NewTileSimpleArrow(ArrowDirectionRightTurnFromLeft),
		},
		{
			x:       13,
			y:       7,
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

		{
			x:         9,
			y:         6,
			id:        GhostOrange,
			direction: GhostDirectionLeft,
		},
	},
	Toolbox: []Tile{
		NewTileSimpleArrow(ArrowDirectionStraightRight),
		NewTileSimpleArrow(ArrowDirectionStraightRight),
	},
}
