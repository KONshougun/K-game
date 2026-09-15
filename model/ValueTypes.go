package model

type DamageType byte

const (
	Flat DamageType = iota
	Crit
	SuperCrit
	Pure
	Miss
)

type SourceType byte
const(
	Ability SourceType = iota
	Passive
	Effect
)

type Damage struct {
	DmgType DamageType
	Amount  float64
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
