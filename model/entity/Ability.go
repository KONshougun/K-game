package entity

import (
	"gioco/game/combat/effect"
	"gioco/game/combat/event"
	"gioco/game/comm"
	"gioco/model/stat"
)

type Ability struct {
	AbilityData

	Cooldown int8
	execute  effect.Effect
}
type AbilityData struct {
	Name             string
	Description      string
	Unlock           bool
	TargetCount      uint8
	DifferentTargets bool

	CooldownCost  int8
	StCost        uint8
	StCostPercent uint8
	MpCost        uint8
	MpCostPercent uint8
	SpdCost       uint16

	AtkMultiplier   uint8
	SpAtkMultiplier uint8
	Acc             uint8
}

func NewAbility(ad AbilityData) Ability {
	return Ability{
		AbilityData: ad,
	}
}
func (a *Ability) SetAbility(f func(a ...any) any) {
	a.execute = effect.NewEffect(effect.AbilityEff, event.CastAbility, effect.None, f)
}

func (a *Ability) CastAbility(stats *stat.Stats) {
	stats[stat.SPD] -= float32(a.SpdCost)
	stats[stat.ST] *= (100 - float32(a.StCostPercent)) / 100
	stats[stat.ST] -= float32(a.StCost)
	stats[stat.MP] *= (100 - float32(a.MpCostPercent)) / 100
	stats[stat.MP] -= float32(a.MpCost)
	a.Cooldown = a.CooldownCost
}

func (a Ability) IsUsable(e Entity) bool {
	if !a.Unlock || a.Cooldown >= 0 ||
		e.Stats[stat.ST] < float32(a.StCost)+e.Stats[stat.ST]*float32(a.StCostPercent)/100 ||
		e.Stats[stat.MP] < float32(a.MpCost)+e.Stats[stat.MP]*float32(a.MpCostPercent)/100 {
		return false
	}
	return true
}

func (a Ability) String() string {
	return comm.Yellow.Write(a.Name)
}
