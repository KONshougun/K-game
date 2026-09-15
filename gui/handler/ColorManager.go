package handler

import "image/color"

func LerpColor(a, b color.RGBA, t float64) color.RGBA {
	return color.RGBA{
		R: uint8(float64(a.R) + (float64(b.R)-float64(a.R))*t),
		G: uint8(float64(a.G) + (float64(b.G)-float64(a.G))*t),
		B: uint8(float64(a.B) + (float64(b.B)-float64(a.B))*t),
		A: 255,
	}
}
func ThreeColorLerp(a, b, c color.RGBA, value float64) color.RGBA {
	if value <= 50 {
		return LerpColor(a, b, value/50)
	}

	return LerpColor(b, c, (value-50)/50)
}
