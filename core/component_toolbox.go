package core

import "github.com/faiface/pixel"

type ToolboxComponent struct {
	systems    *Systems
	baseMatrix pixel.Matrix
	level      *Level
}

func NewToolboxComponent(systems *Systems, level *Level) *ToolboxComponent {
	return &ToolboxComponent{
		systems:    systems,
		baseMatrix: pixel.IM,
		level:      level,
	}
}

func (toolbox *ToolboxComponent) ResetBaseMatrix(baseMatrix pixel.Matrix) {
	toolbox.baseMatrix = baseMatrix
}

func (toolbox *ToolboxComponent) Draw(target pixel.Target) {

}
