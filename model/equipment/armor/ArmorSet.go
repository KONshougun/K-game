package armor

import (
	"gioco/game/combat/effect"
	"gioco/game/combat/event"
	"gioco/model/stat"
)

type ArmorSet byte

const (
	Ferro ArmorSet = iota
	Lupo
	Ossa
	Scaglie
	Ketos
	Fenice

	SetCount
)

func (a ArmorSet) String() string {
	return [...]string{
		"Ferro",
		"Pelliccia di lupo",
		"Ossa",
		"Scaglie",
		"Armatura di Ketos",
		"Armatura della Fenice",
	}[a]
}

type ArmorData struct {
	Effects    [2]effect.Effect
	EffectsStr [2]string
}

var ArmorEffects = [SetCount]ArmorData{
	Ferro: {
		[2]effect.Effect{
			0: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.ATK] += 20
				return nil
			}),
			1: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.HP] += 20
				subStats[stat.DEF] += 16
				return nil
			}),
		},
		[2]string{"Aumenta l'attacco di 20", "Aumenta gli hp di 20 la difesa di 16"},
	},
	Lupo: {
		[2]effect.Effect{
			0: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.ST] += 25
				return nil
			}),
			1: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.ATK] += 15
				subStats[stat.SPD] += 20
				return nil
			}),
		},
		[2]string{"Aumenta la stamina di 25", "Aumenta l'attacco di 15 e la velocità di 20"},
	},
	Ossa: {
		[2]effect.Effect{
			0: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.DEF] += 10
				return nil
			}),
			1: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.ATK] += 20
				subStats[stat.HP] += 25
				return nil
			}),
		},
		[2]string{"Aumenta la difesa di 10", "Aumenta gli hp di 25 e l'attacco di 20"},
	},
	Scaglie: {
		[2]effect.Effect{
			0: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.MP] += 25
				return nil
			}),
			1: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.SP_ATK] += 20
				subStats[stat.DEF] += 20
				return nil
			}),
		},
		[2]string{"Aumenta il mana di 25", "Aumenta l'attacco speciale di 20 e la difesa di 20"},
	},
	Ketos: {
		[2]effect.Effect{
			0: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.HP] += 50
				subStats[stat.SPD] += 100
				return nil
			}),
			1: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.DEF] += 25
				subStats[stat.SP_DEF] += 25
				return nil
			}),
		},
		[2]string{"Aumenta gli hp di 50 e la velocità di 100",
			"Aumenta la difesa e la difesa speciale di 25"},
	},
	Fenice: {
		[2]effect.Effect{
			0: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.DEF] += 20
				subStats[stat.MP] += 30
				return nil
			}),
			1: effect.NewEffect(0, effect.SetEff, event.Equipment, effect.None, effect.Normal, func(a ...any) any {
				subStats := a[1].(*stat.Stats)
				subStats[stat.SP_PRF] += 7
				subStats[stat.SP_ATK] += 25
				return nil
			}),
		},
		[2]string{"Aumenta la difesa di 20 e il mana di 30", "Aumenta l'attacco speciale di 25 e la perforazione speciale  di 7"},
	},
}
