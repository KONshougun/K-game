package characters

import (
	"gioco/game/combat/effect"
	"gioco/game/combat/event"
	"gioco/model"
	"gioco/model/entity"
	"gioco/model/entity/character"
	"gioco/model/equipment/armor"
	"gioco/model/equipment/weapon"
	"gioco/model/stat"
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
	k.Abilities[0].AbilityData = entity.AbilityData{
		Name:          "Normal attack",
		Unlock:        true,
		TargetCount:   1,
		StCost:        10,
		SpdCost:       500,
		AtkMultiplier: 1,
		Acc:           35,
	}

	k.Abilities[0].SetAbility(func(a ...any) any {
		if len(a) != 1 {
			panic("fk urself pt2")
		}
		opponentTeam := a[0].([]*entity.EntityData)
		t := k.GetTarget(opponentTeam, k.Abilities[0].TargetCount, k.Abilities[0].DifferentTargets)[0]

		attacks := make(map[int][]model.Damage, k.Abilities[0].TargetCount)
		if !k.HasHit(opponentTeam[t].Stats, float32(k.Abilities[0].Acc)) {
			attacks[t] = append(attacks[t], model.Damage{DmgType: model.Miss, Amount: 0})
			return attacks
		}

		var dmg float32
		dmgType := k.HasCritted()
		attacks[t] = append(attacks[t], k.GetDamage(dmg, 0, opponentTeam[t].Stats, dmgType))
		return attacks
	})

	//SECONDA ABILITà
	k.Abilities[1].AbilityData = entity.AbilityData{
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
	}
	k.Abilities[1].SetAbility(func(a ...any) any {
		if len(a) != 1 {
			panic("fk urself pt2")
		}
		opponentTeam := a[0].([]*entity.EntityData)
		t := k.GetTarget(opponentTeam, k.Abilities[1].TargetCount, k.Abilities[1].DifferentTargets)[0]

		attacks := make(map[int][]model.Damage, k.Abilities[1].TargetCount)
		if !k.HasHit(opponentTeam[t].Stats, float32(k.Abilities[1].Acc)) {
			attacks[t] = append(attacks[t], model.Damage{DmgType: model.Miss, Amount: 0})
			return attacks
		}

		dmg := k.Stats[stat.ATK]
		dmgType := k.HasCritted()
		attacks[t] = append(attacks[t], k.GetDamage(dmg, 0, opponentTeam[t].Stats, dmgType))

		//Applico i debuff al target
		opponentTeam[t].AddBuff(effect.Buff{
			Count:  255,
			IsBuff: false,
			GetBuff: func(a ...any) any {
				opponentTeam[t].Stats[stat.DEF] *= 0.95
				opponentTeam[t].Stats[stat.SP_DEF] *= 0.95
				return nil
			},
			RemoveBuff: func(a ...any) any {
				opponentTeam[t].Stats[stat.DEF] *= 1.05
				opponentTeam[t].Stats[stat.SP_DEF] *= 1.05
				return nil
			},
		})
		return attacks
	})

	//Terza ABILITà
	k.Abilities[2].AbilityData = entity.AbilityData{
		Name:          "Controllo del berserk",
		Description:   "Entra in uno stato di incoscienza per 3 turni",
		CooldownCost:  5,
		StCost:        25,
		StCostPercent: 5,
		MpCost:        30,
		MpCostPercent: 10,
		SpdCost:       375,
	}
	k.Abilities[2].SetAbility(func(a ...any) any {
		k.AddBuff(effect.Buff{
			Count:  3,
			IsBuff: true,
			GetBuff: func(a ...any) any {
				k.berserkMode = true
				k.Stats[stat.HP] *= 0.90

				k.Stats[stat.ATK] *= 1.20
				k.Stats[stat.SP_ATK] *= 1.20
				k.Stats[stat.CRIT_DMG] *= 1.20
				k.Stats[stat.ATK_SPD] *= 1.10

				k.Stats[stat.DEF] *= 0.85
				k.Stats[stat.SP_DEF] *= 0.85
				return nil
			},
			RemoveBuff: func(a ...any) any {
				k.berserkMode = false
				k.Stats[stat.ATK] *= 0.80
				k.Stats[stat.SP_ATK] *= 0.80
				k.Stats[stat.CRIT_DMG] *= 0.80
				k.Stats[stat.ATK_SPD] *= 0.90

				k.Stats[stat.DEF] *= 1.15
				k.Stats[stat.SP_DEF] *= 1.15

				k.Stats[stat.ST] *= 0.70
				return nil
			},
		})
		return map[int][]model.Damage{}
	})

	//Quarta abilità
	k.Abilities[3].AbilityData = entity.AbilityData{
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
	}
	k.Abilities[3].SetAbility(func(a ...any) any {

		if len(a) != 1 {
			panic("fk urself")
		}
		opponentTeam, ok := a[0].([]*entity.EntityData)
		if !ok {
			panic("fk me")
		}
		target := k.GetTarget(opponentTeam, k.Abilities[3].TargetCount, k.Abilities[3].DifferentTargets)[0]

		attacks := make(map[int][]model.Damage, 3)
		for i := target - 1; i <= target+1; i++ {

			if i < 0 || i >= len(opponentTeam) {
				continue
			}
			if !k.HasHit(opponentTeam[i].Stats, float32(k.Abilities[3].Acc)) {
				attacks[i] = append(attacks[i], model.Damage{DmgType: model.Miss, Amount: 0})
				continue
			}

			dmgType := k.HasCritted()
			spDmg := k.Stats[stat.SP_ATK] * float32(k.Abilities[3].SpAtkMultiplier)
			var dmg float32 = 0
			if i == target {
				dmg = k.Stats[stat.ATK] * float32(k.Abilities[3].AtkMultiplier)
			}
			dmg, spDmg = k.GetCritDmg(opponentTeam[i].Stats, dmg, spDmg, dmgType)

			attacks[i] = append(attacks[i], k.GetDamage(dmg, spDmg, opponentTeam[i].Stats, dmgType))
		}
		return attacks
	})
}
func (k *Knight) setPassives() {
	k.Passives = append(k.Passives, entity.NewPassive(false, "Aumenta attacco subendo danni", event.DamageTaken, effect.After, effect.Normal, func(a ...any) any {
		return nil
	}), entity.NewPassive(false, "Si cura in base al danno inflitto", event.DamageDealt, effect.After, effect.Normal, func(a ...any) any {
		return nil
	}))
}

