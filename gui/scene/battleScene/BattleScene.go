package battlescene

import (
	"github.com/KONshougun/K-game/game/combat"
	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/gui/component"
	"github.com/KONshougun/K-game/gui/component/container"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	"github.com/KONshougun/K-game/model"
	"github.com/KONshougun/K-game/model/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type BattleScene struct {
	*container.SceneData

	ActionPanel *ActionPanel

	BtlData combat.BattleData
}

func NewBattleScene(btlData combat.BattleData) *BattleScene {

	s := &BattleScene{
		SceneData:   container.NewScene(guihandler.GetCombatBackground(btlData.Area.Id)),
		ActionPanel: nil,
		BtlData:     btlData,
	}
	charactersView := container.NewTeam(btlData.Team.GetEntities(), component.Area{
		X:      container.GameWidth * 0.05,
		Y:      container.GameHeight * 0.1,
		Width:  container.GameWidth * 0.4,
		Height: container.GameHeight * 0.5,
	})
	enemiesView := container.NewTeam(btlData.Opponents.GetEntities(), component.Area{
		X:      container.GameWidth * 0.55,
		Y:      container.GameHeight * 0.1,
		Width:  container.GameWidth * 0.4,
		Height: container.GameHeight * 0.5,
	})
	if !s.AddChild(container.Child{C: charactersView},
		container.Child{C: enemiesView}) {
		panic("dajeeeee")
	}
	s.BtlData.StartFight()
	return s
}

func (s *BattleScene) Update() error {
	s.Panel.Update()
	//Controllo se un team ha perso
	if len(s.BtlData.Team) == 0 || len(s.BtlData.Opponents) == 0 {
		gui.Game.RemoveScene(&guihandler.FightData{
			BtlData: s.BtlData,
			HasWon:  s.BtlData.EndFight(),
		})
	}

	//Controllo se sto spettando un output
	if s.ActionPanel == nil {
		//Se non sto aspettando, cerco il prossimo attacco
		eType, actorIdx, abilities := s.BtlData.StartTurn()
		switch eType {
		case entity.Character:
			actor := s.BtlData.Team[actorIdx].GetCharacter().GetEntity()
			//Controllo se c'è almeno un'abilità da poter utilizzare
			var found bool
			for i := range abilities {
				if abilities[i] {
					found = true
				}
			}
			if found {
				//Chiedo l'input all'utente
				s.ActionPanel = NewActionMenu(s, actor, abilities, component.Area{
					X:      0,
					Y:      s.Height * 0.7,
					Width:  s.Width * 0.6,
					Height: s.Height * 0.3,
				})
			} else {
				s.Attack(eType, actor, -1)
			}
		case entity.Enemy:
			actor := s.BtlData.Opponents[actorIdx].GetEnemy()
			s.Attack(eType, actor.GetEntity(), actor.GetAbility(abilities))
		}
	} else {
		s.ActionPanel.Update()
	}
	return nil
}
func (s *BattleScene) Draw(screen *ebiten.Image) {
	s.Panel.Draw(screen)
	if s.ActionPanel != nil {
		s.ActionPanel.Draw(screen)
	}
}

func (s *BattleScene) GetSceneData() *container.SceneData {
	return s.SceneData
}

func (s *BattleScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return container.GameWidth, container.GameHeight
}

func (s *BattleScene) Attack(eType entity.EntityType, actor *entity.EntityData, abilityIdx int) {
	if eType == entity.Character {
		s.ActionPanel = nil
	}
	targets := s.BtlData.StartAttack(eType, actor, abilityIdx)
	if len(targets) > 0 {
		s.ShowDmg(targets)
		s.BtlData.CheckDeath(targets)
	}
	//magari dovrei pure controllare un possibile qtr / magari un endTurn
	/*qtr :=*/
	s.BtlData.EndTurn(eType, actor)
}

// Mostra i danni subiti a schermo
func (s *BattleScene) ShowDmg(targets map[int][]model.Damage) {

}
