package guiadventure

import (
	"image/color"

	"github.com/KONshougun/K-game/game/adventure"
	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/gui/component"
	"github.com/KONshougun/K-game/gui/component/container"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StartAdventure struct {
	*container.SceneData
}

var lbls = []string{
	"Benvenuto avventuriero\nLe aspetta un'avventura piena di sorprese e incontri inaspettati",
	"La sua avventura inizierà nei panni di un guerriero di nome Knight",
	"Durante la tua avventura affronterai diverse avversità\ndiventa sempre più forte e scopri le leggi che governano questo mondo",
}
var i = 0

func StartAdventureScene() *StartAdventure {
	s := &StartAdventure{
		SceneData: container.NewScene(guihandler.LoadImage("asset/background/ui/Menu_Background.png")),
	}
	initMsg := component.NewButton([4]*ebiten.Image{}, component.NewLabel(lbls[i], color.White,
		component.VeryLargeSize, component.VeryLargeSpace, text.AlignCenter, text.AlignCenter, s.Area,
	), s.Area, nil)
	initMsg.OnClick = func() {
		i++
		if i < len(lbls) {
			initMsg.Text = lbls[i]
		} else if i == len(lbls) {
			gui.Game.RemoveScene(&guihandler.NewAdventureData{AdvData: adventure.InitAdventure()})
			gui.Game.AddScene(NewAdvenureScene())
		}
	}

	if !s.Panel.AddChild(container.Child{C: initMsg}) {
		panic("dake")
	}
	return s
}

func (s *StartAdventure) Update() error {
	return s.Panel.Update()
}
func (s *StartAdventure) Draw(screen *ebiten.Image) {
	s.Panel.Draw(screen)
}

func (s *StartAdventure) GetSceneData() *container.SceneData {
	return s.SceneData
}

func (s *StartAdventure) Layout(outsideWidth, outsideHeight int) (int, int) {
	return container.GameWidth, container.GameHeight
}
