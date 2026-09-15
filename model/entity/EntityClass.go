package entity

import "github.com/KONshougun/K-game/model/stat"

type EntityClass byte

const (
	Warrior EntityClass = iota
	Magician
	Brawler
	Marksman
	Assassin
	Tank
	Supporter
	Debuffer
	Dualist
	Area
	Speeder

	EnemyClass
	EnemyBossClass

	EntityClassCount
)

func (e EntityClass) String() string {
	return [...]string{
		"Warrior",
		"Magician",
		"Brawler",
		"Marksman",
		"Assassin",
		"Tank",
		"Supporter",
		"Debuffer",
		"Dualist",
		"Area",
		"Enemy",
		"Enemy boss",
	}[e]
}

//Moltiplicatori
const (
	none       float64 = 1.0
	extremeLow float64 = 1.75
	veryLow    float64 = 2.5000
	low        float64 = 4.5
	medium     float64 = 5.9497
	high       float64 = 7.0622
	extreme    float64 = 8.0
	maximum    float64 = 8.7500
)

var EntiityClassesMultiplier = [EntityClassCount]stat.Stats{
	Warrior: {
		stat.Hp:       medium,
		stat.Atk:      extreme,
		stat.SpAtk:    low,
		stat.Def:      medium,
		stat.SpDef:    medium,
		stat.Prf:      medium,
		stat.SpPrf:    low,
		stat.Spd:      medium,
		stat.CritRate: veryLow,
		stat.CritDmg:  low,
		stat.CritRes:  low,
		stat.St:       high,
		stat.Mp:       low,
	},

	Magician: {
		stat.Hp:       medium,
		stat.Atk:      low,
		stat.SpAtk:    extreme,
		stat.Def:      medium,
		stat.SpDef:    medium,
		stat.Prf:      low,
		stat.SpPrf:    medium,
		stat.Spd:      medium,
		stat.CritRate: veryLow,
		stat.CritDmg:  low,
		stat.CritRes:  low,
		stat.St:       low,
		stat.Mp:       high,
	},

	Brawler: {
		stat.Hp:       medium,
		stat.Atk:      high,
		stat.SpAtk:    high,
		stat.Def:      medium,
		stat.SpDef:    medium,
		stat.Prf:      low,
		stat.SpPrf:    low,
		stat.Spd:      low,
		stat.CritRate: veryLow,
		stat.CritDmg:  low,
		stat.CritRes:  medium,
		stat.St:       medium,
		stat.Mp:       medium,
	},

	Marksman: {
		stat.Hp:       low,
		stat.Atk:      medium,
		stat.SpAtk:    medium,
		stat.Def:      low,
		stat.SpDef:    low,
		stat.Prf:      medium,
		stat.SpPrf:    medium,
		stat.Spd:      medium,
		stat.CritRate: low,
		stat.CritDmg:  medium,
		stat.CritRes:  low,
		stat.St:       medium,
		stat.Mp:       medium,
	},

	Assassin: {
		stat.Hp:       low,
		stat.Atk:      high,
		stat.SpAtk:    high,
		stat.Def:      low,
		stat.SpDef:    low,
		stat.Prf:      medium,
		stat.SpPrf:    medium,
		stat.Spd:      medium,
		stat.CritRate: low,
		stat.CritDmg:  medium,
		stat.CritRes:  low,
		stat.St:       low,
		stat.Mp:       low,
	},

	Tank: {
		stat.Hp:       extreme,
		stat.Atk:      low,
		stat.SpAtk:    low,
		stat.Def:      high,
		stat.SpDef:    high,
		stat.Prf:      low,
		stat.SpPrf:    low,
		stat.Spd:      low,
		stat.CritRate: veryLow,
		stat.CritDmg:  low,
		stat.CritRes:  high,
		stat.St:       low,
		stat.Mp:       low,
	},

	Supporter: {
		stat.Hp:       high,
		stat.Atk:      low,
		stat.SpAtk:    low,
		stat.Def:      medium,
		stat.SpDef:    medium,
		stat.Prf:      low,
		stat.SpPrf:    low,
		stat.Spd:      medium,
		stat.CritRate: veryLow,
		stat.CritDmg:  low,
		stat.CritRes:  low,
		stat.St:       medium,
		stat.Mp:       medium,
		stat.StRg:     low,
		stat.MpRg:     low,
	},

	Debuffer: {
		stat.Hp:       medium,
		stat.Atk:      medium,
		stat.SpAtk:    medium,
		stat.Def:      medium,
		stat.SpDef:    medium,
		stat.Prf:      low,
		stat.SpPrf:    low,
		stat.Spd:      high,
		stat.CritRate: veryLow,
		stat.CritDmg:  low,
		stat.CritRes:  low,
		stat.St:       low,
		stat.Mp:       low,
		stat.StRg:     low,
		stat.MpRg:     low,
	},

	Dualist: {
		stat.Hp:       low,
		stat.Atk:      medium,
		stat.SpAtk:    medium,
		stat.Def:      low,
		stat.SpDef:    low,
		stat.Prf:      high,
		stat.SpPrf:    high,
		stat.Spd:      low,
		stat.CritRate: low,
		stat.CritDmg:  low,
		stat.CritRes:  low,
		stat.St:       medium,
		stat.Mp:       medium,
	},

	Area: {
		stat.Hp:       low,
		stat.Atk:      high,
		stat.SpAtk:    high,
		stat.Def:      medium,
		stat.SpDef:    medium,
		stat.Prf:      low,
		stat.SpPrf:    low,
		stat.Spd:      medium,
		stat.CritRate: veryLow,
		stat.CritDmg:  medium,
		stat.CritRes:  low,
		stat.St:       medium,
		stat.Mp:       medium,
	},

	Speeder: {
		stat.Hp:       low,
		stat.Atk:      low,
		stat.SpAtk:    low,
		stat.Def:      low,
		stat.SpDef:    low,
		stat.Prf:      low,
		stat.SpPrf:    low,
		stat.Spd:      high,
		stat.AtkSpd:   extremeLow,
		stat.CritRate: veryLow,
		stat.CritDmg:  low,
		stat.CritRes:  low,
		stat.St:       high,
		stat.Mp:       medium,
		stat.StRg:     low,
		stat.MpRg:     low,
	},

	EnemyClass: {
		stat.Hp:       extreme,
		stat.Atk:      high,
		stat.SpAtk:    high,
		stat.Def:      high,
		stat.SpDef:    high,
		stat.Prf:      medium,
		stat.SpPrf:    medium,
		stat.Spd:      medium,
		stat.CritRate: low,
		stat.CritDmg:  medium,
		stat.CritRes:  medium,
		stat.St:       medium,
		stat.Mp:       medium,
		stat.StRg:     low,
		stat.MpRg:     low,
	},

	EnemyBossClass: {
		stat.Hp:       maximum,
		stat.Atk:      high,
		stat.SpAtk:    high,
		stat.Def:      high,
		stat.SpDef:    high,
		stat.Prf:      medium,
		stat.SpPrf:    medium,
		stat.Spd:      medium,
		stat.CritRate: low,
		stat.CritDmg:  medium,
		stat.CritRes:  medium,
		stat.St:       medium,
		stat.Mp:       medium,
		stat.StRg:     medium,
		stat.MpRg:     medium,
	},
}
