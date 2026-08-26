package combat

import (
	"fmt"
	"gioco/game/combat/effect"
	"gioco/game/combat/event"
	"gioco/game/comm"
	"gioco/game/comm/area"
	"gioco/model"
	"gioco/model/entity"
	"gioco/model/entity/character"
	"gioco/model/entity/enemy"
	"gioco/model/stat"
	"maps"
	"slices"
)

func Fight(pg []character.Character, opponents []enemy.Enemy, area area.Area) {

	team := make([]entity.Entity, len(pg))
	for i := range pg {
		team[i] = pg[i].GetCharacter()
	}
	opponentTeam := make([]entity.Entity, len(opponents))
	for i := range opponents {
		opponentTeam[i] = opponents[i].GetEnemy()
	}

	triggerEvent(event.StartFight, team, opponentTeam, area, effect.None)

	for len(team) > 0 && len(opponentTeam) > 0 {
		index, isTeam := getFasterEntity(team, opponentTeam)
		if isTeam {
			opponentTeam = attack(team, index, opponentTeam)
		} else {
			team = attack(opponentTeam, index, team)
		}
	}

	if len(team) > 0 {
		triggerEvent(event.EndFight, team, opponentTeam, area, effect.None)
	} else {
		panic("YOU LOOSE!!!")
	}
}
func attack(team []entity.Entity, actor int, opponents []entity.Entity) []entity.Entity {

	teamData := make([]*entity.EntityData, len(team))
	for i := range team {
		teamData[i] = team[i].GetEntity()
	}
	opponentsData := make([]*entity.EntityData, len(opponents))
	for i := range opponents {
		opponentsData[i] = opponents[i].GetEntity()
	}
	fmt.Println(comm.Cyan.Write(teamData[actor].Name) + " sta attaccando:")

	teamData[actor].TriggerEffects(event.StartTurn, effect.None, team, opponents)

qtrTrigged:
	iAbility := team[actor].GetAbility()
	fmt.Printf("iAbility: %v\n", iAbility)
	if iAbility == -1 {
		fmt.Println(comm.BrightYellow.Write(teamData[actor].Name) + " non ha abbastanza energia per utilizzare nessuna abilità")

		teamData[actor].Stats[stat.ST] *= 1.30
		if teamData[actor].Stats[stat.ST] > teamData[actor].MaxSt {
			teamData[actor].Stats[stat.ST] = teamData[actor].MaxSt
		}
		teamData[actor].Stats[stat.MP] *= 1.30
		if teamData[actor].Stats[stat.MP] > teamData[actor].MaxMp {
			teamData[actor].Stats[stat.MP] = teamData[actor].MaxMp
		}
		teamData[actor].Stats[stat.SPD] += 500 

		eff := teamData[actor].TriggerEffects(event.EndTurn, effect.None, team, opponents)
		for ef, value := range eff {
			if ef == effect.QtrEff {
				if value.(bool) && len(opponents) > 0 {
					fmt.Println(comm.Cyan.Write(teamData[actor].Name) + " ha attivato il QTR!!")
					teamData[actor].TriggerEffects(event.QtrTrigged, effect.None, team, opponents)
					goto qtrTrigged
				}
			}
		}
		return opponents
	}
	targets := teamData[actor].ActiveAbility(iAbility, opponentsData)

	for i := range targets {
		for j := range targets[i] {
			teamData[actor].TriggerEffects(event.DamageDealt, effect.Before, team, opponents, actor, targets[i][j], targets[i][j].DmgType)
			opponentsData[actor].TriggerEffects(event.DamageTaken, effect.Before, opponents, team, actor, targets[i][j], targets[i][j].DmgType)

			switch targets[i][j].DmgType {
			case model.Crit:
				fmt.Println(comm.BrightBlue.Write(teamData[actor].Name) + " ha fatto un danno critico!")
			case model.SuperCrit:
				fmt.Println(comm.BrightBlue.Write(teamData[actor].Name) + " ha fatto un danno super critico!!")
			case model.Miss:
				fmt.Println(comm.BrightBlue.Write(teamData[actor].Name) + " ha mancato l'attacco")
			}
			opponentsData[actor].Stats[stat.HP] -= targets[i][j].Amount

			fmt.Printf(comm.Cyan.Write(teamData[actor].Name)+" ha inflitto %v a "+opponentsData[actor].Name+"\n", targets[i][j].Amount)

			teamData[actor].TriggerEffects(event.DamageDealt, effect.After, team, opponents, actor, targets[i][j], targets[i][j].DmgType)
			if opponentsData[actor].Stats[stat.HP] <= 0 {
				break
			}
			opponentsData[actor].TriggerEffects(event.DamageTaken, effect.After, opponents, team, actor, targets[i][j], targets[i][j].DmgType)
		}
	}
	app := 0
	for i := range slices.Sorted(maps.Keys(targets)) {
		if opponentsData[i-app].Stats[stat.HP] <= 0 {
			effs := opponentsData[i-app].TriggerEffects(event.Died, effect.Before)
			if !effs[effect.ReviveEff].(bool) {
				opponentsData[i-app].TriggerEffects(event.Died, effect.After)
				fmt.Println(comm.BrightCyan.Write(opponentsData[i-app].Name) + " è stato eliminato")
				opponents = append(opponents[:i-app], opponents[i+1-app:]...)
				opponentsData = append(opponentsData[:i-app], opponentsData[i+1-app:]...)
				app++
			} else {
				fmt.Println(comm.BrightCyan.Write(opponentsData[i-app].Name) + " è sopravvissuto")
			}
		}
	}

	eff := teamData[actor].TriggerEffects(event.EndTurn, effect.None, team, opponents)
	for ef, value := range eff {
		if ef == effect.QtrEff {
			if value.(bool) && len(opponents) > 0 {
				fmt.Println(comm.Cyan.Write(teamData[actor].Name) + " ha attivato il QTR!!")
				teamData[actor].TriggerEffects(event.QtrTrigged, effect.None, team, opponents)
				goto qtrTrigged
			}
		}
	}
	return opponents
}

func triggerEvent(ev event.EventType, team []entity.Entity, opponentTeam []entity.Entity, area area.Area, trg effect.TriggerTiming, a ...any) {
	for i := range team {
		team[i].GetEntity().TriggerEffects(ev, trg, team, opponentTeam, area, a)
	}
	for i := range opponentTeam {
		opponentTeam[i].GetEntity().TriggerEffects(ev, trg, opponentTeam, team, area, a)
	}
	area.TriggerEffects(ev, trg, a)
}
func getFasterEntity(team []entity.Entity, opponentTeam []entity.Entity) (int, bool) {
	var idEntity int
	var maxSpd float32
	isTeam := true
	for i := range team {
		if team[i].GetEntity().Stats[stat.SPD] > maxSpd {
			maxSpd = team[i].GetEntity().Stats[stat.SPD]
			idEntity = i
		}
	}
	for i := range opponentTeam {
		if opponentTeam[i].GetEntity().Stats[stat.SPD] > maxSpd {
			maxSpd = opponentTeam[i].GetEntity().Stats[stat.SPD]
			isTeam = false
		}
	}
	return idEntity, isTeam
}
