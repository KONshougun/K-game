package component

import "github.com/hajimehoshi/ebiten/v2"

type ComponentController interface {
	Update() error
	Draw(screen *ebiten.Image, width, height float64)
}

type Area struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

type Component struct {
	ComponentController
	Area
}

func NewComponent(controller ComponentController, area Area) Component {
	return Component{
		ComponentController: controller,
		Area:                area,
	}
}
