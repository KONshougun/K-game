package combat

import (
	"fmt"

	"github.com/KONshougun/K-game/game/combat/effect"
	"github.com/KONshougun/K-game/game/combat/event"
	"github.com/KONshougun/K-game/game/combat/handler"
	"github.com/KONshougun/K-game/game/comm"
	"github.com/KONshougun/K-game/model"
	"github.com/KONshougun/K-game/model/area"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/KONshougun/K-game/model/entity/character"
	"github.com/KONshougun/K-game/model/entity/enemy"
	"github.com/KONshougun/K-game/model/stat"
)

type BattleData struct {
	Team      character.Team
	Opponents enemy.Opponents
	EnemiesDeath enemy.Opponents
	Area      area.Area
	Turn      int
}

func (btlData BattleData) StartFight() {
	for i := range btlData.Team {
		btlData.Team[i].GetCharacter().TriggerEffects(event.SetEntityData, effect.None, &handler.SetEntityData{
			FlatStats:    stat.Stats{},
			PercentStats: stat.Stats{},
		})
	}
	for i := range btlData.Opponents {
		btlData.Opponents[i].GetEnemy().TriggerEffects(event.SetEntityData, effect.None, &handler.SetEntityData{
			FlatStats:    stat.Stats{},
			PercentStats: stat.Stats{},
		})
	}
	team := btlData.Team.GetEntities()
	opponents := btlData.Opponents.GetEntities()
	TriggerEvent(event.StartFight, team, opponents, btlData.Area, effect.None, &handler.StartFight{
		Team:      team,
		Opponents: opponents,
		Area:      btlData.Area.Id,
	})
}

func (btlData BattleData) EndFight() bool {
	team := make([]*entity.EntityData, 0, len(btlData.Team))
	if len(team) > 0 {
		TriggerEvent(event.EndFight, btlData.Team.GetAliveEntities(), btlData.Opponents.GetAliveEntities(), btlData.Area, effect.None, &handler.EndFight{
			Team: team,
			Area: btlData.Area.Id,
		})
		return true
	} else {
		return false
	}
}
func (btlData BattleData) StartTurn() (entity.EntityType, int, []bool) {
	idx, eType := GetFasterEntity(btlData.Team.GetAliveEntities(), btlData.Opponents.GetAliveEntities())
	evData := &handler.StartTurn{
		Team:      btlData.Team.GetAliveEntities(),
		Opponents: btlData.Opponents.GetAliveEntities(),
		Area:      btlData.Area.Id,
	}
	switch eType {
	case entity.Character:
		actor := btlData.Team[idx].GetCharacter().GetEntity()
		actor.TriggerEffects(event.StartTurn, effect.None, evData)
		return eType, idx, actor.AvailableAbilities()
	case entity.Enemy:
		actor := btlData.Opponents[idx].GetEnemy().GetEntity()
		actor.TriggerEffects(event.StartTurn, effect.None, evData)
		return eType, idx, actor.AvailableAbilities()
	default:
		panic("vattene a fanculo")
	}
}
func (btlData BattleData) StartAttack(actorType entity.EntityType, actor *entity.EntityData, abilityIdx int) map[int][]model.Damage {
	var team, opponents []*entity.EntityData
	switch actorType {
	case entity.Character:
		team = btlData.Team.GetAliveEntities()
		opponents = btlData.Opponents.GetAliveEntities()
	case entity.Enemy:
		team = btlData.Opponents.GetAliveEntities()
		opponents = btlData.Team.GetAliveEntities()
	}
	targets := actor.CastAbility(abilityIdx, opponents)

	//TODO FORSE è RIMOVIBILE
	for i := range targets {
		for j := range targets[i] {
			actor.TriggerEffects(event.DamageDealt, effect.Before, &handler.DamageDealt{
				Team:      team,
				Opponents: opponents,
				Area:      btlData.Area.Id,
				Target:    opponents[i],
				Amount:    targets[i][j].Amount,
				DmgType:   targets[i][j].DmgType,
				DmgTarget: targets[i][j].DmgTarget,
			})
			opponents[i].TriggerEffects(event.DamageTaken, effect.Before, &handler.DamageTaken{
				Team:      team,
				Opponents: opponents,
				Area:      btlData.Area.Id,
				Actor:     actor,
				Amount:    targets[i][j].Amount,
				DmgType:   targets[i][j].DmgType,
				DmgTarget: targets[i][j].DmgTarget,
			})

			opponents[i].Dameged(targets[i][j].Amount, targets[i][j].DmgType, targets[i][j].DmgTarget)

			actor.TriggerEffects(event.DamageDealt, effect.After, &handler.DamageDealt{
				Team:      team,
				Opponents: opponents,
				Area:      btlData.Area.Id,
				Target:    opponents[i],
				Amount:    targets[i][j].Amount,
				DmgType:   targets[i][j].DmgType,
				DmgTarget: targets[i][j].DmgTarget,
			})
			if opponents[i].Stats[stat.Hp] <= 0 {
				break
			}
			opponents[i].TriggerEffects(event.DamageTaken, effect.After, &handler.DamageTaken{
				Team:      team,
				Opponents: opponents,
				Area:      btlData.Area.Id,
				Actor:     actor,
				Amount:    targets[i][j].Amount,
				DmgType:   targets[i][j].DmgType,
				DmgTarget: targets[i][j].DmgTarget,
			})
		}
	}
	return targets
}

