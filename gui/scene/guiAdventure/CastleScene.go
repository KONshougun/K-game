package guiadventure

import (
	"github.com/KONshougun/K-game/game/adventure"
	"github.com/KONshougun/K-game/gui/component/container"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"

	"github.com/hajimehoshi/ebiten/v2"
)

type CastleScene struct {
	*container.SceneData
	AdvData *adventure.Adventure
}

func NewCastleScene(adventureData *adventure.Adventure) *CastleScene {
	s := &CastleScene{
		SceneData: container.NewScene(guihandler.LoadImage("asset/background/ui/Menu_Background.png")),
		AdvData:   adventureData,
	}
	return s
}

func (s *CastleScene) Update() error {
	return nil
}
func (s *CastleScene) Draw(screen *ebiten.Image) {
	s.Background.Draw(screen)
}
func (s *CastleScene) GetSceneData() *container.SceneData {
	return s.SceneData
}

func (s *CastleScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return container.GameWidth, container.GameHeight
}
