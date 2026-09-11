package component

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type BorderOptions struct {
	BorderColor color.Color
	BorderWidth float64
}

type ProgressBar struct {
	*Component

	CurrValue *float64
	MaxValue  *float64

	Border          BorderOptions
	BackgroundColor color.Color
	BarColor        color.Color
	Antialias       bool
}

func NewProgressBar(value, maxValue *float64, backgroundColor, barColor color.Color,
	antialias bool, border BorderOptions, area Area) *ProgressBar {
	if *value > *maxValue {
		panic("Pozza ingroppase")
	}
	return &ProgressBar{
		Component:       NewComponent(area),
		CurrValue:       value,
		MaxValue:        maxValue,
		BackgroundColor: backgroundColor,
		BarColor:        barColor,
		Border:          border,
		Antialias:       antialias,
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
		b.BackgroundColor,
		b.Antialias,
	)

	// Parte piena
	vector.FillRect(
		screen,
		float32(b.X),
		float32(b.Y),
		float32(b.Width*ratio),
		float32(b.Height),
		b.BarColor,
		b.Antialias,
	)
	
	//Bordo
	vector.StrokeRect(
		screen,
		float32(b.X),
		float32(b.Y),
		float32(b.Width),
		float32(b.Height),
		float32(b.Border.BorderWidth),
		b.Border.BorderColor,
		false,
	)
}
