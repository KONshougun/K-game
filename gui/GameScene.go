package gui

import (
	"image/color"

	"github.com/KONshougun/K-game/gui/component"
	"github.com/KONshougun/K-game/gui/handler"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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

	b *component.Button
}

// FUNZIONE DI TESTING
func NewGame() *Game {
	area := component.Area{
		X:      GameWidth/2 - GameWidth/8,
		Y:      GameHeight/2 - GameHeight/8,
		Width:  GameWidth / 4,
		Height: GameHeight / 4,
	}
	return &Game{b: component.NewButton([component.CountBtnState]*ebiten.Image{
		handler.LoadImage("assets/ui/component/button/button-up.png"),
		handler.LoadImage("assets/ui/component/button/button-up.png"),
		handler.LoadImage("assets/ui/component/button/button-down.png"),
		handler.LoadImage("assets/ui/component/button/button-up.png"),
	}, component.NewLabel("Hello", color.Black, component.VeryLargeSize, component.NoneSpace,
		text.AlignCenter, text.AlignCenter, area), area, func() {})}
}

func (g *Game) Update() error {
	g.b.Update()
	return nil
	//return g.Scene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.b.Draw(screen)
	//g.Scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GameWidth, GameHeight
}
