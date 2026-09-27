package adventure

import (
	"github.com/KONshougun/K-game/model/entity/character"
	"github.com/KONshougun/K-game/model/equipment/armor"
	"github.com/KONshougun/K-game/model/equipment/weapon"
)

type AdventureInventory struct {
	Characters [character.CharacterCount]character.Character
	WeaponInv []*weapon.Weapon
	ArmorInv []*armor.ArmorPiece
	ItemInv [ItemCount]int
	UniqueItemInv [UniqueItemCount]bool

	WorldLevel, WorldCore, Area, Island, Day int
}

func NewAdvetureInventory()*AdventureInventory{
	return &AdventureInventory{
		WeaponInv: make([]*weapon.Weapon, 10),
		ArmorInv: make([]*armor.ArmorPiece, 20),
		WorldCore: 1,
		Island: 1,
		Area: 1,
	}
}