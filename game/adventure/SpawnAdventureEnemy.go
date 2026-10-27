package adventure

import (
	"math/rand"

	"github.com/KONshougun/K-game/game/enemies"
	"github.com/KONshougun/K-game/model/area"
	"github.com/KONshougun/K-game/model/entity/enemy"
)

func SpawnAdventureEnemy(adv *Adventure) []enemy.Enemy {

	if adv.Inventory.Day == 0 {
		return []enemy.Enemy{enemies.NewBandit(1, nil, nil)}
	}
	switch adv.Inventory.Area.Id {
	case area.Plains:
		r := rand.Intn(100)
		if r < 30 {

		}
	default:
		panic("Che cazzo gli hai passato??")
	}

	return nil
}
