package core

type GhostDirection uint

const (
	GhostDirectionUp GhostDirection = iota
	GhostDirectionDown
	GhostDirectionLeft
	GhostDirectionRight
)

type GhostPosition struct {
	x         uint
	y         uint
	id        GhostID
	direction GhostDirection
	inventory []MemoryType
}

func (p GhostPosition) HasMemory(requiredMemType MemoryType) bool {
	for _, memType := range p.inventory {
		if memType == requiredMemType {
			return true
		}
	}
	return false
}

func (p *GhostPosition) RemoveFirstMemory(target MemoryType) {
	for idx, memType := range p.inventory {
		if memType == target {
			p.inventory = append(p.inventory[:idx], p.inventory[idx+1:]...)
			return
		}
	}
}

func (p GhostPosition) Clone() GhostPosition {
	return GhostPosition{
		x:         p.x,
		y:         p.y,
		direction: p.direction,
		id:        p.id,
		inventory: append([]MemoryType(nil), p.inventory...),
	}
}
func TurnLeft(direction GhostDirection) GhostDirection {
	switch direction {
	case GhostDirectionDown:
		return GhostDirectionRight
	case GhostDirectionUp:
		return GhostDirectionLeft
	case GhostDirectionLeft:
		return GhostDirectionDown
	case GhostDirectionRight:
		return GhostDirectionUp
	}
	panic("Impossible")
}

func TurnRight(direction GhostDirection) GhostDirection {
	switch direction {
	case GhostDirectionDown:
		return GhostDirectionLeft
	case GhostDirectionUp:
		return GhostDirectionRight
	case GhostDirectionLeft:
		return GhostDirectionUp
	case GhostDirectionRight:
		return GhostDirectionDown
	}
	panic("Impossible")
}
