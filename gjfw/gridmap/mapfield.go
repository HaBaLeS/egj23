package gridmap

import (
	"github.com/habales/egj23/egmath"
	"image/color"
	"math"
)

type MapField struct {
	pos          egmath.Vec2
	overlayColor color.Color
	walkCost     int
}

func NewMapField(x, y int, c color.Color) *MapField {
	mf := &MapField{
		pos:          egmath.Vec2{x, y},
		overlayColor: c,
		walkCost:     math.MaxInt, //fixme constants here
	}
	return mf
}
