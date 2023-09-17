package gjfw

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

var DEFAULT_COLOR = color.RGBA{255, 255, 0, 255}

type Element struct {
	X, Y, W, H float32
	Color      color.Color
	//fixme use rectangle or similr
}

func NewSquare(x, y, s float32) *Element {
	e := &Element{X: x, Y: y, W: s, H: s, Color: DEFAULT_COLOR}
	return e
}
func NewSquareC(x, y, s float32, c color.Color) *Element {
	e := &Element{X: x, Y: y, W: s, H: s, Color: c}
	return e
}

func NewElementRectC(x, y, w, h float32, c color.Color) *Element {
	e := &Element{X: x, Y: y, W: w, H: h, Color: c}
	return e
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
