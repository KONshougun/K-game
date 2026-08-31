package event

type EventType byte

const (
	SetEntityData EventType = iota
	StartFight 
	EndFight
	StartTurn
	EndTurn

	BuffApplied
	BuffExpired

	DamageDealt
	DamageTaken

	QtrTrigged

	HealDealt
	HealTaken

	EnergySpent
	EnergyRecovered

	Revived
	Resist
	Died
	//...

	EventCount
)

type EventData interface {
	Type() EventType
	GetData() []any
}
