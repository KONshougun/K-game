package main

import (
	"gioco/game/characters"
	"gioco/game/combat"
	"gioco/game/comm/area"
	"gioco/game/enemies"
	"gioco/model/entity"
)

func main() {
	b := enemies.NewBandit(100, nil, nil)
	k := characters.NewKnight(100, nil, nil)
	combat.Fight([]*entity.Entity{k.Entity}, []*entity.Entity{b.Entity}, area.Area{Id: area.Plains})
}
