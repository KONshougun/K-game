package gui

import (
	"github.com/KONshougun/K-game/gui/component/container"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameStruct struct {
	scenes []container.Scene
	data   guihandler.Data
}

var Game GameStruct

func NewGame(scene container.Scene) {
	Game.scenes = append(Game.scenes, scene)

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Gioco")

	ebiten.SetFullscreen(true)
	ebiten.RunGame(&Game)
}

func (g *GameStruct) Update() error {
	return g.scenes[0].Update()
}

func (g *GameStruct) Draw(screen *ebiten.Image) {
	g.scenes[0].Draw(screen)
}
func (g *GameStruct) GetSceneData() *container.SceneData {
	return g.scenes[0].GetSceneData()
}

func (g *GameStruct) AddScene(scene container.Scene) {
	if scene != nil {
		g.scenes = append([]container.Scene{scene}, g.scenes...)
	}
}

func (g *GameStruct) RemoveScene(sceneData guihandler.Data) {
	g.data = sceneData
	g.scenes = append(g.scenes[1:], nil)
}
func (g *GameStruct) GetData() guihandler.Data {
	app := g.data
	g.data = nil
	return app
}
func (g *GameStruct) SetData(data guihandler.Data) {
	g.data = data
}
func (g *GameStruct) Layout(outsideWidth, outsideHeight int) (int, int) {
	return container.GameWidth, container.GameHeight
}
