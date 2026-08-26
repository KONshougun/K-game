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

func NewPassive(unlock bool, description string, ev event.EventType, trg effect.TriggerTiming, prio effect.PriorityType, f func(a ...any) any) Passive {
	return Passive{
		Unlock:      unlock,
		Description: description,
		Effect:      effect.NewEffect(0, effect.PassiveEff, ev, trg, prio, f),
	}
}
