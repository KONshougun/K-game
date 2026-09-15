package battlescene

import (
	"github.com/KONshougun/K-game/game/comm/area"
	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/model/entity/character"
	"github.com/KONshougun/K-game/model/entity/enemy"

	"github.com/hajimehoshi/ebiten/v2"
)

type BattleScene struct {
	Characters *TeamView
	Enemies    *TeamView
	ActionMenu *ActionMenu
	Background *Background
}

func NewBattleScene(characters []character.Character, enemies []enemy.Enemy, area *area.Area) *BattleScene {
	return &BattleScene{
		Characters: NewTeamView(),
		Enemies:    NewTeamView(),
		Background: NewBackground(area),
	}
}

func (s *BattleScene) Update() error {
	s.Background.Update()
	s.Characters.Update()
	s.Enemies.Update()
	s.ActionMenu.Update()
	return nil
}
func (s *BattleScene) Draw(screen *ebiten.Image) {
	s.Background.Draw(screen)
	//s.Characters.Draw(screen.SubImage(image.Rect()))
	//s.Enemies.Draw(screen.SubImage(image.Rect()))
	//s.ActionMenu.Draw(screen.SubImage(image.Rect()))

}

func (s *BattleScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gui.GameWidth, gui.GameHeight
}
