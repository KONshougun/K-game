package guihandler

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"github.com/KONshougun/K-game/model/area"
	"github.com/hajimehoshi/ebiten/v2"
)

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
