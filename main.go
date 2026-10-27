package main

import (
	"github.com/KONshougun/K-game/gui"
	startgame "github.com/KONshougun/K-game/gui/scene/startGame"
)

func main() {
	initGui()
}

func initGui() {
	gui.NewGame(startgame.NewStartGameScene())
}
