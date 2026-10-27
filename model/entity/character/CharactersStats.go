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
	Knight: {
		stat.Hp:       203,
		stat.Atk:      21,
		stat.SpAtk:    19.5,
		stat.Def:      15.7,
		stat.SpDef:    15.5,
		stat.Prf:      5.2,
		stat.SpPrf:    5,
		stat.Spd:      515,
		stat.AtkSpd:   250,
		stat.CritRate: 9.85,
		stat.CritDmg:  10,
		stat.CritRes:  5,
		stat.St:       78,
		stat.Mp:       75,
		stat.StRg:     20,
		stat.MpRg:     20,
		stat.Acc:      30,
		stat.Eva:      5,
		stat.Res:      5,
		stat.Qtr:      5,
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
