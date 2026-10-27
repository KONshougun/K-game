package component

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

/*type Controller interface {
	Update()
	Draw(screen *ebiten.Image)
}*/

type Component interface {
	Update() error
	Draw(screen *ebiten.Image)

	GetComponent() *ComponentData
}

type Area struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (a Area) Contains(other Area) bool {
	if !(other.X >= a.X &&
		other.Y >= a.Y &&
		other.X+other.Width <= a.X+a.Width &&
		other.Y+other.Height <= a.Y+a.Height) {
			
		if !(other.X >= a.X) {
			fmt.Println("vsdvdv")
			fmt.Printf("other.X: %v\n", other.X)
			fmt.Printf("a.X: %v\n", a.X)
		}

		if !(other.Y >= a.Y) {
			fmt.Println("ca")
			fmt.Printf("other.X: %v\n", other.Y)
			fmt.Printf("a.X: %v\n", a.Y)
		}
		if !(other.X+other.Width <= a.X+a.Width) {
			fmt.Println("da")
			fmt.Printf("other.X: %v\n", other.X+other.Width)
			fmt.Printf("a.X: %v\n", a.X+a.Width)
		}
		if !(other.Y+other.Height <= a.Y+a.Height){
			fmt.Println("aus")
			fmt.Printf("other.X: %v\n", other.Y+other.Height)
			fmt.Printf("a.X: %v\n", a.Y+a.Height)
		}
	}
	return other.X >= a.X &&
		other.Y >= a.Y &&
		other.X+other.Width <= a.X+a.Width &&
		other.Y+other.Height <= a.Y+a.Height
}

type ComponentData struct {
	Id uint8
	Area
}

func NewComponentData( /*controller Controller,*/ area Area) *ComponentData {
	return &ComponentData{
		//Controller: controller,
		Area: area,
	}
}
