package weapon

import (
	effect "gioco/game/combat/effect"
	"gioco/game/combat/event"
	"gioco/model/stat"
)

type WeaponSet byte

const (
	Std WeaponSet = iota
	Scaglie
	Ossa

	LunaInsanguinata

	SetCount
)

func (w WeaponSet) String() string {
	return [...]string{
		"Std",
		"Scaglie",
		"Ossa",
		"Luna insanguinata",
	}[w]
}

type WeaponData struct {
	Rarity    uint8
	MainStats stat.Stats
	SubStats  stat.Stats
	Effect    []effect.Effect
	EffectStr []string
}

var WeaponStats = [SetCount]WeaponData{
	Std: {
		MainStats: stat.Stats{
			stat.ATK: 5,
		},
		SubStats: stat.Stats{
			stat.ACC: 30,
			stat.SPD: 20,
		},
	},
	Ossa: {
		MainStats: stat.Stats{
			stat.ATK: 5,
		},
		SubStats: stat.Stats{
			stat.ACC: 30,
			stat.HP:  30,
		},
	},
	Scaglie: {
		MainStats: stat.Stats{
			stat.ATK: 5,
		},
		SubStats: stat.Stats{
			stat.ACC: 30,
			stat.DEF: 25,
		},
	},
	LunaInsanguinata: {
		Rarity:    4,
		EffectStr: []string{"DA PENSARE.... (magari qualcosa col crit (niente di sbroccato grazie))"},
		Effect: []effect.Effect{effect.NewEffect(0, event.SetEntityData, effect.After, effect.Normal, func(evData event.EventData) event.EventData {

			if evData.Type() != event.SetEntityData{
				panic("")
			}

			//DA PENSARE.... (magari qualcosa col crit (niente di sbroccato grazie))
			return evData
		})},
		MainStats: stat.Stats{
			stat.ATK:    7,
			stat.SP_ATK: 7,
		},
		SubStats: stat.Stats{
			stat.ACC:       30,
			stat.CRIT_RATE: 5,
			stat.CRIT_DMG:  10,
		},
	},
}
