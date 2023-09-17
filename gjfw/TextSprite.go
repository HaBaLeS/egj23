package gjfw

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image"
	"image/color"
)

//fixme can ewe extract a drawable ?!?!?

type TextSprite struct {
	x, y     int
	text     string
	fontFace font.Face
	color    color.Color
	dirty    bool //do we need to calc bounds
	bounds   image.Rectangle
}

func NewTextSprite(face font.Face, text string, col color.Color) *TextSprite {
	return &TextSprite{
		fontFace: face,
		color:    col,
		text:     text,
		dirty:    true,
	}
}

func (ts *TextSprite) Render(surface *ebiten.Image) {
	text.Draw(surface, ts.text, ts.fontFace, ts.x, ts.y, ts.color)
}

func (ts *TextSprite) SetPos(x, y int) {
	ts.x = x
	ts.y = y
}

func (ts *TextSprite) SetText(text string) {
	ts.text = text
	ts.dirty = true
}

func (ts *TextSprite) Bounds() image.Rectangle {
	if ts.dirty { //cache the calc
		ts.bounds = text.BoundString(ts.fontFace, ts.text)
	}
	return ts.bounds
}
