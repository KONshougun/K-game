package startgame

import (
	"fmt"
	"image/color"

	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/gui/component"
	"github.com/KONshougun/K-game/gui/component/container"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	guiadventure "github.com/KONshougun/K-game/gui/scene/guiAdventure"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StartGame struct {
	*container.SceneData
}

var chooseMode = false

func NewStartGameScene() *StartGame {
	s := &StartGame{
		SceneData: container.NewScene(guihandler.LoadImage("asset/background/ui/Menu_Background.png")),
	}

	initBtn := component.NewButton([4]*ebiten.Image{}, component.NewLabel(
		"Buongiorno questa è una beta,\nse vi volete lamentare per la grafica andate da Paolo Schiantabasti,\nio mi occupo solo del backend",
		color.White, component.VeryLargeSize, component.VeryLargeSpace, text.AlignCenter, text.AlignCenter, s.Area,
	), s.Area, nil)
	initBtn.OnClick = func() {
		chooseMode = true
		s.RemoveChild(initBtn.Id)

		//Creo i bottoni per le modalità
		adventureBtn := component.NewButton([4]*ebiten.Image{
			guihandler.LoadImage("asset/ui/component/button/button-up.png"),
			guihandler.LoadImage("asset/ui/component/button/button-up.png"),
			guihandler.LoadImage("asset/ui/component/button/button-down.png"),
			guihandler.LoadImage("asset/ui/component/button/button-up.png"),
		}, component.NewLabel("Adventure", color.White, component.VeryLargeSize, component.NoneSpace,
			text.AlignCenter, text.AlignCenter, component.Area{
				X:      s.Area.Width/2 - s.Area.Width/8,
				Y:      s.Area.Height/2 - s.Area.Height/12,
				Width:  s.Area.Width / 4,
				Height: s.Area.Height / 6,
			}), component.Area{
			X:      s.Area.Width/2 - s.Area.Width/8,
			Y:      s.Area.Height/2 - s.Area.Height/12,
			Width:  s.Area.Width / 4,
			Height: s.Area.Height / 6,
		}, func() {
			gui.Game.AddScene(guiadventure.StartAdventureScene())
		})
		if !s.AddChild(container.Child{C: adventureBtn}) {
			fmt.Println("wtf")
		}
	}
	if !s.AddChild(container.Child{C: initBtn}) {
		fmt.Println("Daje")
	}
	return s
}

func (s *StartGame) Update() error {
	return s.Panel.Update()
}
func (s *StartGame) Draw(screen *ebiten.Image) {
	s.Panel.Draw(screen)
}
func (s *StartGame) GetSceneData() *container.SceneData {
	return s.SceneData
}

func (s *StartGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return container.GameWidth, container.GameHeight
}
