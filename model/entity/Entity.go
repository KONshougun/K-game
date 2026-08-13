package entity

import (
	"gioco/game/combat/effect"
	"gioco/game/combat/event"
	"gioco/game/comm/area"
	"gioco/model"
	"gioco/model/equipment/armor"
	"gioco/model/equipment/weapon"
	"gioco/model/stat"
	"math"
	"math/rand"
)

type Controller interface {
	GetAbility(Entity) int
	GetTarget(targets []*Entity, targetCount uint8, differentTargets bool) []int
}

type Entity struct {
	Level     uint8
	MaxLvl    uint8
	Name      string //rimuovibile
	Class     []EntityClass
	Weapons   []*weapon.Weapon
	Armor     armor.Armor
	Abilities []Ability
	Passives  []Passive
	BaseStats stat.Stats
	Stats     stat.Stats

	Controller Controller

	Effects [event.EventCount][]effect.Effect
	Buffs   []effect.Buff

	//hp non in battaglia
	CurrHp float32
	//fight variables
	MaxHp float32
	MaxSt float32
	MaxMp float32
}

func NewEntity(c Controller, name string) *Entity {
	e := &Entity{
		Level:      1,
		MaxLvl:     100,
		Name:       name,
		Controller: c,
		Weapons:    make([]*weapon.Weapon, 0, 1),
	}
	e.AddBaseEvent()
	return e
}

func (e *Entity) SetBaseStats(s stat.Stats) {
	e.BaseStats = s
	e.CurrHp = s[stat.HP]
}

func (e *Entity) SetStatsOf(lvl uint8, baseStat stat.Stats) {

	if e.MaxLvl <= 1 {
		panic("MaxLvl must be greater than 1")
	}
	if lvl < 1 {
		lvl = 1
	}
	if lvl > e.MaxLvl {
		lvl = e.MaxLvl
	}

	var app stat.Stats

	// Somma dei moltiplicatori delle classi
	for _, class := range e.Class {
		EntiityClassesMultiplier[class].ForEach(func(sn stat.StatName, f float32) {
			app[stat.StatName(sn)] += f
		})
	}

	// Media dei moltiplicatori
	classCount := float32(len(e.Class))
	app.ForEach(func(sn stat.StatName, value float32) {
		app[sn] = value / classCount
	})
	// Progressione del livello
	exponent := float64(lvl-1) / float64(e.MaxLvl-1)

	// Calcolo statistiche
	app.ForEach(func(sn stat.StatName, f float32) {
		e.BaseStats[sn] = baseStat[sn] * float32(math.Pow(float64(f), exponent))
	})
}

func (e *Entity) CalculateStats() stat.Stats {
	var s stat.Stats
	e.BaseStats.ForEach(func(sn stat.StatName, f float32) {
		s[sn] = f
	})

	mainStats, subStats := e.GetEquipmentStats()
	mainStats.ForEach(func(sn stat.StatName, f float32) {
		s[sn] *= (100 + f) / 100
	})
	subStats.ForEach(func(sn stat.StatName, f float32) {
		s[sn] += f
	})

	mainStats.ForEach(func(sn stat.StatName, f float32) {
		s[sn] *= (100 + f) / 100
		s[sn] += f
	})

	e.MaxHp = s[stat.HP]
	e.MaxSt = s[stat.ST]
	e.MaxMp = s[stat.MP]
	return s
}

func (e *Entity) GetEquipmentStats() (stat.Stats, stat.Stats) {
	var mainStats, subStats stat.Stats

	if len(e.Weapons) > 0 {
		for _, w := range e.Weapons {
			w.MainStats.ForEach(func(sn stat.StatName, f float32) {
				mainStats[sn] += f
			})
			w.SubStats.ForEach(func(sn stat.StatName, f float32) {
				subStats[sn] += f
			})
			for _, eff := range weapon.WeaponStats[w.Set].Effect {
				if eff.Event == event.Equipment {
					eff.Activate(&mainStats, &subStats)
				}
			}
		}
		subStats[stat.ACC] /= float32(len(e.Weapons))
	}

	e.Armor.ForEach(func(ap armor.ArmorPiece) {
		ap.MainStats.ForEach(func(sn stat.StatName, f float32) {
			mainStats[sn] += f
		})
		ap.SubStats.ForEach(func(sn stat.StatName, f float32) {
			subStats[sn] += f
		})
	})
	effs := e.Armor.GetEffects()
	for _, eff := range effs {
		if eff.Event == event.Equipment {
			eff.Activate(&mainStats, &subStats)
		}
	}

	return mainStats, subStats
}

