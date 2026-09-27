package adventure

import (
	"github.com/KONshougun/K-game/game/characters"
	"github.com/KONshougun/K-game/model/entity/character"
	"github.com/KONshougun/K-game/model/equipment/weapon"
)

type Adventure struct {
	Inventory *AdventureInventory
}

func InitAdventure() *Adventure {
	a := Adventure{
		Inventory: NewAdvetureInventory(),
	}
	a.Inventory.Characters[character.Knight] = characters.NewKnight(1, weapon.NewWeapon(weapon.Sword, weapon.Std), nil)
	a.Inventory.WeaponInv[0] = a.Inventory.Characters[0].GetCharacter().Weapons[0]

	a.Inventory.ItemInv[Monete] = 50

	return &a
}
