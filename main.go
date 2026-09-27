package main

import (
	"github.com/KONshougun/K-game/gui"
	startgame "github.com/KONshougun/K-game/gui/scene/startGame"
)

func main() {

	/*b := enemies.NewBandit(100, nil, nil)
	k := characters.NewKnight(100, nil, nil)

	team := []character.Character{k}
	opponentTeam := []enemy.Enemy{b}

	characters, enemies := []*entity.EntityData{k.GetCharacter().GetEntity()}, []*entity.EntityData{b.GetEnemy().GetEntity()}
	for _, c := range characters {
		c.TriggerEffects(event.SetEntityData, effect.None, &handler.SetEntityData{
			FlatStats:    stat.Stats{},
			PercentStats: stat.Stats{},
		})
		c.Shield = c.Stats[stat.Hp] / 2
	}
	for _, e := range enemies {
		e.TriggerEffects(event.SetEntityData, effect.None, &handler.SetEntityData{
			FlatStats:    stat.Stats{},
			PercentStats: stat.Stats{},
		})
		//e.Shield = e.Stats[stat.Hp] / 2
	}

	combat.NewBattle(team, opponentTeam)*/

	initGui()
}

func initGui() {
	gui.NewGame(startgame.NewStartGameScene())
}
