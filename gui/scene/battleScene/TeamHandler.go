package battlescene

import (
	"github.com/KONshougun/K-game/gui"
	"github.com/KONshougun/K-game/gui/component"

	"github.com/hajimehoshi/ebiten/v2"
)

type TeamView struct {
	*component.Component
	Entities []*gui.EntityView
}

func NewTeamView() *TeamView {
	return &TeamView{}
}

func (t *TeamView) Update() error {
	return nil
}
func (t *TeamView) Draw(screen *ebiten.Image) {

}
