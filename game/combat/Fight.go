package combat

import (
	"fmt"
	"gioco/game/combat/effect"
	"gioco/game/combat/event"
	"gioco/game/comm"
	"gioco/game/comm/area"
	"gioco/model"
	"gioco/model/entity"
	"gioco/model/stat"
	"maps"
	"slices"
)

func Fight(characters []*entity.Entity, enemies []*entity.Entity, area area.Area) {
	triggerEvent(event.StartFight, characters, enemies, area, effect.None)

	//fmt.Printf("characters[0].Stats: %v\n", characters[0].Stats)
	//fmt.Printf("enemies[0].Stats: %v\n", enemies[0].Stats)

	for len(characters) > 0 && len(enemies) > 0 {
		iChar := getFasterEntity(characters)
		iEnem := getFasterEntity(enemies)
		if characters[iChar].Stats[stat.SPD] > enemies[iEnem].Stats[stat.SPD] {
			enemies = attack(characters, iChar, enemies)
		} else {
			characters = attack(enemies, iEnem, characters)
		}
	}

	if len(characters) > 0 {
		triggerEvent(event.EndFight, characters, enemies, area, effect.None)
	} else {
		panic("YOU LOOSE!!!")
	}
}
func attack(team []*entity.Entity, actor int, opponents []*entity.Entity) []*entity.Entity {
	fmt.Println(comm.Cyan.Write(team[actor].Name) + " sta attaccando:")

	team[actor].TriggerEffects(event.StartTurn, effect.None, team, opponents)

qtrTrigged:
	iAbility := team[actor].Controller.GetAbility(*team[actor])
	targets := team[actor].ActiveAbility(iAbility, opponents)

	for i := range targets {
		for j := range targets[i] {
			team[actor].TriggerEffects(event.DamageDealt, effect.Before, team, opponents, actor, targets[i][j], targets[i][j].DmgType)
			opponents[i].TriggerEffects(event.DamageTaken, effect.Before, opponents, team, actor, targets[i][j], targets[i][j].DmgType)

			switch targets[i][j].DmgType {
			case model.Crit:
				fmt.Println(comm.BrightBlue.Write(team[actor].Name) + " ha fatto un danno critico!")
			case model.SuperCrit:
				fmt.Println(comm.BrightBlue.Write(team[actor].Name) + " ha fatto un danno super critico!!")
			case model.Miss:
				fmt.Println(comm.BrightBlue.Write(team[actor].Name) + " ha mancato l'attacco")
			}
			opponents[i].Stats[stat.HP] -= targets[i][j].Amount

			fmt.Printf(comm.Cyan.Write(team[actor].Name)+" ha inflitto %v a "+opponents[i].Name+"\n", targets[i][j].Amount)

			team[actor].TriggerEffects(event.DamageDealt, effect.After, team, opponents, actor, targets[i][j], targets[i][j].DmgType)
			if opponents[i].Stats[stat.HP] <= 0 {
				break
			}
			opponents[i].TriggerEffects(event.DamageTaken, effect.After, opponents, team, actor, targets[i][j], targets[i][j].DmgType)
		}
	}
	app := 0
	for i := range slices.Sorted(maps.Keys(targets)) {
		if opponents[i-app].Stats[stat.HP] <= 0 {
			effs := opponents[i-app].TriggerEffects(event.Died, effect.Before)
			if !effs[effect.Revive].(bool) {
				opponents[i-app].TriggerEffects(event.Died, effect.After)
				fmt.Println(comm.BrightCyan.Write(opponents[i-app].Name) + " è stato eliminato")
				opponents = append(opponents[:i-app], opponents[i+1-app:]...)
				app++
			} else {
				fmt.Println(comm.BrightCyan.Write(opponents[i-app].Name) + " è sopravvissuto")
			}
		}
	}

	eff := team[actor].TriggerEffects(event.EndTurn, effect.None, team, opponents)
	for ef, value := range eff {
		if ef == effect.QtrEff {
			v := ef.HandleEffect(value)
			if v.(bool) && len(opponents) > 0 {
				fmt.Println(comm.Cyan.Write(team[actor].Name) + " ha attivato il QTR!!")
				team[actor].TriggerEffects(event.QtrTrigged, effect.None, team, opponents)
				goto qtrTrigged
			}
		}
	}
	return opponents
}

func triggerEvent(ev event.EventType, characters, enemies []*entity.Entity, area area.Area, trg effect.TriggerTiming, a ...any) {
	for i := range characters {
		(characters)[i].TriggerEffects(ev, trg, characters, enemies, area, a)
	}
	for i := range enemies {
		(enemies)[i].TriggerEffects(ev, trg, enemies, characters, area, a)
	}
	area.TriggerEffects(ev, trg, a)
}
func getFasterEntity(team []*entity.Entity) int {
	maxSpd := 0
	for i := 1; i < len(team); i++ {
		if team[i].Stats[stat.SPD] > team[maxSpd].Stats[stat.SPD] {
			maxSpd = i
		}
	}
	return maxSpd
}
