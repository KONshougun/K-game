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
	StCost        float64
	StCostPercent uint8
	MpCost        float64
	MpCostPercent uint8
	SpdCost       uint16

	AtkMultiplier   float64
	SpAtkMultiplier float64
	Acc             uint8
}

func NewAbility(ad AbilityData, f func(opponents []*EntityData) map[int][]model.Damage) Ability {
	return Ability{
		AbilityData: ad,
		Execute:     f,
	}
}

func (a *Ability) CastAbility(stats *stat.Stats) {
	stats[stat.Spd] -= float64(a.SpdCost) - stats[stat.AtkSpd]
	stats[stat.St] *= (100 - float64(a.StCostPercent)) / 100
	stats[stat.St] -= float64(a.StCost)
	stats[stat.Mp] *= (100 - float64(a.MpCostPercent)) / 100
	stats[stat.Mp] -= float64(a.MpCost)
	a.Cooldown = a.CooldownCost
}

func (a *Ability) IsUsable(e *EntityData) (string, bool) {
	if !a.Unlock {
		return "Non sbloccata", false
	}
	if a.Cooldown >= 0 {
		return "In ricarica", false
	}
	if e.Stats[stat.St] < float64(a.StCost)+e.Stats[stat.St]*float64(a.StCostPercent)/100 {
		return "Stamina insufficiente", false
	}
	if e.Stats[stat.Mp] < float64(a.MpCost)+e.Stats[stat.Mp]*float64(a.MpCostPercent)/100 {
		return "Mana insufficiente", false
	}
	return "", true
}

func (a Ability) String() string {
	return comm.Yellow.Write(a.Name)
}
