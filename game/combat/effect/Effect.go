package effect

import "gioco/game/combat/event"

type IdEffect byte

const (
	FightBaseEff IdEffect = iota
	BuffEff
	AbilityEff
	PassiveEff
	SetEff
	QtrEff
	Revive

	IdEffCount
)

type Effect struct {
	IdEff    IdEffect
	Event    event.EventType
	Trigger  TriggerTiming
	Activate func(a ...any) any
}

func NewEffect(id IdEffect, e event.EventType, trg TriggerTiming, f func(a ...any) any) Effect {
	return Effect{
		IdEff:    id,
		Event:    e,
		Trigger:  trg,
		Activate: f,
	}
}

func (idEff IdEffect)HandleEffect(value any)any{
	switch idEff{
	case QtrEff:
		v, ok := value.(bool)
		if ok{
			return v
		}else{
			panic("cazzo è possibile")
		}
	default:
		panic("daje roma dajeeeee")
	}
}
