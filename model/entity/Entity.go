package entity

import (
	"fmt"
	"gioco/game/combat/effect"
	"gioco/game/combat/event"
	"gioco/game/comm/area"
	"gioco/model"
	"gioco/model/equipment/armor"
	"gioco/model/equipment/weapon"
	"gioco/model/stat"
	"math"
	"math/rand"
	"sort"
)

type Entity interface {
	GetAbility() int
	GetTarget(targets []*EntityData, targetCount uint8, differentTargets bool) []int
	GetEntity() *EntityData
}

type EntityData struct {
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

	Effects [event.EventCount][]effect.Effect
	Buffs   []effect.Buff

	//hp non in battaglia
	CurrHp float32
	//fight variables
	MaxHp float32
	MaxSt float32
	MaxMp float32
}

func NewEntity(name string) *EntityData {
	e := &EntityData{
		Level:   1,
		MaxLvl:  100,
		Name:    name,
		Weapons: make([]*weapon.Weapon, 0, 1),
	}
	e.AddBaseEvent()
	return e
}

func (e *EntityData) SetBaseStats(s stat.Stats) {
	e.BaseStats = s
	e.CurrHp = s[stat.HP]
}

func (e *EntityData) SetStatsOf(lvl uint8, baseStat stat.Stats) {

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

func (e *EntityData) CalculateStats() stat.Stats {
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

func (e *EntityData) GetEquipmentStats() (stat.Stats, stat.Stats) {
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

func (e *EntityData) GetDamage(dmg, spDmg float32, enemyStat stat.Stats, dmgT model.DamageType) model.Damage {
	switch dmgT {
	case model.Pure:
		return model.Damage{
			DmgType: dmgT,
			Amount:  dmg + spDmg,
		}
	case model.Miss:
		return model.Damage{DmgType: dmgT, Amount: 0}
	}
	dmg, spDmg = e.GetCritDmg(enemyStat, dmg, spDmg, dmgT)
	def := enemyStat[stat.DEF] - e.Stats[stat.PRF]
	if def > 0 {
		dmg -= def
		if dmg < 0 {
			dmg = 0
		}
	} else {
		dmg -= def * 1.25
	}
	spDef := enemyStat[stat.SP_DEF] - e.Stats[stat.SP_PRF]
	if spDef > 0 {
		spDmg -= spDef
		if spDmg < 0 {
			spDmg = 0
		}
	} else {
		spDmg -= spDef * 1.25
	}
	return model.Damage{
		DmgType: dmgT,
		Amount:  dmg + spDmg,
	}
}

func (e *EntityData) AddBaseEvent() {
	e.AddEffect(event.StartFight, effect.FightBaseEff, effect.None, effect.Normal, func(a ...any) any {
		e.AbilitiesCdReset()
		e.Stats = e.CalculateStats()
		return nil
	})
	e.AddEffect(event.StartTurn, effect.FightBaseEff, effect.None, effect.Normal, func(a ...any) any {
		e.RestoreEnergy()
		e.AbilityCdReduction()
		e.BuffCdReduction()
		return nil
	})
	e.AddEffect(event.CastAbility, effect.FightBaseEff, effect.None, effect.Normal, func(a ...any) any {
		v, ok := a[0].(int)
		if !ok {
			panic("Ammazzati grazie")
		}
		e.Abilities[v].CastAbility(&e.Stats)
		return nil
	})
	e.AddEffect(event.EndTurn, effect.QtrEff, effect.None, effect.Normal, func(a ...any) any {
		return e.QtrActivasion()
	})
	e.AddEffect(event.Died, effect.ReviveEff, effect.Before, effect.Normal, func(a ...any) any {
		resRate := float32(rand.Intn(100))
		return resRate < e.Stats[stat.RES]
	})
}
func (e *EntityData) AddEffect(ev event.EventType, ef effect.IdEffect, trg effect.TriggerTiming, priority effect.PriorityType, f func(a ...any) any) uint16 {
	var id uint16
	if len(e.Effects) > 0 {
		id = 1
	} else {
		id = e.Effects[ev][len(e.Effects[ev])-1].Id + 1
	}
	e.Effects[ev] = append(e.Effects[ev], effect.NewEffect(id, ef, ev, trg, priority, func(a ...any) any {
		e.EventHandler(ev, a)
		return f(a)
	}))
	return id
}
func (e *EntityData) RemoveEffect(ev event.EventType, id uint16) {
	for i := range e.Effects[ev] {
		if e.Effects[ev][i].Id == id {
			e.Effects[ev] = append(e.Effects[ev][:i], e.Effects[ev][i+1:]...)
			return
		}
	}
}
func (e *EntityData) AddBuff(b effect.Buff, a ...any) uint16 {
	var id uint16
	if len(e.Buffs) > 0 {
		id = e.Buffs[len(e.Buffs)-1].Id + 1
	} else {
		id = 1
	}
	b.Id = id
	e.Buffs = append(e.Buffs, b)
	e.Buffs[len(e.Buffs)-1].GetBuff(a)
	return id
}
func (e *EntityData) RemoveBuff(id uint16) {
	for i := range e.Buffs {
		if e.Buffs[i].Id == id {
			e.Buffs[i].RemoveBuff()
			e.Buffs = append(e.Buffs[:i], e.Buffs[i+1:]...)
		}
	}
}
func (e *EntityData) TriggerEffects(ev event.EventType, trg effect.TriggerTiming, a ...any) map[effect.IdEffect]any {
	foo := make(map[effect.IdEffect]any)
	for _, ef := range e.GetEffectsOrdered(ev) {
		if ef.Trigger == trg {
			//	GLI EFFETTI SUCCESSIVI DOVREBBERO SAPERE COSA HANNO FATTO GLI EFFETTI PRECEDENTI
			app := ef.Activate(a...)
			if app != nil {
				foo[ef.IdEff] = app
			}
		}
	}
	return foo
}
func (e *EntityData) GetEffectsOrdered(ev event.EventType) []effect.Effect {
	effect := append([]effect.Effect(nil), e.Effects[ev]...)

	sort.SliceStable(effect, func(i, j int) bool {
		return effect[i].Priority < effect[j].Priority
	})

	return effect
}

func (e *EntityData) AbilitiesCdReset() {
	for i := range e.Abilities {
		e.Abilities[i].Cooldown = 0
	}
}
func (e *EntityData) RestoreEnergy() {
	e.Stats[stat.ST] += e.Stats[stat.ST_RG]
	if e.Stats[stat.ST] > e.MaxSt {
		e.Stats[stat.ST] = e.MaxSt
	}
	e.Stats[stat.MP] += e.Stats[stat.MP_RG]
	if e.Stats[stat.MP] > e.MaxMp {
		e.Stats[stat.MP] = e.MaxMp
	}
}
func (e *EntityData) AbilityCdReduction() {
	for i := range e.Abilities {
		e.Abilities[i].Cooldown--
	}
}
func (e *EntityData) BuffCdReduction() {
	for i := len(e.Buffs) - 1; i >= 0; i-- {
		if !e.Buffs[i].DecreaseDuration() {
			e.Buffs = append(e.Buffs[:i], e.Buffs[i+1:]...)
		}
	}
}
func (e *EntityData) QtrActivasion() bool {
	return rand.Intn(100) < int(e.Stats[stat.QTR])
}

//FIGHT FUNCS

func (e *EntityData) ActiveAbility(ability int, opponents any) map[int][]model.Damage {
	e.TriggerEffects(event.CastAbility, effect.None, ability)
	result := e.Abilities[ability].execute.Activate(opponents)

	atks, ok := result.(map[int][]model.Damage)
	if !ok {
		fmt.Printf("ability=%v\n", ability)
		fmt.Printf("result=%#v\n", result)
		fmt.Printf("type=%T\n", result)
		panic(fmt.Sprintf("Activate returned %T, expected map[int][]model.Damage", result))
	}

	return atks
}

func (e *EntityData) HasHit(defenderStat stat.Stats, abilityAcc float32) bool {
	return rand.Intn(100) < int(e.Stats[stat.ACC]+abilityAcc-defenderStat[stat.EVA])
}

func (e *EntityData) HasCritted() model.DamageType {
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

func (e *EntityData) GetCritDmg(defenderStat stat.Stats, dmg, spDmg float32, critType model.DamageType) (float32, float32) {
	if defenderStat[stat.CRIT_RES] >= 100 {
		panic("COGLIONE")
	}
	switch critType {
	case model.Flat:
		return dmg, spDmg
	case model.Crit:
		return dmg * ((1 + e.Stats[stat.CRIT_DMG]/100) * (1 - defenderStat[stat.CRIT_RES]/100)), spDmg * ((1 + e.Stats[stat.CRIT_DMG]/100) * (1 - defenderStat[stat.CRIT_RES]/100))
	case model.SuperCrit:
		return dmg * ((1 + e.Stats[stat.CRIT_DMG]/100) * 1.5 * (1 - defenderStat[stat.CRIT_RES]/100)), spDmg * ((1 + e.Stats[stat.CRIT_DMG]/100) * 1.5 * (1 - defenderStat[stat.CRIT_RES]/100))
	default:
		panic("get out")
	}
}

func (e *EntityData) EventHandler(ev event.EventType, a ...any) {
	switch ev {
	case event.StartFight:
		_, ok := a[0].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[1].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[2].(area.AreaId)
		if !ok {
			panic("Daje roma cazzo")
		}
	case event.EndFight:
		_, ok := a[0].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
	case event.StartTurn, event.EndTurn, event.QtrTrigged:
		_, ok := a[0].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[1].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}

	case event.CastAbility:
		_, ok := a[0].(int)
		if !ok {
			panic("Daje roma cazzo")
		}
	case event.DamageDealt, event.DamageTaken:
		_, ok := a[0].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[1].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}

		_, ok = a[2].(int)
		if !ok {
			panic("Daje roma cazzo")
		}

		_, ok = a[3].(float32)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[4].(model.DamageType)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[5].(bool)
		if !ok {
			panic("Daje roma cazzo")
		}
	case event.HealDealt, event.HealDTaken:
		_, ok := a[0].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[1].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[2].(int)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[3].(float32)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[4].(model.HealType)
		if !ok {
			panic("Daje roma cazzo")
		}
	case event.BuffApplied, event.BuffExpired:
		_, ok := a[0].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[1].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[2].(*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[3].(bool)
		if !ok {
			panic("Daje roma cazzo")
		}
	case event.Died:
		_, ok := a[0].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[1].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[2].(*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[3].(float32)
		if !ok {
			panic("Daje roma cazzo")
		}
	case event.EnergyRecovered, event.EnergySpent:
		_, ok := a[0].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[1].([]*EntityData)
		if !ok {
			panic("Daje roma cazzo")
		}
		_, ok = a[2].(float32)
		if !ok {
			panic("Daje roma cazzo")
		}
	}
}
