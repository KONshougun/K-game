package battlescene

import (
	"github.com/KONshougun/K-game/game/comm/area"
	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/model/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type BattleScene struct {
	Characters *TeamView
	Enemies    *TeamView
	ActionMenu *ActionMenu
	Background *Background
}

func NewBattleScene(characters, enemies []*entity.EntityData, area *area.Area) *BattleScene {
	return &BattleScene{
		Characters: NewTeamView(characters),
		Enemies:    NewTeamView(enemies),
		Background: NewBackground(area),
	}
}

func (s *BattleScene) Update() error {
	if s.Characters != nil {
		s.Characters.Update()
	}

	if s.Enemies != nil {
		s.Enemies.Update()
	}

	if s.ActionMenu != nil {
		s.ActionMenu.Update()
	}

	return nil
}
func (s *BattleScene) Draw(screen *ebiten.Image) {
	//s.Characters.Draw(screen.SubImage(image.Rect()))
	//s.Enemies.Draw(screen.SubImage(image.Rect()))
	//s.ActionMenu.Draw(screen.SubImage(image.Rect()))

}

func (s *BattleScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gui.GameWidth, gui.GameHeight
}
