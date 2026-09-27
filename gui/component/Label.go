package component

import (
	"bytes"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
)

type Size byte

const (
	VerySmallSize Size = 8
	SmallSize     Size = 12
	NormalSize    Size = 16
	LargeSize     Size = 24
	VeryLargeSize Size = 32
)

type LineSpacing byte

const (
	NoneSpace      LineSpacing = 0
	SmallSpace     LineSpacing = 8
	NormalSpace    LineSpacing = 16
	LargeSpace     LineSpacing = 24
	VeryLargeSpace LineSpacing = 32
)

type Label struct {
	*Component

	Text   string
	Face   text.Face
	Layout *text.DrawOptions
}

func NewLabel(str string, c color.Color, size Size, space LineSpacing, horizontalAlign, verticalAlign text.Align, area Area) *Label {
	return &Label{
		Component: NewComponent(area),

		Text: str,
		Face: LoadFont(NotoSerif, float64(size)),
		Layout: &text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				LineSpacing:    float64(space),
				PrimaryAlign:   horizontalAlign,
				SecondaryAlign: verticalAlign,
			},
			DrawImageOptions: ebiten.DrawImageOptions{
				ColorScale: colorScale(c),
			},
		},
	}
}

func (l *Label) Update() error {
	return nil
}
func (l *Label) Draw(screen *ebiten.Image) {
	str := l.fitText()
	l.Layout.GeoM.Reset()
	l.Layout.GeoM.Translate(l.X+l.Width/2*float64(l.Layout.PrimaryAlign), l.Y+l.Height/2*float64(l.Layout.SecondaryAlign))
	text.Draw(screen, str, l.Face, l.Layout)
}

/*func (l *Label) fitText() string {
	str := []rune(l.Text)
	suffix := "..."

	fits := func(s string) bool {
		w, h := text.Measure(
			s,
			l.Face,
			l.Layout.LineSpacing,
		)

		return w <= l.Width && h <= l.Height
	}

	if fits(l.Text) {
		return l.Text
	}

	for len(str) > 0 {
		candidate := string(str) + suffix

		if fits(candidate) {
			return candidate
		}

		str = str[:len(str)-1]
	}

	if fits(suffix) {
		return suffix
	}

	fmt.Println("ciao")
	return ""
}*/

func (l *Label) fitText() string {
	const suffix = "..."

	fits := func(s string) bool {
		w, h := text.Measure(
			s,
			l.Face,
			l.Layout.LineSpacing,
		)

		return w <= l.Width && h <= l.Height
	}

	if fits(l.Text) {
		return l.Text
	}

	str := []rune(l.Text)

	for len(str) > 0 {
		// Evita che il testo finisca con un newline prima di "..."
		if str[len(str)-1] == '\n' {
			str = str[:len(str)-1]
		}

		candidate := string(str) + suffix

		if fits(candidate) {
			return candidate
		}

		str = str[:len(str)-1]
	}

	if fits(suffix) {
		return suffix
	}

	return ""
}

func colorScale(c color.Color) ebiten.ColorScale {
	var scale ebiten.ColorScale
	scale.ScaleWithColor(c)
	return scale
}

//FONT

type Font string

const (
	NotoSerif = "asset/ui/font/NotoSerif.ttf"
)

func LoadFont(font Font, size float64) text.Face {
	data, err := os.ReadFile(string(font))
	if err != nil {
		panic(err)
	}

	source, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	return &text.GoTextFace{
		Source:   source,
		Language: language.Italian,
		Size:     size,
	}
}
