package entity

import (
	"gioco/game/combat/effect"
	"gioco/game/combat/event"
)

type Passive struct {
	Effect      effect.Effect
	Unlock      bool
	Description string
}

func NewPassive(unlock bool, description string, activasionType event.EventType, trg effect.TriggerTiming, f func(a ...any) any) Passive {
	return Passive{
		Unlock:      unlock,
		Description: description,
		Effect:      effect.NewEffect(effect.PassiveEff, activasionType, trg, f),
	}
}
