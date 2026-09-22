package component

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type BarType byte

const (
	HealthBar BarType = iota
	ShieldBar
	StaminaBar
	ManaBar
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
	BarType
}

func NewProgressBar(value, maxValue *float64, t BarType, border BorderOptions, area Area) *ProgressBar {
	if *value > *maxValue {
		panic("Pozza ingroppase")
	}
	return &ProgressBar{
		Component:       NewComponent(area),
		CurrValue:       value,
		MaxValue:        maxValue,
		BackgroundColor: color.Black,
		BarType:         t,
		Border:          border,
	}
}

func (b *ProgressBar) Update() error {
	return nil
}
func (b *ProgressBar) Draw(screen *ebiten.Image) {
	ratio := *b.CurrValue / *b.MaxValue
	if ratio > 1 {
		ratio = 1
	}
	b.BarColor = GetBarColor(b.BarType, ratio)
	// Sfondo
	vector.FillRect(
		screen,
		float32(b.X),
		float32(b.Y),
		float32(b.Width),
		float32(b.Height),
		b.BackgroundColor,
		false,
	)

	// Parte piena
	vector.FillRect(
		screen,
		float32(b.X),
		float32(b.Y),
		float32(b.Width*ratio),
		float32(b.Height),
		b.BarColor,
		false,
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

func GetBarColor(barT BarType, value float64) color.NRGBA {
	switch barT {
	case HealthBar:
		return ThreeColorLerp(
			color.NRGBA{0, 200, 0, 255},
			color.NRGBA{255, 255, 0, 255},
			color.NRGBA{255, 0, 0, 255},
			value,
		)
	case ShieldBar:
		return color.NRGBA{R: 245, G: 245, B: 245, A: 192}
	case StaminaBar:
		return ThreeColorLerp(
			color.NRGBA{255, 210, 50, 255},
			color.NRGBA{255, 140, 0, 255},
			color.NRGBA{255, 0, 0, 255},
			value,
		)
	case ManaBar:
		return ThreeColorLerp(
			color.NRGBA{0, 200, 255, 255},
			color.NRGBA{0, 0, 255, 255},
			color.NRGBA{150, 0, 255, 255},
			value,
		)
	default:
		panic("Daje")
	}
}

// COLOR
func LerpColor(a, b color.NRGBA, t float64) color.NRGBA {
	return color.NRGBA{
		R: uint8(float64(a.R) + (float64(b.R)-float64(a.R))*t),
		G: uint8(float64(a.G) + (float64(b.G)-float64(a.G))*t),
		B: uint8(float64(a.B) + (float64(b.B)-float64(a.B))*t),
		A: (a.A + b.A) / 2,
	}
}
func ThreeColorLerp(a, b, c color.NRGBA, value float64) color.NRGBA {
	value = 1 - value

	// Curva esponenziale
	value = math.Pow(value, 2.0)

	if value <= 0.5 {
		return LerpColor(a, b, value/0.5)
	}

	return LerpColor(b, c, (value-0.5)/0.5)
}
