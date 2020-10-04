package core

import (
	"math"

	"github.com/faiface/pixel"
)

type ToolboxComponent struct {
	systems    *Systems
	baseMatrix pixel.Matrix
}

func NewToolboxComponent(systems *Systems) *ToolboxComponent {
	return &ToolboxComponent{
		systems:    systems,
		baseMatrix: pixel.IM,
	}
}

func (toolbox *ToolboxComponent) ResetBaseMatrix(baseMatrix pixel.Matrix) {
	toolbox.baseMatrix = baseMatrix
}

func (toolbox *ToolboxComponent) Draw(target pixel.Target) {
	for idx, tile := range toolbox.systems.input.RemainingTools {
		tile.Draw(toolbox.systems.Sprites, target, toolbox.baseMatrix.Moved(pixel.V(25, 25+60*float64(idx))))
	}
}

func (toolbox *ToolboxComponent) SelectItem(position pixel.Vec) {
	toolboxPosition := toolbox.baseMatrix.Unproject(position)

	tileID := int(math.Floor((toolboxPosition.Y) / 60))
	//fmt.Printf("Click at: %+v, tileID is %d\n", toolboxPosition, tileID)
	if 0 <= tileID && tileID < len(toolbox.systems.input.RemainingTools) {
		if toolbox.systems.input.SelectedTile == toolbox.systems.input.RemainingTools[tileID] {
			toolbox.systems.input.SelectedTile.Mark(false)
			toolbox.systems.input.SelectedTile = nil
		} else {
			if toolbox.systems.input.SelectedTile != nil {
				toolbox.systems.input.SelectedTile.Mark(false)
			}
			toolbox.systems.input.SelectedTile = toolbox.systems.input.RemainingTools[tileID]
			toolbox.systems.input.SelectedTile.Mark(true)
		}

	} else {
		if toolbox.systems.input.SelectedTile != nil {
			toolbox.systems.input.SelectedTile.Mark(false)
			toolbox.systems.input.SelectedTile = nil
		}
	}
}
