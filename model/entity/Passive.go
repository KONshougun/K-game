package entity

import (
	"github.com/KONshougun/K-game/game/combat/effect"
	"github.com/KONshougun/K-game/game/combat/event"
)

type Passive struct {
	Effect      effect.Effect
	Unlock      bool
	Description string
}

func NewPassive(unlock bool, description string, ev event.EventType, trg effect.TriggerTiming, prio effect.PriorityType, f func(evt event.EventData) event.EventData) Passive {
	return Passive{
		Unlock:      unlock,
		Description: description,
		Effect:      effect.NewEffect(0, ev, trg, prio, f),
	}
}
