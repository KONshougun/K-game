package characters

import (
	"github.com/KONshougun/K-game/game/combat/effect"
	"github.com/KONshougun/K-game/game/combat/event"
	"github.com/KONshougun/K-game/model"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/KONshougun/K-game/model/entity/character"
	"github.com/KONshougun/K-game/model/equipment/armor"
	"github.com/KONshougun/K-game/model/equipment/weapon"
	"github.com/KONshougun/K-game/model/stat"
)

type Knight struct {
	*character.CharacterData

	berserkMode bool
}

func NewKnight(lvl uint8, w *weapon.Weapon, a *armor.Armor) *Knight {
	k := &Knight{
		CharacterData: character.NewCharacter(character.Knight, a, w, weapon.Sword),
	}
	k.setAbilities()
	k.setPassives()
	k.SetBaseStats(character.CharactersBaseStats[k.CharacterName])
	k.Class = append(k.Class, entity.Brawler)
	k.SetLevel(lvl)
	return k
}

func (k *Knight) GetCharacter() *character.CharacterData {
	return k.CharacterData
}
func (k *Knight) setAbilities() {
	k.Abilities[0] = entity.NewAbility(entity.AbilityData{
		Name:          "Normal attack",
		Unlock:        true,
		TargetCount:   1,
		StCost:        10,
		SpdCost:       500,
		AtkMultiplier: 1,
		Acc:           35,
	}, func(opponents []*entity.EntityData) map[int][]model.Damage {
		t := k.GetTarget(opponents, k.Abilities[0].TargetCount, k.Abilities[0].DifferentTargets, false)[0]

		attacks := make(map[int][]model.Damage, k.Abilities[0].TargetCount)
		if !k.HasHit(opponents[t].Stats, float64(k.Abilities[0].Acc)) {
			attacks[t] = append(attacks[t], model.Damage{DmgType: model.Miss, Amount: 0})
			return attacks
		}

		dmg := k.Stats[stat.Atk] * k.Abilities[0].AtkMultiplier
		dmgType := k.HasCritted()
		attacks[t] = append(attacks[t], k.CalculateDmg(dmg, 0, opponents[t].Stats, dmgType))
		return attacks
	})
	//SECONDA ABILITà
	k.Abilities[1] = entity.NewAbility(entity.AbilityData{
		Name:          "Fendente sfiancante",
		Description:   "Colpisce un avversario con un fendente e gli abbassa la difesa e difesa speciale",
		Unlock:        true,
		CooldownCost:  2,
		TargetCount:   1,
		StCost:        35,
		StCostPercent: 3,
		SpdCost:       550,
		AtkMultiplier: 1,
		Acc:           35,
	}, func(opponents []*entity.EntityData) map[int][]model.Damage {
		t := k.GetTarget(opponents, k.Abilities[1].TargetCount, k.Abilities[1].DifferentTargets, false)[0]

		attacks := make(map[int][]model.Damage, k.Abilities[1].TargetCount)
		if !k.HasHit(opponents[t].Stats, float64(k.Abilities[1].Acc)) {
			attacks[t] = append(attacks[t], model.Damage{DmgType: model.Miss, Amount: 0})
			return attacks
		}

		dmg := k.Stats[stat.Atk] * k.Abilities[1].AtkMultiplier
		dmgType := k.HasCritted()
		attacks[t] = append(attacks[t], k.CalculateDmg(dmg, 0, opponents[t].Stats, dmgType))

		//Applico i debuff al target
		opponents[t].AddBuff(effect.Buff{
			Count:  255,
			IsBuff: false,
			GetBuff: func() {
				opponents[t].Stats[stat.Def] *= 0.95
				opponents[t].Stats[stat.SpDef] *= 0.95
			},
			RemoveBuff: func() {
				opponents[t].Stats[stat.Def] *= 1.05
				opponents[t].Stats[stat.SpDef] *= 1.05
			},
		})
		return attacks
	})
	//Terza ABILITà
	k.Abilities[2] = entity.NewAbility(entity.AbilityData{
		Name:          "Controllo del berserk",
		Description:   "Entra in uno stato di incoscienza per 3 turni",
		CooldownCost:  5,
		StCost:        25,
		StCostPercent: 5,
		MpCost:        30,
		MpCostPercent: 10,
		SpdCost:       375,
	}, func(opponents []*entity.EntityData) map[int][]model.Damage {
		k.AddBuff(effect.Buff{
			Count:  3,
			IsBuff: true,
			GetBuff: func() {
				k.berserkMode = true
				k.Stats[stat.Hp] *= 0.90

				k.Stats[stat.Atk] *= 1.20
				k.Stats[stat.SpAtk] *= 1.20
				k.Stats[stat.CritDmg] *= 1.20
				k.Stats[stat.AtkSpd] *= 1.10

				k.Stats[stat.Def] *= 0.85
				k.Stats[stat.SpDef] *= 0.85
			},
			RemoveBuff: func() {
				k.berserkMode = false
				k.Stats[stat.Atk] *= 0.80
				k.Stats[stat.SpAtk] *= 0.80
				k.Stats[stat.CritDmg] *= 0.80
				k.Stats[stat.AtkSpd] *= 0.90

				k.Stats[stat.Def] *= 1.15
				k.Stats[stat.SpDef] *= 1.15

				k.Stats[stat.St] *= 0.70
			},
		})
		return map[int][]model.Damage{}
	})
	//Quarta abilità
	k.Abilities[3] = entity.NewAbility(entity.AbilityData{
		Name:          "Excalibur",
		Description:   "Colpisce con un fendente carico di magia",
		CooldownCost:  6,
		StCost:        30,
		StCostPercent: 25,
		MpCost:        30,
		MpCostPercent: 30,
		SpdCost:       600,

		TargetCount:     1,
		AtkMultiplier:   1.30,
		SpAtkMultiplier: 1,
		Acc:             40,
	}, func(opponents []*entity.EntityData) map[int][]model.Damage {

		target := k.GetTarget(opponents, k.Abilities[3].TargetCount, k.Abilities[3].DifferentTargets, false)[0]

		attacks := make(map[int][]model.Damage, 3)
		for i := target - 1; i <= target+1; i++ {
			if i < 0 || i >= len(opponents) {
				continue
			}
			if !k.HasHit(opponents[i].Stats, float64(k.Abilities[3].Acc)) {
				attacks[i] = append(attacks[i], model.Damage{DmgType: model.Miss, Amount: 0})
				continue
			}

			dmgType := k.HasCritted()
			spDmg := k.Stats[stat.SpAtk] * float64(k.Abilities[3].SpAtkMultiplier)
			var dmg float64 = 0
			if i == target {
				dmg = k.Stats[stat.Atk] * float64(k.Abilities[3].AtkMultiplier)
			}

			attacks[i] = append(attacks[i], k.CalculateDmg(dmg, spDmg, opponents[i].Stats, dmgType))
		}
		return attacks
	})
}

