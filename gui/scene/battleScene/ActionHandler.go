package battlescene

import (
	"github.com/KONshougun/K-game/model/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type ActionMenu struct {

}

func NewActionMenu(entity *entity.AbilityData) *ActionMenu{
	return &ActionMenu{}
}

func (m *ActionMenu) Update() {

}
func (m *ActionMenu) Draw(screen *ebiten.Image) {

}