func (e *Entity) AddBaseEvent() {
	e.AddEffect(event.StartFight, effect.FightBaseEff, effect.None, func(a ...any) any {
		e.AbilitiesCdReset()
		e.Stats = e.CalculateStats()
		return nil
	})
	e.AddEffect(event.StartTurn, effect.FightBaseEff, effect.None, func(a ...any) any {
		e.RestoreEnergy()
		e.AbilityCdReduction()
		e.BuffCdReduction()
		return nil
	})

	e.AddEffect(event.EndTurn, effect.FightBaseEff, effect.None, func(a ...any) any {
		e.AbilityCdReduction()
		e.BuffCdReduction()
		return nil
	})
	e.AddEffect(event.CastAbility, effect.FightBaseEff, effect.None, func(a ...any) any {
		v, ok := e.eventHandler(event.CastAbility, a...)[0].(int)
		if !ok {
			panic("Ammazzati grazie")
		}
		e.Abilities[v].CastAbility(&e.Stats)
		return nil
	})
	e.AddEffect(event.EndTurn, effect.QtrEff, effect.None, func(a ...any) any {
		return e.QtrActivasion()
	})
	e.AddEffect(event.Died, effect.Revive, effect.Before, func(a ...any) any {
		resRate := float32(rand.Intn(100))
		return resRate < e.Stats[stat.RES]
	})
}
func (e *Entity) AddEffect(ev event.EventType, ef effect.IdEffect, trg effect.TriggerTiming, f func(a ...any) any) {
	e.Effects[ev] = append(e.Effects[ev], effect.NewEffect(ef, ev, trg, f))
}
func (e *Entity) TriggerEffects(ev event.EventType, trg effect.TriggerTiming, a ...any) map[effect.IdEffect]any {
	foo := make(map[effect.IdEffect]any)
	for _, ef := range e.Effects[ev] {
		if ef.Trigger == trg {
			app := ef.Activate(a...)
			if app != nil {
				foo[ef.IdEff] = app
			}
		}
	}
	return foo
}

func (e *Entity) AbilitiesCdReset() {
	for i := range e.Abilities {
		e.Abilities[i].Cooldown = 0
	}
}
func (e *Entity) RestoreEnergy() {
	e.Stats[stat.ST] += e.Stats[stat.ST_RG]
	if e.Stats[stat.ST] > e.MaxSt {
		e.Stats[stat.ST] = e.MaxSt
	}
	e.Stats[stat.MP] += e.Stats[stat.MP_RG]
	if e.Stats[stat.MP] > e.MaxMp {
		e.Stats[stat.MP] = e.MaxMp
	}
}
func (e *Entity) AbilityCdReduction() {
	for i := range e.Abilities {
		e.Abilities[i].Cooldown--
	}
}
func (e *Entity) BuffCdReduction() {
	for i := len(e.Buffs) - 1; i >= 0; i-- {
		if !e.Buffs[i].DecreaseDuration() {
			e.Buffs = append(e.Buffs[:i], e.Buffs[i+1:]...)
		}
	}
}
func (e *Entity) QtrActivasion() bool {
	return rand.Intn(100) < int(e.Stats[stat.QTR])
}

//FIGHT FUNCS

func (e *Entity) ActiveAbility(ability int, opponents any) map[int][]model.Damage {
	e.TriggerEffects(event.CastAbility, effect.None, ability)
	if atks, ok := e.Abilities[ability].execute.Activate(opponents).(map[int][]model.Damage); ok {
		return atks
	} else {
		panic("vai cazzo")
	}
}

func (e *Entity) HasHit(defenderStat stat.Stats, abilityAcc float32) bool {
	return rand.Intn(100) < int(e.Stats[stat.ACC]+abilityAcc-defenderStat[stat.EVA])
}

