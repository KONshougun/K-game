package weapon

import (
	effect "github.com/KONshougun/K-game/game/combat/effect"
	"github.com/KONshougun/K-game/game/combat/event"
	"github.com/KONshougun/K-game/model/stat"
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
			stat.Atk: 5,
		},
		SubStats: stat.Stats{
			stat.Acc: 30,
			stat.Spd: 20,
		},
	},
	Ossa: {
		MainStats: stat.Stats{
			stat.Atk: 5,
		},
		SubStats: stat.Stats{
			stat.Acc: 30,
			stat.Hp:  30,
		},
	},
	Scaglie: {
		MainStats: stat.Stats{
			stat.Atk: 5,
		},
		SubStats: stat.Stats{
			stat.Acc: 30,
			stat.Def: 25,
		},
	},
	LunaInsanguinata: {
		Rarity:    4,
		EffectStr: []string{"DA PENSARE.... (magari qualcosa col crit (niente di sbroccato grazie))"},
		Effect: []effect.Effect{effect.NewEffect(0, event.SetEntityData, effect.After, effect.Normal, func(evData event.EventData) event.EventData {

			if evData.Type() != event.SetEntityData {
				panic("")
			}

			//DA PENSARE.... (magari qualcosa col crit (niente di sbroccato grazie))
			return evData
		})},
		MainStats: stat.Stats{
			stat.Atk:   7,
			stat.SpAtk: 7,
		},
		SubStats: stat.Stats{
			stat.Acc:      30,
			stat.CritRate: 5,
			stat.CritDmg:  10,
		},
	},
}
