package component

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ProgressBar struct {
	Component

	CurrValue *float64
	MaxValue  *float64

	Antialias bool
}

func NewProgressBar(value, maxValue *float64, antialias bool, area Area) *ProgressBar {
	return &ProgressBar{
		Component: *NewComponent(area),
		CurrValue: value,
		MaxValue:  maxValue,
		Antialias: antialias,
	}
}

func (b *ProgressBar) Update() error {
	return nil
}
func (b *ProgressBar) Draw(screen *ebiten.Image) {
	ratio := *b.CurrValue / *b.MaxValue

	// Sfondo
	vector.FillRect(
		screen,
		float32(b.X),
		float32(b.Y),
		float32(b.Width),
		float32(b.Height),
		color.RGBA{50, 50, 50, 255},
		b.Antialias,
	)

	// Parte piena
	vector.FillRect(
		screen,
		float32(b.X),
		float32(b.Y),
		float32(b.Width*ratio),
		float32(b.Height),
		color.RGBA{200, 50, 50, 255},
		b.Antialias,
	)
}
