package armor

import (
	"github.com/KONshougun/K-game/game/combat/effect"
	"github.com/KONshougun/K-game/game/combat/event"
	"github.com/KONshougun/K-game/model/stat"
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

type ArmorEffect struct {
	Effect         effect.Effect
	Description    string
	RequiredPieces uint8
}

type ArmorData struct {
	effects []ArmorEffect
}

var ArmorEffects = [SetCount]ArmorData{
	Ferro: {
		[]ArmorEffect{
			{
				RequiredPieces: 2,
				Description:    "Aumenta l'attacco di 20",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.Atk] += 20
					return evData
				}),
			},
			{
				RequiredPieces: 5,
				Description:    "Aumenta gli hp di 20 la difesa di 16",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.Hp] += 20
					evData.GetData()[0].(*stat.Stats)[stat.Def] += 16
					return evData
				}),
			},
		},
	},
	Lupo: {
		[]ArmorEffect{
			{
				RequiredPieces: 2,
				Description:    "Aumenta la stamina di 25",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.St] += 25
					return evData
				}),
			}, {
				RequiredPieces: 5,
				Description:    "Aumenta l'attacco di 15 e la velocità di 20",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.Atk] += 15
					evData.GetData()[0].(*stat.Stats)[stat.Spd] += 20
					return evData
				}),
			},
		},
	},
	Ossa: {
		[]ArmorEffect{
			{
				RequiredPieces: 2,
				Description:    "Aumenta la difesa di 10",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.Def] += 10
					return evData
				}),
			}, {
				RequiredPieces: 5,
				Description:    "Aumenta gli hp di 25 e l'attacco di 20",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.Hp] += 25
					evData.GetData()[0].(*stat.Stats)[stat.Atk] += 20
					return evData
				}),
			},
		},
	},
	Scaglie: {
		[]ArmorEffect{
			{
				RequiredPieces: 2,
				Description:    "Aumenta il mana di 25",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.Mp] += 25
					return evData
				}),
			}, {
				RequiredPieces: 5,
				Description:    "Aumenta l'attacco speciale di 20 e la difesa di 20",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.SpAtk] += 20
					evData.GetData()[0].(*stat.Stats)[stat.Def] += 20
					return evData
				}),
			},
		},
	},
	Ketos: {
		[]ArmorEffect{
			{
				RequiredPieces: 2,
				Description:    "Aumenta gli hp di 50 e la velocità di 100",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.Hp] += 50
					evData.GetData()[0].(*stat.Stats)[stat.Spd] += 100
					return evData
				}),
			}, {
				RequiredPieces: 5,
				Description:    "Aumenta la difesa e la difesa speciale di 25",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.Def] += 25
					evData.GetData()[0].(*stat.Stats)[stat.SpDef] += 25
					return evData
				}),
			},
		},
	},
	Fenice: {
		[]ArmorEffect{
			{
				RequiredPieces: 2,
				Description:    "Aumenta la difesa speciale di 20 e il mana di 30",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.SpDef] += 20
					evData.GetData()[0].(*stat.Stats)[stat.Mp] += 30
					return evData
				}),
			}, {
				RequiredPieces: 5,
				Description:    "Aumenta l'attacco speciale di 25 e la perforazione speciale  di 7",
				Effect: effect.NewEffect(0, event.SetEntityData, effect.None, effect.Normal, func(evData event.EventData) event.EventData {
					evData.GetData()[0].(*stat.Stats)[stat.SpPrf] += 7
					evData.GetData()[0].(*stat.Stats)[stat.SpAtk] += 25
					return evData
				}),
			},
		},
	},
}

func getEffects(set ArmorSet, count uint8) []effect.Effect {
	effs := make([]effect.Effect, 0, len(ArmorEffects[set].effects))
	for _, eff := range ArmorEffects[set].effects {
		if count >= eff.RequiredPieces {
			effs = append(effs, eff.Effect)
		}
	}
	return effs
}
