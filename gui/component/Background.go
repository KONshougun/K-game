package component

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	*Component

	Image *ebiten.Image
}

func NewBackground(img *ebiten.Image, area Area) *Background {
	return &Background{
		Component: NewComponent(area),
		Image:     img,
	}
}

func (b *Background) Update() error {
	return nil
}
func (b *Background) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}

	imgBounds := b.Image.Bounds()
	scaleX := b.Area.Width / float64(imgBounds.Dx())
	scaleY := b.Area.Height / float64(imgBounds.Dy())
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(b.Area.X, b.Area.Y)
	screen.DrawImage(b.Image, &opts)
}
