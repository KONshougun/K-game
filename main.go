package main

import (
	"gioco/game/characters"
	"gioco/game/combat"
	"gioco/game/enemies"
	"gioco/gui"
	"gioco/model/entity/character"
	"gioco/model/entity/enemy"

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
	ebiten.RunGame(&gui.Game{
		Scene: gui.NewGame().Scene,
	})
}
