package core

type Systems struct {
	sprites *SpriteSystem
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
