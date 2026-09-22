package main

import (
	"github.com/KONshougun/K-game/game/characters"
	"github.com/KONshougun/K-game/game/combat/effect"
	"github.com/KONshougun/K-game/game/combat/event"
	"github.com/KONshougun/K-game/game/combat/handler"
	"github.com/KONshougun/K-game/game/comm/area"
	"github.com/KONshougun/K-game/game/enemies"
	"github.com/KONshougun/K-game/gui"
	battlescene "github.com/KONshougun/K-game/gui/scene/battleScene"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/KONshougun/K-game/model/stat"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	b := enemies.NewBandit(100, nil, nil)
	k := characters.NewKnight(100, nil, nil)

	//team := []character.Character{k}
	//opponentTeam := []enemy.Enemy{b}

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
	initGui(characters, enemies, area.Plains)
	//combat.InitBattle(team, opponentTeam)
}

func initGui(characters, enemies []*entity.EntityData, area area.AreaId) {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Gioco")

	ebiten.SetFullscreen(true)
	ebiten.RunGame(gui.NewGame(
		battlescene.NewBattleScene(characters, enemies, area),
	))
}
