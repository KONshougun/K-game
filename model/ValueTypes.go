package model

type DamageType byte

const (
	Flat DamageType = iota
	Crit
	SuperCrit
	Pure
	Miss
)

type DamageTarget byte

const (
	Normal DamageTarget = iota
	Health
	ShieldPenetration
)

type SourceType byte

const (
	Ability SourceType = iota
	Passive
	Effect
)

type Damage struct {
	Amount    float64
	DmgType   DamageType
	DmgTarget DamageTarget
}

type HealType byte

const (
	Heal HealType = iota
	Shield
)

type EnergyType byte

const (
	St EnergyType = iota
	Mp
)
