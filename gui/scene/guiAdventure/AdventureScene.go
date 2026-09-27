package guiadventure

import (
	"github.com/KONshougun/K-game/game/adventure"
	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/gui/component"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"

	"github.com/hajimehoshi/ebiten/v2"
)

type AdvenureManager struct {
	*component.Background

	AdventureData *adventure.Adventure
}

func NewAdvenureScene(adventureData *adventure.Adventure) *AdvenureManager {
	area := component.Area{
		X:      0,
		Y:      0,
		Width:  gui.GameWidth,
		Height: gui.GameHeight,
	}
	s := &AdvenureManager{
		Background:    component.NewBackground(guihandler.LoadImage("asset/background/ui/Menu_Background.png"), area),
		AdventureData: adventureData,
	}
	return s
}

func (s *AdvenureManager) Update() error {
	return nil
}
func (s *AdvenureManager) Draw(screen *ebiten.Image) {
	s.Background.Draw(screen)
}

func (s *AdvenureManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gui.GameWidth, gui.GameHeight
}
