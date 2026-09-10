package entity

import (
	"github.com/KONshougun/K-game/game/comm"
	"github.com/KONshougun/K-game/model"
	"github.com/KONshougun/K-game/model/stat"
)

type Ability struct {
	AbilityData

	Cooldown int8
	Execute  func(opponents []*EntityData) map[int][]model.Damage
}
type AbilityData struct {
	Name             string
	Description      string
	Unlock           bool
	TargetCount      uint8
	DifferentTargets bool

	CooldownCost  int8
	StCost        float32
	StCostPercent uint8
	MpCost        float32
	MpCostPercent uint8
	SpdCost       uint16

	AtkMultiplier   float32
	SpAtkMultiplier float32
	Acc             uint8
}

func NewAbility(ad AbilityData, f func(opponents []*EntityData) map[int][]model.Damage) Ability {
	return Ability{
		AbilityData: ad,
		Execute:     f,
	}
}

func (a *Ability) CastAbility(stats *stat.Stats) {
	stats[stat.SPD] -= float32(a.SpdCost) - stats[stat.ATK_SPD]
	stats[stat.ST] *= (100 - float32(a.StCostPercent)) / 100
	stats[stat.ST] -= float32(a.StCost)
	stats[stat.MP] *= (100 - float32(a.MpCostPercent)) / 100
	stats[stat.MP] -= float32(a.MpCost)
	a.Cooldown = a.CooldownCost
}

func (a *Ability) IsUsable(e *EntityData) (string, bool) {
	if !a.Unlock {
		return "Non sbloccata", false
	}
	if a.Cooldown >= 0{
		return "In ricarica", false
	}
	if e.Stats[stat.ST] < float32(a.StCost)+e.Stats[stat.ST]*float32(a.StCostPercent)/100 {
		return "Stamina insufficiente", false
	}
	if e.Stats[stat.MP] < float32(a.MpCost)+e.Stats[stat.MP]*float32(a.MpCostPercent)/100 {
		return "Mana insufficiente", false
	}
	return "", true
}

func (a Ability) String() string {
	return comm.Yellow.Write(a.Name)
}
