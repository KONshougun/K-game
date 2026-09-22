package main

import (
	"github.com/KONshougun/K-game/game/characters"
	"github.com/KONshougun/K-game/game/comm/area"
	"github.com/KONshougun/K-game/game/enemies"
	"github.com/KONshougun/K-game/gui"
	battlescene "github.com/KONshougun/K-game/gui/scene/battleScene"
	"github.com/KONshougun/K-game/model/entity/character"
	"github.com/KONshougun/K-game/model/entity/enemy"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	b := enemies.NewBandit(100, nil, nil)
	k := characters.NewKnight(100, nil, nil)

	team := []character.Character{k}
	opponentTeam := []enemy.Enemy{b}

	initGui(team, opponentTeam, area.Plains)
	//combat.InitBattle(team, opponentTeam)
}

func initGui(characters []character.Character, enemies []enemy.Enemy, area area.AreaId) {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Gioco")

	ebiten.SetFullscreen(true)
	ebiten.RunGame(gui.NewGame(
		battlescene.NewBattleScene(characters, enemies, area),
	))
}
