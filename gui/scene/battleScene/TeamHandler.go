package battlescene

import (
	"github.com/KONshougun/K-game/gui/component"
	"github.com/KONshougun/K-game/model/entity"

	"github.com/hajimehoshi/ebiten/v2"
)

type TeamView struct {
	*component.Component
	TeamAvatar []*EntityView
}

func NewTeamView(team []*entity.EntityData, area component.Area) *TeamView {
	if area.Height < area.Width/2 {
		panic("Lunghezza dell'area non corretta")
	}
	t := &TeamView{
		Component:  component.NewComponent(area),
		TeamAvatar: make([]*EntityView, len(team)),
	}
	for i := range team {
		switch team[i].EntityType {
		case entity.Character:
			t.TeamAvatar[i] = NewEntityView(team[i], component.Area{
				X:      area.X + area.Width/4*float64(i),
				Y:      area.Y,
				Width:  area.Width / 4,
				Height: area.Width / 2,
			})
		case entity.Enemy:
			t.TeamAvatar[i] = NewEntityView(team[i], component.Area{
				X:      area.X + area.Width - area.Width/4*float64(i+1),
				Y:      area.Y,
				Width:  area.Width / 4,
				Height: area.Width / 2,
			})
		}
	}
	return t
}

func (t *TeamView) Update() error {
	return nil
}
func (t *TeamView) Draw(screen *ebiten.Image) {
	for _, a := range t.TeamAvatar {
		a.Draw(screen)
	}
}
