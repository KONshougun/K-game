package guihandler

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func IsMousePressed(mouseBtn ebiten.MouseButton, area *Area) bool {

	if ebiten.IsMouseButtonPressed(mouseBtn) {
		if area != nil {
			area.IsIn(ebiten.CursorPositionF())
		}
		return true
	}
	return false
}