// TODO
func (k *Knight) setPassives() {
	k.Passives = append(k.Passives, entity.NewPassive(false, "Aumenta attacco subendo danni", event.DamageTaken, effect.After, effect.Normal, func(evt event.EventData) event.EventData {

		return evt
	}), entity.NewPassive(false, "Si cura in base al danno inflitto", event.DamageDealt, effect.After, effect.Normal, func(evt event.EventData) event.EventData {

		return evt
	}))
}

func (k *Knight) AwakenAbility(index int) {
	switch index {
	case 1:
		k.GetCharacter().AbilityAwakened[1] = true
		k.Abilities[1] = entity.NewAbility(entity.AbilityData{
			Name:          "Fendente sfiancante",
			Description:   "Colpisce un avversario con un fendente e gli abbassa difesa e difesa speciale, inoltre riduce le cure ricevute per 3 turni",
			Unlock:        true,
			CooldownCost:  3,
			TargetCount:   1,
			StCost:        35,
			StCostPercent: 10,
			SpdCost:       600,
			AtkMultiplier: 1,
			Acc:           35,
		}, func(opponents []*entity.EntityData) map[int][]model.Damage {

			targets := make([]int, 0, 3)
			if !k.berserkMode {
				targets = k.GetTarget(opponents, k.Abilities[1].TargetCount, k.Abilities[1].DifferentTargets, false)
			} else {
				targets = k.GetTarget(opponents, 3, k.Abilities[1].DifferentTargets, true)
			}
			attacks := make(map[int][]model.Damage, len(targets))
			for _, t := range targets {
				if !k.HasHit(opponents[t].Stats, float64(k.Abilities[1].Acc)) {
					attacks[t] = append(attacks[t], model.Damage{DmgType: model.Miss, Amount: 0})
					return attacks
				}

				dmgType := k.HasCritted()
				//PSY DMG
				dmg := k.Stats[stat.Atk] * k.Abilities[1].AtkMultiplier
				attacks[t] = append(attacks[t], k.CalculateDmg(dmg, 0, opponents[t].Stats, dmgType))

				var appIdEff uint16
				//Applico i debuff al target
				opponents[t].AddBuff(
					effect.NewBuff(0, 255, false, true, func() {
						opponents[t].Stats[stat.Def] *= 0.95
						opponents[t].Stats[stat.SpDef] *= 0.95
						appIdEff = opponents[t].AddEffect(event.HealTaken, effect.Before, effect.Normal, func(evt event.EventData) event.EventData {
							if *(evt.GetData()[5].(*model.HealType)) == model.Heal {
								*(evt.GetData()[4].(*float64)) *= 0.90
							}
							return evt
						})
					}, func() {
						opponents[t].Stats[stat.Def] *= 1.05
						opponents[t].Stats[stat.SpDef] *= 1.05

						opponents[t].RemoveEffect(event.HealTaken, appIdEff)
					}))
			}
			return attacks
		})
	case 2:
		k.GetCharacter().AbilityAwakened[2] = true
		k.Abilities[2] = entity.NewAbility(entity.AbilityData{
			Name:          "Controllo del berserk",
			Description:   "Entra in uno stato di incoscienza per 4 turni",
			Unlock:        true,
			CooldownCost:  5,
			StCost:        0,
			StCostPercent: 100,
			MpCost:        40,
			MpCostPercent: 10,
			SpdCost:       300,
		}, func(opponents []*entity.EntityData) map[int][]model.Damage {

			appStCost := make([]float64, len(k.Abilities))
			appStCostPer := make([]uint8, len(k.Abilities))
			k.AddBuff(effect.Buff{
				Count:  4,
				IsBuff: true,
				GetBuff: func( /*evt event.EventData*/ ) /*event.EventData*/ {
					k.berserkMode = true
					k.Stats[stat.Hp] *= 0.90

					k.Stats[stat.Atk] *= 1.25
					k.Stats[stat.SpAtk] *= 1.25
					k.Stats[stat.CritDmg] *= 1.20
					k.Stats[stat.AtkSpd] *= 1.15

					k.Stats[stat.Def] *= 0.90
					k.Stats[stat.SpDef] *= 0.90

					for i, a := range k.Abilities {
						appStCost[i] = float64(a.StCost)
						appStCostPer[i] = uint8(k.Stats[stat.St]/100) * a.StCostPercent
						k.Abilities[i].StCost = 0
						k.Abilities[i].StCostPercent = 0
					}
				},
				RemoveBuff: func( /*evt event.EventData*/ ) /*event.EventData*/ {
					k.berserkMode = false
					k.Stats[stat.Atk] *= 0.75
					k.Stats[stat.SpAtk] *= 0.75
					k.Stats[stat.CritDmg] *= 0.80
					k.Stats[stat.AtkSpd] *= 0.85

					k.Stats[stat.Def] *= 1.10
					k.Stats[stat.SpDef] *= 1.10

					for i := range k.Abilities {
						k.Abilities[i].StCost = appStCost[i]
						k.Abilities[i].StCostPercent = appStCostPer[i]
					}
				},
			})
			return map[int][]model.Damage{}
		})
	case 3:
		k.GetCharacter().AbilityAwakened[3] = true
		k.Abilities[3] = entity.NewAbility(entity.AbilityData{
			Name:          "Excalibur",
			Description:   "Colpisce con un fendente carico di magia",
			Unlock:        true,
			CooldownCost:  6,
			StCost:        30,
			StCostPercent: 30,
			MpCost:        50,
			MpCostPercent: 30,
			SpdCost:       700,

			TargetCount:     1,
			AtkMultiplier:   1.30,
			SpAtkMultiplier: 1,
			Acc:             40,
		}, func(opponents []*entity.EntityData) map[int][]model.Damage {
			target := k.GetTarget(opponents, k.Abilities[3].TargetCount, k.Abilities[3].DifferentTargets, false)[0]

			attacks := make(map[int][]model.Damage, len(opponents))
			dmg := k.Stats[stat.Atk] * float64(k.Abilities[3].AtkMultiplier)
			spDmg := k.Stats[stat.SpAtk]*float64(k.Abilities[3].SpAtkMultiplier) + dmg

			for i := range len(opponents) {
				if !k.HasHit(opponents[i].Stats, float64(k.Abilities[3].Acc)) {
					attacks[i] = append(attacks[i], model.Damage{DmgType: model.Miss, Amount: 0})
					continue
				}

				dmgType := k.HasCritted()
				if i == target {
					if k.berserkMode {
						attacks[i] = append(attacks[i], k.CalculateDmg(dmg, spDmg, opponents[i].Stats, model.Pure))
					} else {
						attacks[i] = append(attacks[i], k.CalculateDmg(dmg, spDmg, opponents[i].Stats, dmgType))
					}
				} else {
					attacks[i] = append(attacks[i], k.CalculateDmg(0, spDmg, opponents[i].Stats, dmgType))
				}
			}
			return attacks
		})
	default:
		panic("get out of here")
	}
}
