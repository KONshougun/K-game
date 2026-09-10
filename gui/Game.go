package gui

import (
	"gioco/gui/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	//Layout(outsideWidth, outsideHeight int) (int, int)
}

const (
	GameWidth  = 1280
	GameHeight = 720
)

type Game struct {
	Scene Scene
	l     component.Label
}

//FUNZIONE DI TESTING
func NewGame()Game{
	return Game{l: component.NewLabel("Hello world", component.NormalSize, component.NoneSpace, text.AlignCenter, text.AlignCenter, component.Area{
		X: 0,Y: 0,Width: GameWidth,Height: GameHeight,
	})}
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
