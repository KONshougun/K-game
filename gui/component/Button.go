package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ButtonState byte

const (
	Normal ButtonState = iota
	MouseOver
	Pressed
	Locked

	CountBtnState
)

type Button struct {
	*Component
	*Label

	Locked bool

	State [CountBtnState]*ebiten.Image

	OnClick func()
}

func NewButton(states [CountBtnState]*ebiten.Image, label *Label, area Area, f func()) *Button {
	return &Button{
		Component: NewComponent(area),
		Label:     label,
		State:     states,
		OnClick:   f,
	}
}

func (b *Button) Update() error {
	if b.Locked {
		return nil
	}

	mouseX, mouseY := ebiten.CursorPosition()

	if b.contains(float64(mouseX), float64(mouseY)) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if b.OnClick != nil {
				b.OnClick()
			}
		}
	}

	if b.Label != nil {
		if err := b.Label.Update(); err != nil {
			return err
		}
	}
	return nil
}

func (b *Button) Draw(screen *ebiten.Image) {

	if b.State != [4]*ebiten.Image{nil, nil, nil, nil} {
		img := b.State[b.getState()]
		if img == nil {
			img = b.State[Normal]
		}
		imgBounds := img.Bounds()
		scaleX := float64(b.Area.Width) / float64(imgBounds.Dx())
		scaleY := float64(b.Area.Height) / float64(imgBounds.Dy())

		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Scale(scaleX, scaleY)
		opts.GeoM.Translate(
			float64(b.Area.X),
			float64(b.Area.Y),
		)
		screen.DrawImage(img, &opts)
	}

	if b.Label != nil {
		b.Label.Draw(screen)
	}
}
func (b *Button) getState() ButtonState {
	if b.Locked {
		return Locked
	}

	mouseX, mouseY := ebiten.CursorPosition()

	if !b.contains(float64(mouseX), float64(mouseY)) {
		return Normal
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return Pressed
	}

	return MouseOver
}
func (b *Button) contains(x, y float64) bool {
	return x >= b.X &&
		x <= b.X+b.Width &&
		y >= b.Y &&
		y <= b.Y+b.Height
}
