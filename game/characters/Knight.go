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
	*character.Character
}

func NewKnight(lvl uint8, w *weapon.Weapon, a *armor.Armor) *Knight {
	k := &Knight{
		Character: character.NewCharacter(character.Knight, a, w, weapon.Sword),
	}
	k.setAbilities()
	k.setPassives()
	k.SetBaseStats(character.CharactersBaseStats[k.CharacterName])
	k.Class = append(k.Class, entity.Brawler)
	k.SetLevel(lvl)
	return k
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

	//Setto l'abilità
	k.Abilities[0].SetAbility(func(a ...any) any {
		if len(a) != 1 {
			panic("fk urself pt2")
		}
		opponentTeam, ok := a[0].([]*entity.Entity)

		if !ok {
			panic("fk me pt2")
		}
		targets := k.Controller.GetTarget(opponentTeam, k.Abilities[0].TargetCount, k.Abilities[0].DifferentTargets)

		attacks := make(map[int][]model.Damage, len(targets))
		for _, t := range targets {
			if !k.HasHit(opponentTeam[t].Stats, float32(k.Abilities[0].Acc)) {
				attacks[t] = append(attacks[t], model.Damage{DmgType: model.Miss, Amount: 0})
				continue
			}

			var dmg float32
			dmgType := k.HasCritted()
			//PSY DMG
			dmg = k.Stats[stat.ATK]
			dmg = k.GetCritDmg(opponentTeam[t].Stats, dmg, dmgType)

			def := opponentTeam[t].Stats[stat.DEF] - k.Stats[stat.PRF]
			if def > 0 {
				dmg -= def
				if dmg < 0 {
					dmg = 0
				}
			} else {
				dmg -= def * 1.25
			}
			attacks[t] = append(attacks[t], model.Damage{DmgType: dmgType, Amount: dmg})
		}
		return attacks
	})
}
func (k *Knight) setPassives() {
	k.Passives = append(k.Passives, entity.NewPassive(false, "Aumenta attacco subendo danni", event.DamageTaken, effect.After, func(a ...any) any {
		return nil
	}), entity.NewPassive(false, "Si cura in base al danno inflitto", event.DamageDealt, effect.After, func(a ...any) any {
		return nil
	}))
}

