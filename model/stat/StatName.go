package stat

type StatName byte

const (
	Hp StatName = iota

	Atk
	SpAtk

	Def
	SpDef

	Prf
	SpPrf

	Spd
	AtkSpd

	CritRate
	CritDmg
	CritRes

	St
	Mp

	StRg
	MpRg

	Acc
	Eva
	Res
	Qtr

	StatCount
)

func (s StatName) String() string {
	return [...]string{
		"HP",
		"ATK",
		"SP ATK",
		"DEF",
		"SP DEF",
		"PRF",
		"SP PRF",
		"SPD",
		"ATK SPD",
		"CRIT RATE",
		"CRIT DMG",
		"CRIT RES",
		"ST",
		"MP",
		"ST RG",
		"MP RG",
		"ACC",
		"EVA",
		"RES",
		"QTR",
	}[s]
}
