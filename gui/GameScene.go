package gui

import (
	"github.com/KONshougun/K-game/game/characters"
	"github.com/KONshougun/K-game/game/combat/effect"
	"github.com/KONshougun/K-game/game/combat/event"
	"github.com/KONshougun/K-game/game/combat/handler"
	"github.com/KONshougun/K-game/gui/component"
	"github.com/KONshougun/K-game/model/stat"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)
}

const (
	GameWidth  = 1280
	GameHeight = 720
)

type Game struct {
	Scene Scene

	e *EntityView
}

// FUNZIONE DI TESTING
func NewGame() *Game {
	c := characters.NewKnight(100, nil, nil).EntityData
	c.TriggerEffects(event.SetEntityData, effect.None, &handler.SetEntityData{
		FlatStats:    stat.Stats{},
		PercentStats: stat.Stats{},
	})
	return &Game{
		e: NewEntityView(c, component.Area{
			X:      0,
			Y:      0,
			Width:  GameWidth /10,
			Height: GameHeight/5,
		}),
	}
}

func (g *Game) Update() error {
	return g.e.Update()
	//return g.Scene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.e.Draw(screen)
	//g.Scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GameWidth, GameHeight
}
