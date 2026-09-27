package startgame

import (
	"image/color"

	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/gui/component"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	guiadventure "github.com/KONshougun/K-game/gui/scene/guiAdventure"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StartGame struct {
	*component.Background
	InitLbl *component.Button
	Modes   []*component.Button
}

var chooseMode = false

func NewStartGameScene() *StartGame {

	area := component.Area{
		X:      0,
		Y:      0,
		Width:  gui.GameWidth,
		Height: gui.GameHeight,
	}
	s := &StartGame{
		Background: component.NewBackground(guihandler.LoadImage("asset/background/ui/Menu_Background.png"), area),
		InitLbl: component.NewButton([4]*ebiten.Image{}, component.NewLabel(
			"Buongiorno questa è una beta,\nse vi volete lamentare per la grafica andate da Paolo Schiantabasti,\nio mi occupo solo del backend",
			color.White, component.VeryLargeSize, component.VeryLargeSpace, text.AlignCenter, text.AlignCenter, area,
		), area, func() {
			chooseMode = true
		}),
		Modes: []*component.Button{component.NewButton([4]*ebiten.Image{
			guihandler.LoadImage("asset/ui/component/button/button-up.png"),
			guihandler.LoadImage("asset/ui/component/button/button-up.png"),
			guihandler.LoadImage("asset/ui/component/button/button-down.png"),
			guihandler.LoadImage("asset/ui/component/button/button-up.png"),
		}, component.NewLabel("Adventure", color.White, component.VeryLargeSize, component.NoneSpace,
			text.AlignCenter, text.AlignCenter, component.Area{
				X:      area.Width/2 - area.Width/8,
				Y:      area.Height/2 - area.Height/12,
				Width:  area.Width / 4,
				Height: area.Height / 6,
			}), component.Area{
			X:      area.Width/2 - area.Width/8,
			Y:      area.Height/2 - area.Height/12,
			Width:  area.Width / 4,
			Height: area.Height / 6,
		}, func() {
			gui.Game.Scenes = append(gui.Game.Scenes, guiadventure.StartAdventureScene())
		})},
	}
	return s
}

func (s *StartGame) Update() error {

	if !chooseMode {
		s.InitLbl.Update()
	} else {
		for _, b := range s.Modes {
			b.Update()
		}
	}
	return nil
}
func (s *StartGame) Draw(screen *ebiten.Image) {
	s.Background.Draw(screen)

	if !chooseMode {
		s.InitLbl.Draw(screen)
	} else {
		for _, b := range s.Modes {
			b.Draw(screen)
		}
	}
}

func (s *StartGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gui.GameWidth, gui.GameHeight
}
