package handler

import (
	"bytes"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

/*type AssetManager struct {
	images map[string]*ebiten.Image
}

func (a *AssetManager) Get(path string) *ebiten.Image {
	if img, ok := a.images[path]; ok {
		return img
	}

	img, err := loadImage(path)
	if err != nil {
		panic("Daje")
	}
	a.images[path] = img

	return img
}*/

func LoadImage(path string) *ebiten.Image /*, error*/ {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil //, err
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil //, err
	}
	return ebiten.NewImageFromImage(img) //, nil
}
