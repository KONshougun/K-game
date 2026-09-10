package component

import (
	"gioco/gui/handler"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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
	NoneSpace LineSpacing = iota
	SmallSpace
	NormalSpace
	LargeSpace
	VeryLargeSpace
)

type Label struct {
	Component

	Text   string
	Face   text.Face
	Layout text.DrawOptions
}

func NewLabel(str string, size Size, space LineSpacing, horizontalAlign, verticalAlign text.Align, area Area) Label {
	return Label{
		Component: NewComponent(&Label{}, area),

		Text: str,
		Face: handler.LoadFont(handler.NotoSerif, float64(size)),
		Layout: text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				LineSpacing:    float64(space),
				PrimaryAlign:   horizontalAlign,
				SecondaryAlign: verticalAlign,
			},
		},
	}
}

func (l *Label) Update() error {
	return nil
}
func (l *Label) Draw(screen *ebiten.Image, width, height float64) {
	str := l.fitText(width, height)

	text.Draw(
		screen,
		str,
		l.Face,
		&l.Layout,
	)
}

func (l *Label) fitText(width, height float64) string {
	str := []rune(l.Text)
	suffix := "..."

	fits := func(s string) bool {
		w, h := text.Measure(
			s,
			l.Face,
			l.Layout.LineSpacing,
		)

		return w <= width && h <= height
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

	return ""
}
