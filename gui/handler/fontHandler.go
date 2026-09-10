package handler

import (
	"bytes"
	"os"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
)

type Font string

const (
	NotoSerif = "github.com/KONshougun/K-game/assets/ui/font/NotoSerif.ttf"
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
