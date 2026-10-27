package setcharactersscene

import (
	"github.com/KONshougun/K-game/game/combat"
	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/gui/component/container"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	"github.com/KONshougun/K-game/model/entity/character"

	"github.com/hajimehoshi/ebiten/v2"
)

type SetCharacterScene struct {
	*container.SceneData

	CharactersBtns      *container.ButtonGroup
	AvailableCharacters []int
	Selectable          int

	Characters []character.Character
	BtlData    combat.BattleData
}

func NewSetCharacterScene(characters []character.Character, btlData combat.BattleData) *SetCharacterScene {
	s := &SetCharacterScene{
		SceneData:           container.NewScene(guihandler.LoadImage("asset/background/ui/Menu_Background.png")),
		AvailableCharacters: make([]int, 0, len(characters)),
		Characters:          characters,
		Selectable:          len(btlData.Opponents),
		BtlData:             btlData,
	}
	//TODO charactersBtns
	//CharactersBtns:      component.NewButtonGroup(len(characters), len(btlData.Opponents)),

	for i, c := range characters {
		//TODO controllare se sono ancora vivi
		if c != nil {
			s.AvailableCharacters = append(s.AvailableCharacters, i)
		} else {
			//s.CharactersBtns.Btns[i].Locked = true
		}
	}
	return s
}

func (s *SetCharacterScene) Update() error {

	if len(s.AvailableCharacters) <= s.Selectable {
		data := &guihandler.ChooseCharacterData{
			BtlData: s.BtlData,
		}
		for i := range s.AvailableCharacters {
			data.BtlData.Team = append(data.BtlData.Team, s.Characters[i])
		}
		gui.Game.RemoveScene(data)
	}
	s.Background.Update()
	return nil
}
func (s *SetCharacterScene) Draw(screen *ebiten.Image) {
	s.Background.Draw(screen)
}
func (s *SetCharacterScene) GetSceneData() *container.SceneData {
	return s.SceneData
}

func (s *SetCharacterScene) Layout(outsideWidth, outsideHeight int) (int, int) {
	return container.GameWidth, container.GameHeight
}
