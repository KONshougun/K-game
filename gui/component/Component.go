package component

/*type ComponentController interface {
	Update() error
	Draw(screen *ebiten.Image)
}*/

type Area struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

type Component struct {
	Area
}

func NewComponent( /*controller ComponentController,*/ area Area) *Component {
	return &Component{
		//ComponentController: controller,
		Area: area,
	}
}
