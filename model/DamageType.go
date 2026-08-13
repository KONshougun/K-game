package model

type DamageType byte

const(
	Flat DamageType = iota
	Crit
	SuperCrit
	Pure
	Miss
)

type Damage struct{
	DmgType DamageType
	Amount 	float32
}

type HealType byte

const(
	Heal HealType = iota
	Shield
)