/*
		// Second ability
		abilities[1] = new Ability(abilityData[1]) {

			@Override
			public void attack(int target) {
				super.attack(target);

				setCooldown(2);

				statistiche.multiply(STATS_NAMES.DEF, 1 + (0.1f * Account.enemys.size()), true);
				statistiche.multiply(STATS_NAMES.SP_DEF, 1 + (0.1f * Account.enemys.size()), false);
				// buff(STATS_NAMES.DEF, statistiche.get(STATS_NAMES.DEF) * 0.05f *
				// Account.enemys.size(), true);

				for (int i = Account.enemys.size() - 1; i >= 0; i--) {

					if (hasHit(Account.enemys.get(i))) {

						int crit = hasCritted();

						if (crit == 2) {
							Common.print(nome + " HA FATTO UN SUPER CRITICO!!!", COLORS.brightBlue);
						} else if (crit == 1) {
							Common.print(nome + " HA FATTO UN DANNO CRITICO!!", COLORS.brightBlue);
						}

						float danno = 0, spDanno = 0;
						danno = statistiche.get(STATS_NAMES.ATK) * atkMultiplier;
						danno += getCrit(danno, crit, i);

						float def = Account.enemys.get(i).statistiche.get(STATS_NAMES.DEF)
								- statistiche.get(STATS_NAMES.PRF);
						danno -= (def >= 0 ? def : def * 1.25f);
						if (danno < 0) {
							danno = 0;
						}

						attackEnemy(i, danno + spDanno);

					} else {
						Common.print(Account.enemys.get(i).nome + " ha schivato l'attacco", COLORS.brightRed);
					}
				}
			}

		};

		// Third ability
		abilities[2] = new Ability(abilityData[2]) {

			@Override
			public void attack(int target) {
				super.attack(target);

				setCooldown(3);

				damaged(Knight.this, statistiche.multiply(STATS_NAMES.HP, 0.8f, false));

				statistiche.multiply(STATS_NAMES.ATK, 1.2f, false);
				statistiche.multiply(STATS_NAMES.SP_ATK, 1.2f, false);
				statistiche.multiply(STATS_NAMES.ATK_SPD, 1.1f, false);

				statistiche.divide(STATS_NAMES.DEF, 1.1f, false);
				statistiche.divide(STATS_NAMES.SP_DEF, 1.1f, false);

				effects.add(new Effect(3) {
					@Override
					public void run() {
						statistiche.divide(STATS_NAMES.ATK, 1.2f, false);
						statistiche.divide(STATS_NAMES.SP_ATK, 1.2f, false);
						statistiche.divide(STATS_NAMES.ATK_SPD, 1.1f, false);

						statistiche.multiply(STATS_NAMES.DEF, 1.1f, false);
						statistiche.multiply(STATS_NAMES.SP_DEF, 1.1f, false);
					}
				});
			}

		};

		// Fourth ability
		abilities[3] = new Ability(abilityData[3]) {

			@Override
			public void attack(int target) {
				super.attack(target);

				setCooldown(4);

				statistiche.multiply(STATS_NAMES.ATK, 1.5f, false);
				statistiche.multiply(STATS_NAMES.CRIT_DMG, 1.2f, false);
				statistiche.multiply(STATS_NAMES.CRIT_RATE, 1.1f, false);

				// INFLIGGE DANNI AL TARGET E AGLI AVVERSARI ADIACENTI
				for (int i = target + 1; i >= target - 1; i--) {

					Common.debugRequest("Enemy target index: "+ i + " Enemy size: "+Account.enemys.size());

					for(int j=0; j< Account.enemys.size();j++) {
						Common.debugRequest("Nemico: "+Account.enemys.get(j).nome+" index:"+j);
					}
					if (i < Account.enemys.size() && i >= 0) {

						if (!hasHit(Account.enemys.get(i))) {
							System.out.println(
									COLORS.brightRed + Account.enemys.get(i).nome + " ha schivato l'attacco" + COLORS.reset);
						} else {

							int crit = hasCritted();
							if (crit == 2) {
								Common.print(nome + " HA FATTO UN SUPER CRITICO!!!", COLORS.brightBlue);
							} else if (crit == 1) {
								Common.print(nome + " HA FATTO UN DANNO CRITICO!!", COLORS.brightBlue);
							}
							// NEMICO ADIACENTE
							if (i != target) {

								float spDanno = statistiche.get(STATS_NAMES.SP_ATK) * spAtkMultiplier;
								spDanno += getCrit(spDanno, crit, i);

								float spDef = Account.enemys.get(i).statistiche.get(STATS_NAMES.SP_DEF)
										- statistiche.get(STATS_NAMES.SP_PRF);
								spDanno -= (spDef >= 0 ? spDef : spDef * 1.25f);
								if (spDanno < 0) {
									spDanno = 0;
								}

								attackEnemy(i, spDanno);

							// TARGET
							} else {

								float danno = statistiche.get(STATS_NAMES.ATK) * atkMultiplier;
								danno += getCrit(danno, crit, i);

								float def = Account.enemys.get(i).statistiche.get(STATS_NAMES.DEF)
										- statistiche.get(STATS_NAMES.PRF);
								danno -= (def >= 0 ? def : def * 1.25f);
								if (danno < 0) {
									danno = 0;
								}

								attackEnemy(i, danno);
							}
						}
					}
				}

				statistiche.divide(STATS_NAMES.ATK, 1.5f, false);
				statistiche.divide(STATS_NAMES.CRIT_DMG, 1.2f, false);
				statistiche.divide(STATS_NAMES.CRIT_RATE, 1.1f, false);
			}
		};
	}

}
*/
