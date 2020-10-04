package core

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
)

var _ (Screen) = (*LevelSelectScreen)(nil)

type LevelSelectScreen struct {
	systems *Systems

	selectedLevelID int
}

var levels []Level = []Level{
	TestLevelOne,
	LevelTwo,
	LevelThree,
	LevelFour,
	LevelFive,
}

const maxY float64 = 700

func NewScreenLevelSelect(systems *Systems) Screen {
	fmt.Fprint(systems.MainScreen.UpperButtonText, ">> Play Level! >>")
	fmt.Fprint(systems.MainScreen.LowerButtonText, "Quit")

	return &LevelSelectScreen{systems: systems}
}

func (screen *LevelSelectScreen) Run(target pixel.Target, dt float64) GameState {
	draw := imdraw.New(nil)

	for idx, level := range levels {
		top := maxY - float64(idx*50)
		draw.Color = colornames.White
		draw.Push(pixel.V(50, top))
		draw.Push(pixel.V(600, top))
		draw.Push(pixel.V(600, top-50))
		draw.Push(pixel.V(50, top-50))
		draw.Push(pixel.V(50, top))
		draw.Line(5)

		if idx == screen.selectedLevelID {
			draw.Color = colornames.Blue
			draw.Push(pixel.V(60, top-25))
			draw.Circle(5, 4)
		}
		draw.Draw(target)

		text := text.New(pixel.V(75, maxY-float64(idx*50)-20), screen.systems.MainScreen.Atlas)
		if level.Won {
			text.Color = colornames.Green
		} else {
			text.Color = colornames.White
		}
		fmt.Fprint(text, level.Name)
		text.Draw(target, pixel.IM)
	}
	return GameStateKeep
}

func (screen *LevelSelectScreen) Click(pos pixel.Vec) GameState {
	if screen.systems.MainScreen.UpperButtonBounds.Contains(pos) {
		screen.systems.SetLevel(&levels[screen.selectedLevelID])
		return GameStateEdit
	}
	if screen.systems.MainScreen.LowerButtonBounds.Contains(pos) {
		return GameStateQuit
	}

	if 50 <= pos.X && pos.X <= 600 {
		levelId := -(int(math.Floor((pos.Y-maxY)/50)) + 1)
		if 0 <= levelId && levelId < len(levels) {
			screen.selectedLevelID = levelId
		}
	}
	return GameStateKeep
}
