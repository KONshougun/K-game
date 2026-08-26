package main

import (
	"gioco/game/characters"
	"gioco/game/combat"
	"gioco/game/comm/area"
	"gioco/game/enemies"
	"gioco/model/entity/character"
	"gioco/model/entity/enemy"
)

func main() {
	b := enemies.NewBandit(100, nil, nil)
	k := characters.NewKnight(100, nil, nil)
	k.AwakenAbility(1)
	k.AwakenAbility(2)
	k.AwakenAbility(3)

	team := []character.Character{k}
	opponentTeam := []enemy.Enemy{b}

	combat.Fight(team, opponentTeam, area.Area{Id: area.Plains})
}
