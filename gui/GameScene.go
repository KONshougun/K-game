package gui

import (
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
}

// FUNZIONE DI TESTING
func NewGame(scene Scene) *Game {
	return &Game{Scene: scene}
}

func (g *Game) Update() error {
	return g.Scene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GameWidth, GameHeight
}
