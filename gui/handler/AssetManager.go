package handler

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"github.com/KONshougun/K-game/game/comm/area"
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
		panic(fmt.Sprintf("Immagine non trovata: %s", path))
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}
func GetCombatBackground(area area.AreaId) *ebiten.Image {
	return LoadImage(fmt.Sprintf("asset/background/area/%s.png", area.String()))
}
