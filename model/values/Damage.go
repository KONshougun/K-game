package values

type DamageType byte

const (
	Flat DamageType = iota
	Crit
	Pure
	Miss
)

type DamageTarget byte

const (
	Normal DamageTarget = iota
	HealthDmg
	ShieldDmg
)

type Damage struct {
	Amount float64
	Type   DamageType
	Target DamageTarget

	Source SourceType
}

func NewDamage(amount float64, t DamageType, target DamageTarget, s SourceType) *Damage {
	return &Damage{
		Amount: amount,
		Type:   t,
		Target: target,
		Source: s,
	}
}