func (btlData BattleData) CheckDeath(targets map[int][]model.Damage) {
	//Controllo se ci sono stati morti

	death := true

	for death {
		death = false

		//Controllo i characters
		//for

		//Controllo gli enemy
	}

	/*
		found := true
		for found {
			found = false
			app := 0
			for i := range slices.Sorted(maps.Keys(targets)) {
				if opponents[i-app].Stats[stat.Hp] <= 0 {
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
							found = true
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
		}*/
}
func (btlData BattleData) EndTurn(actorType entity.EntityType, actor *entity.EntityData) bool {
	team := btlData.Team.GetAliveEntities()
	opponents := btlData.Opponents.GetAliveEntities()
	evt := actor.TriggerEffects(event.EndTurn, effect.None, &handler.EndTurn{
		Team:         team,
		Opponents:    opponents,
		Area:         btlData.Area.Id,
		QtrActivated: false,
	})
	if *(evt.GetData()[3].(*bool)) && len(opponents) > 0 {
		fmt.Println(comm.Cyan.Write(actor.Name) + " ha attivato il QTR!!")
		actor.TriggerEffects(event.QtrTrigged, effect.None, &handler.QtrTrigged{
			Team:      team,
			Opponents: opponents,
			Area:      btlData.Area.Id,
		})
		return true
	}
	return false
}
func TriggerEvent(ev event.EventType, team []*entity.EntityData, opponentTeam []*entity.EntityData, area area.Area, trg effect.TriggerTiming, evtData event.EventData) {
	for i := range team {
		team[i].TriggerEffects(ev, trg, evtData)
	}
	for i := range opponentTeam {
		opponentTeam[i].TriggerEffects(ev, trg, evtData)
	}
	area.TriggerEffects(ev, trg, evtData)
}
func GetFasterEntity(team []*entity.EntityData, opponentTeam []*entity.EntityData) (int, entity.EntityType) {
	var idEntity int
	var maxSpd float64
	var eT entity.EntityType
	for i := range team {
		if team[i].Stats[stat.Spd] > maxSpd {
			maxSpd = team[i].Stats[stat.Spd]
			idEntity = i
			eT = entity.Character
		}
	}
	for i := range opponentTeam {
		if opponentTeam[i].Stats[stat.Spd] > maxSpd {
			maxSpd = opponentTeam[i].Stats[stat.Spd]
			eT = entity.Enemy
		}
	}
	return idEntity, eT
}

/*

func attack(team, opponents []*entity.EntityData, area area.Area, actorIdx int) ([]*entity.EntityData, []*entity.EntityData) {

qtrTrigged:
		//Controllo se ci sono stati morti
		found := true
		for found {
			found = false
			app := 0
			for i := range slices.Sorted(maps.Keys(targets)) {
				if opponents[i-app].Stats[stat.Hp] <= 0 {
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
							found = true
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

	func Fight(btlData BattleData) bool {
		team, opponents := NewBattle(btlData)

		for len(team) > 0 && len(opponents) > 0 {
			index, eType := GetFasterEntity(team, opponents)
			switch eType {
			case entity.Character:
				team, opponents = attack(team, opponents, btlData.Area, index)
			case entity.Enemy:
				team, opponents = attack(opponents, team, btlData.Area, index)

			}
		}
		return EndBattle(team, opponents, btlData.Area)
	}
*/
