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

type GameStruct struct {
	Scenes []Scene
}

var Game GameStruct

func NewGame(scene Scene) {
	Game.Scenes = append(Game.Scenes, scene)

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Gioco")

	ebiten.SetFullscreen(true)
	ebiten.RunGame(&Game)
}

func (g *GameStruct) Update() error {
	return g.Scenes[len(g.Scenes)-1].Update()
}

func (g *GameStruct) Draw(screen *ebiten.Image) {
	g.Scenes[len(g.Scenes)-1].Draw(screen)
}

func (g *GameStruct) Layout(outsideWidth, outsideHeight int) (int, int) {
	return GameWidth, GameHeight
}
