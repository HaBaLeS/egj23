package util

import (
	"github.com/habales/egj23/assets"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func LoadFont(file string, size float64) font.Face {
	data, err := assets.Data.ReadFile(file)
	if err != nil {
		panic(err)
	}
	tt, err := opentype.Parse(data)
	if err != nil {
		panic(err)
	}
	const dpi = 72
	fontFact, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}
	return fontFact
}
