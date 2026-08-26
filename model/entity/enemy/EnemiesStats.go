package enemy

import "gioco/model/stat"

type EnemyName uint16

const (
	Bandito EnemyName = iota
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
	Bandito: {
		stat.HP:        100,
		stat.ATK:       25,
		stat.SP_ATK:    15,
		stat.DEF:       10,
		stat.SP_DEF:    7,
		stat.PRF:       4,
		stat.SP_PRF:    3,
		stat.SPD:       350,
		stat.ATK_SPD:   220,
		stat.CRIT_RATE: 8,
		stat.CRIT_DMG:  9,
		stat.CRIT_RES:  5,
		stat.ST:        70,
		stat.MP:        40,
		stat.ST_RG:     15,
		stat.MP_RG:     10,
		stat.ACC:       30,
		stat.EVA:       5,
		stat.RES:       5,
		stat.QTR:       5,
	},
}

func (e EnemyName) String() string {
	return [...]string{
		"Bandito",
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
