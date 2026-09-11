package component

/*type Controller interface {
	Update()
	Draw(screen *ebiten.Image)
}*/

type Area struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

type Component struct {
	//Controller
	Area
}

func NewComponent( /*controller Controller,*/ area Area) *Component {
	return &Component{
		//Controller: controller,
		Area: area,
	}
}
