package core

var TestLevelOne = Level{
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
}
