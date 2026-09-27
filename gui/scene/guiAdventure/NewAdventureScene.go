package guiadventure

import (
	"image/color"

	"github.com/KONshougun/K-game/game/adventure"
	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/gui/component"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StartAdventure struct {
	*component.Background
	InitLbl *component.Button
}

var lbls = []string{
	"Benvenuto avventuriero\nLe aspetta un'avventura piena di sorprese e incontri inaspettati",
	"La sua avventura inizierà nei panni di un guerriero di nome Knight",
	"Durante la tua avventura affronterai diverse avversità\ndiventa sempre più forte e scopri le leggi che governano questo mondo",
}
var i = 0

func StartAdventureScene() *StartAdventure {
	area := component.Area{
		X:      0,
		Y:      0,
		Width:  gui.GameWidth,
		Height: gui.GameHeight,
	}
	s := &StartAdventure{
		Background: component.NewBackground(guihandler.LoadImage("asset/background/ui/Menu_Background.png"), area),
		InitLbl: component.NewButton([4]*ebiten.Image{}, component.NewLabel(lbls[i], color.White,
			component.VeryLargeSize, component.VeryLargeSpace, text.AlignCenter, text.AlignCenter, area,
		), area, func() {
			i++
			if i == len(lbls) {
				gui.Game.Scenes = append(gui.Game.Scenes, NewAdvenureScene(adventure.InitAdventure()))
			}
		}),
	}
	return s
}

func (s *StartAdventure) Update() error {
	s.InitLbl.Update()
	return nil
}
func (s *StartAdventure) Draw(screen *ebiten.Image) {
	if i < len(lbls) {
		s.InitLbl.Label.Text = lbls[i]
	}
	s.Background.Draw(screen)
	s.InitLbl.Draw(screen)
}

func (s *StartAdventure) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gui.GameWidth, gui.GameHeight
}