func (k *Knight) AwakenAbility(index int) {
	switch index {
	case 1:
		k.GetCharacter().AbilityAwakened[1] = true
		k.Abilities[1].AbilityData = entity.AbilityData{
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
		}
		k.Abilities[1].SetAbility(func(a ...any) any {
			if len(a) != 1 {
				panic("fk urself pt2")
			}
			var attacks map[int][]model.Damage
			opponentTeam := a[0].([]*entity.EntityData)
			if !k.berserkMode {
				t := k.GetTarget(opponentTeam, k.Abilities[1].TargetCount, k.Abilities[1].DifferentTargets)[0]
				attacks := make(map[int][]model.Damage, k.Abilities[1].TargetCount)
				if !k.HasHit(opponentTeam[t].Stats, float32(k.Abilities[1].Acc)) {
					attacks[t] = append(attacks[t], model.Damage{DmgType: model.Miss, Amount: 0})
					return attacks
				}

				dmgType := k.HasCritted()
				//PSY DMG
				dmg := k.Stats[stat.ATK] * k.Abilities[1].AtkMultiplier
				dmg, _ = k.GetCritDmg(opponentTeam[t].Stats, dmg, 0, dmgType)
				attacks[t] = append(attacks[t], k.GetDamage(dmg, 0, opponentTeam[t].Stats, dmgType))

				var appIdEff uint16
				//Applico i debuff al target
				opponentTeam[t].AddBuff(
					effect.NewBuff(0, 255, false, true, func(a ...any) any {
						opponentTeam[t].Stats[stat.DEF] *= 0.95
						opponentTeam[t].Stats[stat.SP_DEF] *= 0.95
						appIdEff = opponentTeam[t].AddEffect(event.HealDTaken, effect.BuffEff, effect.Before, effect.Normal, func(a ...any) any {
							opponentTeam[t].EventHandler(event.HealDTaken, a...)
							healAmount := a[3].(float32)
							healType := a[4].(model.HealType)

							if healType == model.Heal {
								healAmount *= 0.90
							}
							return []any{healAmount, healType}
						})
						return nil
					}, func(a ...any) any {
						opponentTeam[t].Stats[stat.DEF] *= 1.05
						opponentTeam[t].Stats[stat.SP_DEF] *= 1.05

						opponentTeam[t].RemoveEffect(event.HealDTaken, appIdEff)
						return nil
					}))
			}
			return attacks
		})
	case 2:
		k.GetCharacter().AbilityAwakened[2] = true
		k.Abilities[2].AbilityData = entity.AbilityData{
			Name:          "Controllo del berserk",
			Description:   "Entra in uno stato di incoscienza per 4 turni",
			Unlock:        true,
			CooldownCost:  5,
			StCost:        0,
			StCostPercent: 100,
			MpCost:        40,
			MpCostPercent: 10,
			SpdCost:       300,
		}
		k.Abilities[2].SetAbility(func(a ...any) any {
			appStCost := make([]float32, len(k.Abilities))
			appStCostPer := make([]uint8, len(k.Abilities))
			k.AddBuff(effect.Buff{
				Count:  4,
				IsBuff: true,
				GetBuff: func(a ...any) any {
					k.berserkMode = true
					k.Stats[stat.HP] *= 0.90

					k.Stats[stat.ATK] *= 1.25
					k.Stats[stat.SP_ATK] *= 1.25
					k.Stats[stat.CRIT_DMG] *= 1.20
					k.Stats[stat.ATK_SPD] *= 1.15

					k.Stats[stat.DEF] *= 0.90
					k.Stats[stat.SP_DEF] *= 0.90

					for i, a := range k.Abilities {
						appStCost[i] = float32(a.StCost)
						appStCostPer[i] = uint8(k.Stats[stat.ST]/100) * a.StCostPercent
						k.Abilities[i].StCost = 0
						k.Abilities[i].StCostPercent = 0
					}
					return nil
				},
				RemoveBuff: func(a ...any) any {
					k.berserkMode = false
					k.Stats[stat.ATK] *= 0.75
					k.Stats[stat.SP_ATK] *= 0.75
					k.Stats[stat.CRIT_DMG] *= 0.80
					k.Stats[stat.ATK_SPD] *= 0.85

					k.Stats[stat.DEF] *= 1.10
					k.Stats[stat.SP_DEF] *= 1.10

					for i := range k.Abilities {
						k.Abilities[i].StCost = appStCost[i]
						k.Abilities[i].StCostPercent = appStCostPer[i]
					}
					return nil
				},
			})
			return map[int][]model.Damage{}
		})
	case 3:
		k.GetCharacter().AbilityAwakened[3] = true
		k.Abilities[3].AbilityData = entity.AbilityData{
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
		}
		k.Abilities[3].SetAbility(func(a ...any) any {
			if len(a) != 1 {
				panic("fk urself")
			}
			opponentTeam, ok := a[0].([]*entity.EntityData)
			if !ok {
				panic("fk me")
			}
			target := k.GetTarget(opponentTeam, k.Abilities[3].TargetCount, k.Abilities[3].DifferentTargets)[0]

			attacks := make(map[int][]model.Damage, len(opponentTeam))
			for i := range target {
				if !k.HasHit(opponentTeam[i].Stats, float32(k.Abilities[3].Acc)) {
					attacks[i] = append(attacks[i], model.Damage{DmgType: model.Miss, Amount: 0})
					continue
				}

				dmg := k.Stats[stat.ATK] * float32(k.Abilities[3].AtkMultiplier)
				spDmg := k.Stats[stat.SP_ATK]*float32(k.Abilities[3].SpAtkMultiplier) + dmg
				dmgType := k.HasCritted()
				dmg, spDmg = k.GetCritDmg(opponentTeam[i].Stats, dmg, spDmg, dmgType)

				if i == target {
					if k.berserkMode {
						attacks[i] = append(attacks[i], k.GetDamage(dmg, spDmg, opponentTeam[i].Stats, dmgType))
						continue
					}
					//PSY DMG
					def := opponentTeam[i].Stats[stat.DEF] - k.Stats[stat.PRF]
					if def > 0 {
						dmg -= def
						if dmg < 0 {
							dmg = 0
						}
					} else {
						dmg -= def * 1.25
					}
				}

				dmg = 0
				attacks[i] = append(attacks[i], k.GetDamage(dmg, spDmg, opponentTeam[i].Stats, dmgType))
			}
			return attacks
		})
	default:
		panic("get out of here")
	}
}
