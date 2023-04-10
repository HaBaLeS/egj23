package gjfw

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Element struct {
	X, Y, W, H float32
	Color      color.Color
	//fixme use rectangle or similr
}

func (pe *Element) Render(surface *ebiten.Image) {
	vector.DrawFilledRect(surface, pe.X, pe.Y, pe.W, pe.H, pe.Color, true)
}

func (pe *Element) Intersect(other *Element) bool {
	if pe.X < other.X+other.W &&
		pe.X+pe.W > other.X &&
		pe.Y < other.Y+other.H &&
		pe.H+pe.Y > other.Y {
		return true
	} else {
		return false
	}
}
