package main

import (
	"gioco/game/characters"
	"gioco/game/combat"
	"gioco/game/enemies"
	"gioco/model/entity/character"
	"gioco/model/entity/enemy"
)

func main() {
	b := enemies.NewBandit(100, nil, nil)
	k := characters.NewKnight(100, nil, nil)

	team := []character.Character{k}
	opponentTeam := []enemy.Enemy{b}

	combat.InitBattle(team, opponentTeam)
}
