package guihandler

import (
	"fmt"

	"github.com/KONshougun/K-game/model/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

func GetAvatar(name string, entityT entity.EntityType) *ebiten.Image {
	return LoadImage(fmt.Sprintf("asset/entity/%s/%s/%s_Avatar.png", entityT.String(), name, name))
}
func GetIcon(name string, entityT entity.EntityType) *ebiten.Image {
	return LoadImage(fmt.Sprintf("asset/entity/%s/%s/%s_Icon.png", entityT.String(), name, name))
}
func GetAbilities(name string, entityT entity.EntityType, l int) []*ebiten.Image {
	abilities := make([]*ebiten.Image, l)
	for i := range l {
		//fmt.Printf("asset/entity/%s/%s/%s_Ability_%v.png\n", entityT.String(), name, name, (i + 1))
		abilities[i] = LoadImage(fmt.Sprintf("asset/entity/%s/%s/%s_Ability_%v.png", entityT.String(), name, name, (i + 1)))
	}
	return abilities
}
