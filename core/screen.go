package core

import "github.com/faiface/pixel"

type Screen interface {
	Run(target pixel.Target, dt float64)
}
