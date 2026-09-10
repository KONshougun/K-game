package character

import "github.com/KONshougun/K-game/model/stat"

type CharacterName uint16

const (
	Knight CharacterName = iota
	Hunter
	Ironforge
	Invoker
	Lightbringer
	Darkwielder
	Viking
	Hawk

	CharacterCount
)

var CharactersBaseStats = map[CharacterName]stat.Stats{
	Knight : {
		stat.HP:        203,
		stat.ATK:       21,
		stat.SP_ATK:    19.5,
		stat.DEF:       15.7,
		stat.SP_DEF:    15.5,
		stat.PRF:       5.2,
		stat.SP_PRF:    5,
		stat.SPD:       515,
		stat.ATK_SPD:   250,
		stat.CRIT_RATE: 9.85,
		stat.CRIT_DMG:  10,
		stat.CRIT_RES:  5,
		stat.ST:        78,
		stat.MP:        75,
		stat.ST_RG:     20,
		stat.MP_RG:     20,
		stat.ACC:       30,
		stat.EVA:       5,
		stat.RES:       5,
		stat.QTR:       5,
	},
}

func (c CharacterName) String() string {
	return [...]string{
		"Knight",
		"Hunter",
		"Ironforge",
		"Invoker",
		"Lightbringer",
		"Darkwielder",
		"Viking",
		"Hawk",
	}[c]
}
