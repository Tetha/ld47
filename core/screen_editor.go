package core

import "github.com/faiface/pixel"

var _ (Screen) = (*EditLevelScreen)(nil)

type EditLevelScreen struct {
	systems     *Systems
	screenBatch *pixel.Batch
}

func NewEditLevelScreen(systems *Systems) *EditLevelScreen {
	return &EditLevelScreen{
		systems:     systems,
		screenBatch: pixel.NewBatch(&pixel.TrianglesData{}, systems.sprites.Sheet),
	}
}
func (screen *EditLevelScreen) Run(target pixel.Target, _ float64) {
}
