package guiadventure

import (
	"github.com/KONshougun/K-game/game/adventure"
	"github.com/KONshougun/K-game/game/combat"
	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/gui/component"
	"github.com/KONshougun/K-game/gui/component/container"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	battlescene "github.com/KONshougun/K-game/gui/scene/battleScene"
	setcharactersscene "github.com/KONshougun/K-game/gui/scene/setCharactersScene"
	"github.com/KONshougun/K-game/model/area"

	"github.com/hajimehoshi/ebiten/v2"
)

type AdvenureManager struct {
	*container.SceneData
	ActionsBtns []component.Button
	AdvData     *adventure.Adventure
}

func NewAdvenureScene() *AdvenureManager {
	s := &AdvenureManager{
		SceneData: container.NewScene(guihandler.LoadImage("asset/background/ui/Menu_Background.png")),
	}
	return s
}

func (s *AdvenureManager) Update() error {
	sceneData := gui.Game.GetData()
	if sceneData != nil {
		switch sceneData.Type() {
		case guihandler.NewAdventure:
			s.AdvData = sceneData.GetData()[0].(*adventure.Adventure)
			gui.Game.SetData(&guihandler.AdventureActionData{AreaId: area.Plains})
		case guihandler.AdventureAction:
			s.AdvData.Inventory.Area = area.Areas[sceneData.GetData()[0].(area.AreaId)]
			if s.AdvData.Inventory.Area.HasFight {
				btlData := combat.BattleData{
					Opponents: adventure.SpawnAdventureEnemy(s.AdvData),
					Area:      s.AdvData.Inventory.Area,
				}
				scene := setcharactersscene.NewSetCharacterScene(s.AdvData.Inventory.Characters[:], btlData)
				gui.Game.AddScene(scene)
				return nil
			} else {
				gui.Game.AddScene(NewCastleScene(s.AdvData))
				return nil
			}
		case guihandler.ChooseCharacter:
			gui.Game.AddScene(battlescene.NewBattleScene(sceneData.GetData()[0].(*combat.BattleData)))
			return nil
		case guihandler.Fight:
			data := sceneData.GetData()
			if data[1].(bool) {
				//Dare le ricompense...
			} else {
				panic("U lost testa de cazz")
			}
		}
	}

	//ASPETTO L'ADVENTURE_ACTION_DATA
	s.Background.Update()
	return nil
}
func (s *AdvenureManager) Draw(screen *ebiten.Image) {
	s.Background.Draw(screen)
}

func (s *AdvenureManager) GetSceneData() *container.SceneData {
	return s.SceneData
}

func (s *AdvenureManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	return container.GameWidth, container.GameHeight
}
