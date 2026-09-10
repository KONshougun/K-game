package battlescene

import (
	"github.com/KONshougun/K-game/game/comm/area"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	area   *area.Area
	sprite *ebiten.Image
}

func NewBackground(area *area.Area) *Background {
	return &Background{
		area:   area,
		sprite: nil,
	}
}

func (b *Background) Update() error {

	return nil
}
func (b *Background) Draw(screen *ebiten.Image) {

}
