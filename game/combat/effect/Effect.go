package effect

import "gioco/game/combat/event"

type IdEffect byte

const (
	FightBaseEff IdEffect = iota
	BuffEff
	AbilityEff
	PassiveEff
	SetEff
	GetAbilityEff
	QtrEff
	ReviveEff

	IdEffCount
)

type Effect struct {
	Id       uint16
	IdEff    IdEffect
	Event    event.EventType
	Trigger  TriggerTiming
	Priority PriorityType
	Activate func(a ...any) any
}

type PriorityType uint8

const (
	Slow PriorityType = iota
	Normal
	Fast
	ExtremeFast
)

func NewEffect(id uint16, idEff IdEffect, e event.EventType, trg TriggerTiming, priority PriorityType, f func(a ...any) any) Effect {

	return Effect{
		Id:       id,
		IdEff:    idEff,
		Event:    e,
		Trigger:  trg,
		Priority: priority,
		Activate: func(a ...any) any {
			idEff.HandleEffect(a)
			return f(a)
		},
	}
}

func (idEff IdEffect) HandleEffect(value any) {
	switch idEff {
	case QtrEff:
		_, ok := value.(bool)
		if !ok {
			panic("cazzo è possibile")
		}
	}
}
