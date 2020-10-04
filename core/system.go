package core

type Systems struct {
	sprites *SpriteSystem

	level      *Level
	simulation *SimulationState
}

func InitSystems() (*Systems, error) {
	result := new(Systems)

	sprites, err := LoadSprites()
	if err != nil {
		return nil, err
	}
	result.sprites = sprites

	return result, nil
}

func (systems *Systems) SetLevel(level *Level) {
	systems.level = level

	// reset the simulation state on level change as well
	systems.simulation = &SimulationState{}
}
