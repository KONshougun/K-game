package handler

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"github.com/KONshougun/K-game/model/entity"
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

func LoadImage(path string) *ebiten.Image {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil
	}
	return ebiten.NewImageFromImage(img)
}
func GetAvatar(name string, entityT entity.EntityType) *ebiten.Image {
	return LoadImage(fmt.Sprintf("asset/entity/%s/avatar/%s.png", entityT.String(), name))
}
