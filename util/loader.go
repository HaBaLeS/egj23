package util

import (
	"github.com/habales/egj23/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadImage(file string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFileSystem(assets.Data, file)
	if err != nil {
		panic(err)
	}
	return img
}
