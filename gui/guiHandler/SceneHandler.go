package guihandler

import (
	"github.com/KONshougun/K-game/game/adventure"
	"github.com/KONshougun/K-game/game/combat"
	"github.com/KONshougun/K-game/model/area"
)

type SceneType byte

const (
	NewAdventure SceneType = iota
	AdventureAction
	Fight
	ChooseCharacter
)

type Data interface {
	GetData() []any
	Type() SceneType
}

type NewAdventureData struct {
	AdvData *adventure.Adventure
}

func (s *NewAdventureData) GetData() []any { return []any{s.AdvData} }
func (NewAdventureData) Type() SceneType   { return NewAdventure }

type AdventureActionData struct {
	AreaId area.AreaId
}

func (s *AdventureActionData) GetData() []any { return []any{s.AreaId} }
func (AdventureActionData) Type() SceneType   { return AdventureAction }

type FightData struct {
	BtlData combat.BattleData
	HasWon  bool
}

func (s *FightData) GetData() []any { return []any{s.BtlData, s.HasWon} }
func (FightData) Type() SceneType   { return Fight }

type ChooseCharacterData struct {
	BtlData combat.BattleData
}

func (s *ChooseCharacterData) GetData() []any { return []any{s.BtlData} }
func (ChooseCharacterData) Type() SceneType   { return ChooseCharacter }
