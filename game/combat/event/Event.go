package event

type EventType byte

/*type EventData interface {
	EventType() EventType
	EventReturnData()
}*/

const (
	StartFight EventType = iota
	EndFight
	StartTurn
	EndTurn

	CastAbility
	Equipment

	BuffApplied
	BuffExpired

	DamageDealt
	DamageTaken

	QtrTrigged

	HealDealt
	HealDTaken

	EnergySpent
	EnergyRecovered

	Revived
	Died
	//...

	EventCount
)
