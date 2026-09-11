package gui

import (
	"image/color"

	"github.com/KONshougun/K-game/gui/component"
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
	p     component.ProgressBar
}

// FUNZIONE DI TESTING
func NewGame() *Game {
	a := 4.0
	b := 10.0
	return &Game{p: *component.NewProgressBar(&a, &b, color.RGBA{50, 50, 50, 255}, color.RGBA{200, 50, 50, 255},
		false, component.BorderOptions{
			BorderColor: color.White,
			BorderWidth: 2,
		}, component.Area{
			X:      GameWidth/2 - GameWidth/8,
			Y:      GameHeight/2 - GameHeight/8,
			Width:  GameWidth / 4,
			Height: GameHeight / 4,
		})}
}

func (g *Game) Update() error {
	if *g.p.CurrValue == 0 {
		*g.p.CurrValue = *g.p.MaxValue
	} else {
		*g.p.CurrValue--
	}
	g.p.Update()
	return nil
	//return g.Scene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.p.Draw(screen)
	//g.Scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GameWidth, GameHeight
}
