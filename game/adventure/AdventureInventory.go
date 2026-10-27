package adventure

import (
	"github.com/KONshougun/K-game/model"
	"github.com/KONshougun/K-game/model/area"
	"github.com/KONshougun/K-game/model/entity/character"
	"github.com/KONshougun/K-game/model/equipment/armor"
	"github.com/KONshougun/K-game/model/equipment/weapon"
)

type AdventureInventory struct {
	Characters    [character.CharacterCount]character.Character
	WeaponInv     []*weapon.Weapon
	ArmorInv      []*armor.ArmorPiece
	ItemInv       [model.ItemCount]int
	UniqueItemInv [model.UniqueItemCount]bool

	WorldLevel, WorldCore, Island, Day int

	Area area.Area
}

func NewAdvetureInventory() *AdventureInventory {
	return &AdventureInventory{
		WeaponInv: make([]*weapon.Weapon, 10),
		ArmorInv:  make([]*armor.ArmorPiece, 20),
		WorldCore: 1,
		Island:    1,
		Area:      area.Areas[area.Plains],
	}
}
