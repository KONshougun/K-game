package battlescene

import (
	"github.com/KONshougun/K-game/game/combat/effect"
	"github.com/KONshougun/K-game/game/combat/event"
	gameHandler "github.com/KONshougun/K-game/game/combat/handler"
	"github.com/KONshougun/K-game/game/comm/area"
	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/gui/component"
	"github.com/KONshougun/K-game/gui/handler"
	"github.com/KONshougun/K-game/model/entity/character"
	"github.com/KONshougun/K-game/model/entity/enemy"
	"github.com/KONshougun/K-game/model/stat"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	GameWidth  = 1280
	GameHeight = 720
)

type BattleScene struct {
	//Characters *TeamView
	//Enemies    *TeamView
	ActionMenu *ActionPanel
	*component.Background
}

func NewBattleScene(characters []character.Character, enemies []enemy.Enemy, area area.AreaId) *BattleScene {

	for i := range characters {
		characters[i].GetCharacter().TriggerEffects(event.SetEntityData, effect.None, &gameHandler.SetEntityData{
			FlatStats:    stat.Stats{},
			PercentStats: stat.Stats{},
		})
		characters[i].GetCharacter().Shield = characters[i].GetCharacter().Stats[stat.Hp] / 2
	}
	return &BattleScene{
		//Characters: NewTeamView(),
		//Enemies:    NewTeamView(),

		ActionMenu: NewActionMenu(characters[0].GetCharacter().GetEntity(), []bool{true, true, true, false}, component.Area{
			X:      0,
			Y:      GameHeight * 0.7,
			Width:  GameWidth * 0.7,
			Height: GameHeight * 0.3,
		}),
		Background: component.NewBackground(handler.GetCombatBackground(area), component.Area{
			Width:  GameWidth,
			Height: GameHeight,
		}),
	}
}

func (s *BattleScene) Update() error {
	s.Background.Update()
	//s.Characters.Update()
	//s.Enemies.Update()
	s.ActionMenu.Update()
	return nil
}
func (s *BattleScene) Draw(screen *ebiten.Image) {
	s.Background.Draw(screen)
	//s.Characters.Draw(screen.SubImage(image.Rect()))
	//s.Enemies.Draw(screen.SubImage(image.Rect()))
	s.ActionMenu.Draw(screen)

}

func (s *BattleScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gui.GameWidth, gui.GameHeight
}
