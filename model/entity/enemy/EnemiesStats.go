package enemy

import "github.com/KONshougun/K-game/model/stat"

type EnemyName uint16

const (
	Bandit EnemyName = iota
	Lupo
	Goblin
	Scheletro
	Zombie
	LupoMannaro
	Vampiro
	Fantasma
	Coccodrillo
	PaguroGigante
	Tritone
	Nokk

	//BOSS
	Kayn
	Akari
	Thalorn
)

var EnemiesBaseStats = map[EnemyName]stat.Stats{
	Bandit: {
		stat.Hp:       100,
		stat.Atk:      25,
		stat.SpAtk:    15,
		stat.Def:      10,
		stat.SpDef:    7,
		stat.Prf:      4,
		stat.SpPrf:    3,
		stat.Spd:      350,
		stat.AtkSpd:   220,
		stat.CritRate: 8,
		stat.CritDmg:  9,
		stat.CritRes:  5,
		stat.St:       70,
		stat.Mp:       40,
		stat.StRg:     15,
		stat.MpRg:     10,
		stat.Acc:      30,
		stat.Eva:      5,
		stat.Res:      5,
		stat.Qtr:      5,
	},
}

func (e EnemyName) String() string {
	return [...]string{
		"Bandit",
		"Lupo",
		"Goblin",
		"Scheletro",
		"Zombie",
		"Lupo Mannaro",
		"Vampiro",
		"Fantasma",
		"Coccodrillo",
		"Paguro Gigante",
		"Tritone",
		"Nokk",

		"Kayn",
		"Akari",
		"Thalorn",
	}[e]
}
