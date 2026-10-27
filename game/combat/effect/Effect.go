package effect

import "github.com/KONshougun/K-game/game/combat/event"

//type IdEffect byte

/*const (
	FightBaseEff IdEffect = iota
	BuffEff
	AbilityEff
	PassiveEff
	SetEff
	GetAbilityEff
	QtrEff
	ReviveEff

	IdEffCount
)*/

type Effect struct {
	Id       uint16
	Event    event.EventType
	Trigger  TriggerTiming
	Priority PriorityType
	Activate func(evData event.EventData) event.EventData
}

type PriorityType uint8

const (
	Slow PriorityType = iota
	Normal
	Fast
	ExtremeFast
)

func NewEffect(id uint16, e event.EventType, trg TriggerTiming, priority PriorityType, f func(evData event.EventData) event.EventData) Effect {
	return Effect{
		Id:       id,
		Event:    e,
		Trigger:  trg,
		Priority: priority,
		Activate: func(evData event.EventData) event.EventData {
			if evData.Type() == e {
				return f(evData)
			} else {
				panic("siuummm")
			}
		},
	}
}
