package battlescene

import (
	"github.com/KONshougun/K-game/game/comm/area"
	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/gui/component"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	"github.com/KONshougun/K-game/model/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GameWidth  = 1280
	GameHeight = 720
)

type BattleScene struct {
	Characters *TeamView
	Enemies    *TeamView
	ActionMenu *ActionPanel
	*component.Background
}

func NewBattleScene(characters, enemies []*entity.EntityData, area area.AreaId) *BattleScene {

	return &BattleScene{
		Characters: NewTeamView(characters, component.Area{
			X:      GameWidth * 0.05,
			Y:      GameHeight * 0.1,
			Width:  GameWidth * 0.4,
			Height: GameHeight * 0.5,
		}),
		Enemies: NewTeamView(enemies, component.Area{
			X:      GameWidth * 0.55,
			Y:      GameHeight * 0.1,
			Width:  GameWidth * 0.4,
			Height: GameHeight * 0.5,
		}),

		ActionMenu: NewActionMenu(characters[0], []bool{true, true, true, false}, component.Area{
			X:      0,
			Y:      GameHeight * 0.7,
			Width:  GameWidth * 0.7,
			Height: GameHeight * 0.3,
		}),
		Background: component.NewBackground(guihandler.GetCombatBackground(area), component.Area{
			Width:  GameWidth,
			Height: GameHeight,
		}),
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
	s.Characters.Draw(screen)
	s.Enemies.Draw(screen)
	s.ActionMenu.Draw(screen)
}

func (s *BattleScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gui.GameWidth, gui.GameHeight
}
