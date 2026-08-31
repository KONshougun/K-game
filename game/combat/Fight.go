package combat

import (
	"fmt"
	"gioco/game/combat/effect"
	"gioco/game/combat/event"
	"gioco/game/combat/handler"
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

func InitBattle(characters []character.Character, enemies []enemy.Enemy) {

	team := make([]*entity.EntityData, len(characters))
	for i := range characters {
		team[i] = characters[i].GetCharacter().GetEntity()
		team[i].TriggerEffects(event.SetEntityData, effect.None, &handler.SetEntityData{
			FlatStats:    stat.Stats{},
			PercentStats: stat.Stats{},
		})
	}
	opponents := make([]*entity.EntityData, len(enemies))
	for i := range enemies {
		opponents[i] = enemies[i].GetEnemy().GetEntity()
		opponents[i].TriggerEffects(event.SetEntityData, effect.None, &handler.SetEntityData{
			FlatStats:    stat.Stats{},
			PercentStats: stat.Stats{},
		})
	}

	Fight(team, opponents, area.Area{Id: area.Plains})
}

func Fight(team, opponents []*entity.EntityData, area area.Area) {

	triggerEvent(event.StartFight, team, opponents, area, effect.None, &handler.StartFight{
		Team:      team,
		Opponents: opponents,
		Area:      area.Id,
	})

	for len(team) > 0 && len(opponents) > 0 {
		index, isTeam := getFasterEntity(team, opponents)
		if isTeam {
			team, opponents = attack(team, opponents, area, index)
		} else {
			team, opponents = attack(opponents, team, area, index)
		}
	}

	if len(team) > 0 {
		triggerEvent(event.EndFight, team, opponents, area, effect.None, &handler.EndFight{
			Team: team,
			Area: area.Id,
		})
	} else {
		panic("YOU LOOSE!!!")
	}
}
func attack(team, opponents []*entity.EntityData, area area.Area, actorIdx int) ([]*entity.EntityData, []*entity.EntityData) {

	actor := team[actorIdx]
	fmt.Println(comm.Cyan.Write(actor.Name) + " sta attaccando:")

	actor.TriggerEffects(event.StartTurn, effect.None, &handler.StartTurn{
		Team:      team,
		Opponents: opponents,
		Area:      area.Id,
	})

qtrTrigged:
	targets := make(map[int][]model.Damage)
	idxAbility := 0
	targets, idxAbility = actor.Attack(actor, opponents)

	if idxAbility == -1 {
		fmt.Println(comm.BrightYellow.Write(actor.Name) + " non ha abbastanza energia per utilizzare nessuna abilità")

		actor.Stats[stat.ST] *= 1.30
		if actor.Stats[stat.ST] > actor.MaxSt {
			actor.Stats[stat.ST] = actor.MaxSt
		}
		actor.Stats[stat.MP] *= 1.30
		if actor.Stats[stat.MP] > actor.MaxMp {
			actor.Stats[stat.MP] = actor.MaxMp
		}
		actor.Stats[stat.SPD] -= 500 + actor.Stats[stat.ATK_SPD]
	}

	for i := range targets {
		for j := range targets[i] {
			actor.TriggerEffects(event.DamageDealt, effect.Before, &handler.DamageDealt{
				Team:      team,
				Opponents: opponents,
				Area:      area.Id,
				Target:    opponents[i],
				Amount:    targets[i][j].Amount,
				DmgType:   targets[i][j].DmgType,
			})
			opponents[i].TriggerEffects(event.DamageTaken, effect.Before, &handler.DamageTaken{
				Team:      team,
				Opponents: opponents,
				Area:      area.Id,
				Actor:     actor,
				Amount:    targets[i][j].Amount,
				DmgType:   targets[i][j].DmgType,
			})

			switch targets[i][j].DmgType {
			case model.Flat:
				fmt.Print("")
			case model.Crit:
				fmt.Println(comm.BrightBlue.Write(actor.Name) + " ha fatto un danno critico!")
			case model.SuperCrit:
				fmt.Println(comm.BrightBlue.Write(actor.Name) + " ha fatto un danno super critico!!")
			case model.Pure:
				fmt.Println(comm.BrightBlue.Write(actor.Name) + " ha fatto danni puri!!")
			case model.Miss:
				fmt.Println(comm.BrightBlue.Write(actor.Name) + " ha mancato l'attacco")
			default:
				fmt.Printf("targets[i][j].DmgType: %v\n", targets[i][j].DmgType)
				panic("sega")
			}
			opponents[i].Stats[stat.HP] -= targets[i][j].Amount

			fmt.Printf(comm.Cyan.Write(actor.Name)+" ha inflitto %v a "+opponents[i].Name+"\n", targets[i][j].Amount)

			actor.TriggerEffects(event.DamageDealt, effect.After, &handler.DamageDealt{
				Team:      team,
				Opponents: opponents,
				Area:      area.Id,
				Target:    opponents[i],
				Amount:    targets[i][j].Amount,
				DmgType:   targets[i][j].DmgType,
			})
			if opponents[i].Stats[stat.HP] <= 0 {
				break
			}
			opponents[i].TriggerEffects(event.DamageTaken, effect.After, &handler.DamageTaken{
				Team:      team,
				Opponents: opponents,
				Area:      area.Id,
				Actor:     actor,
				Amount:    targets[i][j].Amount,
				DmgType:   targets[i][j].DmgType,
			})
		}
	}
	//Controllo se ci sono stati morti
	app := 0
	for i := range slices.Sorted(maps.Keys(targets)) {
		if opponents[i-app].Stats[stat.HP] <= 0 {
			evt := opponents[i-app].TriggerEffects(event.Died, effect.Before, &handler.Died{
				Team:      team,
				Opponents: opponents,
				Area:      area.Id,
				Actor:     actor,
				Amount:    targets[i][len(targets[i])-1].Amount,
				DmgType:   targets[i][len(targets[i])-1].DmgType,
				IsAlive:   false,
			})
			if !*(evt.GetData()[6].(*bool)) {
				evt = opponents[i-app].TriggerEffects(event.Died, effect.After, &handler.Died{
					Team:      team,
					Opponents: opponents,
					Area:      area.Id,
					Actor:     actor,
					Amount:    targets[i][len(targets[i])-1].Amount,
					DmgType:   targets[i][len(targets[i])-1].DmgType,
					IsAlive:   false,
				})
				if !*(evt.GetData()[6].(*bool)) {
					fmt.Println(comm.BrightCyan.Write(opponents[i-app].Name) + " è stato eliminato")
					opponents = append(opponents[:i-app], opponents[i+1-app:]...)
					app++
				} else {
					fmt.Println(comm.BrightCyan.Write(opponents[i-app].Name) + " è rinato")
					opponents[i-app].TriggerEffects(event.Revived, effect.After, &handler.Revived{
						Team:      team,
						Opponents: opponents,
						Area:      area.Id,
					})
				}
			} else {
				fmt.Println(comm.BrightCyan.Write(opponents[i-app].Name) + " è sopravvissuto")
				opponents[i-app].TriggerEffects(event.Resist, effect.After, &handler.Resist{
					Team:      team,
					Opponents: opponents,
					Area:      area.Id,
				})
			}
		}
	}

	evt := actor.TriggerEffects(event.EndTurn, effect.None, &handler.EndTurn{
		Team:         team,
		Opponents:    opponents,
		Area:         area.Id,
		QtrActivated: false,
	})
	if *(evt.GetData()[3].(*bool)) && len(opponents) > 0 {
		fmt.Println(comm.Cyan.Write(actor.Name) + " ha attivato il QTR!!")
		actor.TriggerEffects(event.QtrTrigged, effect.None, &handler.QtrTrigged{
			Team:      team,
			Opponents: opponents,
			Area:      area.Id,
		})
		goto qtrTrigged
	}
	return team, opponents
}

func triggerEvent(ev event.EventType, team []*entity.EntityData, opponentTeam []*entity.EntityData, area area.Area, trg effect.TriggerTiming, evtData event.EventData) {
	for i := range team {
		team[i].TriggerEffects(ev, trg, evtData)
	}
	for i := range opponentTeam {
		opponentTeam[i].TriggerEffects(ev, trg, evtData)
	}
	area.TriggerEffects(ev, trg, evtData)
}
func getFasterEntity(team []*entity.EntityData, opponentTeam []*entity.EntityData) (int, bool) {
	var idEntity int
	var maxSpd float32
	isTeam := true
	for i := range team {
		if team[i].Stats[stat.SPD] > maxSpd {
			maxSpd = team[i].Stats[stat.SPD]
			idEntity = i
		}
	}
	for i := range opponentTeam {
		if opponentTeam[i].Stats[stat.SPD] > maxSpd {
			maxSpd = opponentTeam[i].Stats[stat.SPD]
			isTeam = false
		}
	}
	return idEntity, isTeam
}
