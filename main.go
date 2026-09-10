package main

import (
 "github.com/KONshougun/K-game/game/characters"
 "github.com/KONshougun/K-game/game/combat"
 "github.com/KONshougun/K-game/game/enemies"
 "github.com/KONshougun/K-game/gui"
 "github.com/KONshougun/K-game/model/entity/character"
 "github.com/KONshougun/K-game/model/entity/enemy"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	initGui()

	b := enemies.NewBandit(100, nil, nil)
	k := characters.NewKnight(100, nil, nil)

	team := []character.Character{k}
	opponentTeam := []enemy.Enemy{b}

	combat.InitBattle(team, opponentTeam)
}

func initGui() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Gioco")

	ebiten.SetFullscreen(true)
	ebiten.RunGame(gui.NewGame())
}