// 0 = NO CRIT 1 = NORMAL CRIT 2 = SUPER CRIT
func (e *Entity) HasCritted() model.DamageType {
	rate := float32(rand.Intn(100))
	if rate < e.Stats[stat.CRIT_RATE] {
		if rate*2.5 < e.Stats[stat.CRIT_RATE] {
			return model.SuperCrit
		} else {
			return model.Crit
		}
	}
	return model.Flat
}

func (e *Entity) GetCritDmg(defenderStat stat.Stats, dmg float32, critType model.DamageType) float32 {
	if defenderStat[stat.CRIT_RES] >= 100 {
		panic("COGLIONE")
	}
	switch critType {
	case model.Flat:
		return dmg
	case model.Crit:
		return dmg * ((1 + e.Stats[stat.CRIT_DMG]/100) * (1 - defenderStat[stat.CRIT_RES]/100))
	case model.SuperCrit:
		return dmg * ((1 + e.Stats[stat.CRIT_DMG]/100) * 1.5 * (1 - defenderStat[stat.CRIT_RES]/100))
	default:
		panic("get out")
	}
}

func (e *Entity) eventHandler(ev event.EventType, a ...any) []any {
	switch ev {
	case event.StartFight:
		team, ok := a[0].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		opponentTeam, ok := a[1].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		area, ok := a[2].(area.AreaId)
		if !ok {
			panic("Daje roma cazzo")
		}
		return []any{team, opponentTeam, area}
	case event.EndFight:
		team, ok := a[0].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		return []any{team}
	case event.StartTurn, event.EndTurn, event.QtrTrigged:
		team, ok := a[0].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		opponentTeam, ok := a[1].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		return []any{team, opponentTeam}

	case event.CastAbility:
		iAbility, ok := a[0].(int)
		if !ok {
			panic("Daje roma cazzo")
		}
		return []any{iAbility}
	case event.DamageDealt, event.DamageTaken:
		team, ok := a[0].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		opponentTeam, ok := a[1].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}

		actor, ok := a[2].(int)
		if !ok {
			panic("Daje roma cazzo")
		}

		dmg, ok := a[3].(float32)
		if !ok {
			panic("Daje roma cazzo")
		}
		dmgType, ok := a[4].(model.DamageType)
		if !ok {
			panic("Daje roma cazzo")
		}
		dmgDealt, ok := a[5].(bool)
		if !ok {
			panic("Daje roma cazzo")
		}
		return []any{team, opponentTeam, actor, dmg, dmgType, dmgDealt}
	case event.HealDealt, event.HealDTaken:
		team, ok := a[0].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		opponentTeam, ok := a[1].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		actor, ok := a[2].(Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		healAmount, ok := a[3].(float32)
		if !ok {
			panic("Daje roma cazzo")
		}
		healType, ok := a[4].(model.HealType)
		if !ok {
			panic("Daje roma cazzo")
		}
		return []any{team, opponentTeam, actor, healAmount, healType}
	case event.BuffApplied, event.BuffExpired:
		team, ok := a[0].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		opponentTeam, ok := a[1].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		buffer, ok := a[2].(*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		isBuff, ok := a[3].(bool)
		if !ok {
			panic("Daje roma cazzo")
		}
		return []any{team, opponentTeam, buffer, isBuff}
	case event.Died:
		team, ok := a[0].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		opponentTeam, ok := a[1].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		attacker, ok := a[2].(*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		dmgAmount, ok := a[3].(float32)
		if !ok {
			panic("Daje roma cazzo")
		}
		return []any{team, opponentTeam, attacker, dmgAmount}
	case event.EnergyRecovered, event.EnergySpent:
		team, ok := a[0].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		opponentTeam, ok := a[1].([]*Entity)
		if !ok {
			panic("Daje roma cazzo")
		}
		amount, ok := a[2].(float32)
		if !ok {
			panic("Daje roma cazzo")
		}
		return []any{team, opponentTeam, amount}
	default:
		panic("fatti i cazzi tuoi")
	}
}

/*
type DamageEvent struct {
	Source     *entity.Entity
	Target     *entity.Entity
	Amount     float32
	DamageType DamageType
	IsCritical bool
	IsDodged   bool
}*/
