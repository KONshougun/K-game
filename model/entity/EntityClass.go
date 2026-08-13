package entity

import "gioco/model/stat"

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

	Enemy
	EnemyBoss

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
	none       float32 = 1.0
	extremeLow float32 = 1.75
	veryLow    float32 = 2.5000
	low        float32 = 4.5
	medium     float32 = 5.9497
	high       float32 = 7.0622
	extreme    float32 = 8.0
	maximum    float32 = 8.7500
)

var EntiityClassesMultiplier = [EntityClassCount]stat.Stats{
	Warrior: {
		stat.HP:        medium,
		stat.ATK:       extreme,
		stat.SP_ATK:    low,
		stat.DEF:       medium,
		stat.SP_DEF:    medium,
		stat.PRF:       medium,
		stat.SP_PRF:    low,
		stat.SPD:       medium,
		stat.CRIT_RATE: veryLow,
		stat.CRIT_DMG:  low,
		stat.CRIT_RES:  low,
		stat.ST:        high,
		stat.MP:        low,
	},

	Magician: {
		stat.HP:        medium,
		stat.ATK:       low,
		stat.SP_ATK:    extreme,
		stat.DEF:       medium,
		stat.SP_DEF:    medium,
		stat.PRF:       low,
		stat.SP_PRF:    medium,
		stat.SPD:       medium,
		stat.CRIT_RATE: veryLow,
		stat.CRIT_DMG:  low,
		stat.CRIT_RES:  low,
		stat.ST:        low,
		stat.MP:        high,
	},

	Brawler: {
		stat.HP:        medium,
		stat.ATK:       high,
		stat.SP_ATK:    high,
		stat.DEF:       medium,
		stat.SP_DEF:    medium,
		stat.PRF:       low,
		stat.SP_PRF:    low,
		stat.SPD:       low,
		stat.CRIT_RATE: veryLow,
		stat.CRIT_DMG:  low,
		stat.CRIT_RES:  medium,
		stat.ST:        medium,
		stat.MP:        medium,
	},

	Marksman: {
		stat.HP:        low,
		stat.ATK:       medium,
		stat.SP_ATK:    medium,
		stat.DEF:       low,
		stat.SP_DEF:    low,
		stat.PRF:       medium,
		stat.SP_PRF:    medium,
		stat.SPD:       medium,
		stat.CRIT_RATE: low,
		stat.CRIT_DMG:  medium,
		stat.CRIT_RES:  low,
		stat.ST:        medium,
		stat.MP:        medium,
	},

	Assassin: {
		stat.HP:        low,
		stat.ATK:       high,
		stat.SP_ATK:    high,
		stat.DEF:       low,
		stat.SP_DEF:    low,
		stat.PRF:       medium,
		stat.SP_PRF:    medium,
		stat.SPD:       medium,
		stat.CRIT_RATE: low,
		stat.CRIT_DMG:  medium,
		stat.CRIT_RES:  low,
		stat.ST:        low,
		stat.MP:        low,
	},

	Tank: {
		stat.HP:        extreme,
		stat.ATK:       low,
		stat.SP_ATK:    low,
		stat.DEF:       high,
		stat.SP_DEF:    high,
		stat.PRF:       low,
		stat.SP_PRF:    low,
		stat.SPD:       low,
		stat.CRIT_RATE: veryLow,
		stat.CRIT_DMG:  low,
		stat.CRIT_RES:  high,
		stat.ST:        low,
		stat.MP:        low,
	},

	Supporter: {
		stat.HP:        high,
		stat.ATK:       low,
		stat.SP_ATK:    low,
		stat.DEF:       medium,
		stat.SP_DEF:    medium,
		stat.PRF:       low,
		stat.SP_PRF:    low,
		stat.SPD:       medium,
		stat.CRIT_RATE: veryLow,
		stat.CRIT_DMG:  low,
		stat.CRIT_RES:  low,
		stat.ST:        medium,
		stat.MP:        medium,
		stat.ST_RG:     low,
		stat.MP_RG:     low,
	},

	Debuffer: {
		stat.HP:        medium,
		stat.ATK:       medium,
		stat.SP_ATK:    medium,
		stat.DEF:       medium,
		stat.SP_DEF:    medium,
		stat.PRF:       low,
		stat.SP_PRF:    low,
		stat.SPD:       high,
		stat.CRIT_RATE: veryLow,
		stat.CRIT_DMG:  low,
		stat.CRIT_RES:  low,
		stat.ST:        low,
		stat.MP:        low,
		stat.ST_RG:     low,
		stat.MP_RG:     low,
	},

	Dualist: {
		stat.HP:        low,
		stat.ATK:       medium,
		stat.SP_ATK:    medium,
		stat.DEF:       low,
		stat.SP_DEF:    low,
		stat.PRF:       high,
		stat.SP_PRF:    high,
		stat.SPD:       low,
		stat.CRIT_RATE: low,
		stat.CRIT_DMG:  low,
		stat.CRIT_RES:  low,
		stat.ST:        medium,
		stat.MP:        medium,
	},

	Area: {
		stat.HP:        low,
		stat.ATK:       high,
		stat.SP_ATK:    high,
		stat.DEF:       medium,
		stat.SP_DEF:    medium,
		stat.PRF:       low,
		stat.SP_PRF:    low,
		stat.SPD:       medium,
		stat.CRIT_RATE: veryLow,
		stat.CRIT_DMG:  medium,
		stat.CRIT_RES:  low,
		stat.ST:        medium,
		stat.MP:        medium,
	},

	Speeder: {
		stat.HP:        low,
		stat.ATK:       low,
		stat.SP_ATK:    low,
		stat.DEF:       low,
		stat.SP_DEF:    low,
		stat.PRF:       low,
		stat.SP_PRF:    low,
		stat.SPD:       high,
		stat.ATK_SPD:   extremeLow,
		stat.CRIT_RATE: veryLow,
		stat.CRIT_DMG:  low,
		stat.CRIT_RES:  low,
		stat.ST:        high,
		stat.MP:        medium,
		stat.ST_RG:     low,
		stat.MP_RG:     low,
	},

	Enemy: {
		stat.HP:        extreme,
		stat.ATK:       high,
		stat.SP_ATK:    high,
		stat.DEF:       high,
		stat.SP_DEF:    high,
		stat.PRF:       medium,
		stat.SP_PRF:    medium,
		stat.SPD:       medium,
		stat.CRIT_RATE: low,
		stat.CRIT_DMG:  medium,
		stat.CRIT_RES:  medium,
		stat.ST:        medium,
		stat.MP:        medium,
		stat.ST_RG:     low,
		stat.MP_RG:     low,
	},

	EnemyBoss: {
		stat.HP:        maximum,
		stat.ATK:       high,
		stat.SP_ATK:    high,
		stat.DEF:       high,
		stat.SP_DEF:    high,
		stat.PRF:       medium,
		stat.SP_PRF:    medium,
		stat.SPD:       medium,
		stat.CRIT_RATE: low,
		stat.CRIT_DMG:  medium,
		stat.CRIT_RES:  medium,
		stat.ST:        medium,
		stat.MP:        medium,
		stat.ST_RG:     medium,
		stat.MP_RG:     medium,
	},
}
