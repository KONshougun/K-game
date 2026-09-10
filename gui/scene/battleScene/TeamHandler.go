package battlescene

import (
	"gioco/model/entity"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type TeamView struct {
	team []*Entity
	area image.Rectangle
}
type Entity struct {
	entityData *entity.EntityData
	sprite     *ebiten.Image
}

func NewTeamView(team []*entity.EntityData) *TeamView {
	return &TeamView{team: NewEntities(team...)}
}
func NewEntities(entities ...*entity.EntityData) []*Entity {
	team := make([]*Entity, len(entities))

	for i := range entities {
		team[i] = &Entity{
			entityData: entities[i],
			sprite:     nil,
		}
	}

	return team
}

func (t *TeamView) Update() {

}
func (t *TeamView) Draw(screen *ebiten.Image) {

}
func (t *TeamView) Layout(area image.Rectangle) {
	t.area = area
}
