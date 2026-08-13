package stat

type StatName byte

const(
	HP StatName = iota

	ATK
	SP_ATK

	DEF
	SP_DEF

	PRF
	SP_PRF

	SPD
	ATK_SPD

	CRIT_RATE
	CRIT_DMG
	CRIT_RES

	ST
	MP

	ST_RG
	MP_RG

	ACC
	EVA
	RES
	QTR

	StatCount
)

func (s StatName) String() string{
	return [...]string{
		"HP",
		"ATK",
		"SP_ATK",
		"DEF",
		"SP_DEF",
		"PRF",
		"SP_PRF",
		"SPD",
		"ATK_SPD",
		"CRIT_RATE",
		"CRIT_DMG",
		"CRIT_RES",
		"ST",
		"MP",
		"ST_RG",
		"MP_RG",
		"ACC",
		"EVA",
		"RES",
		"QTR",
	}[s]
}