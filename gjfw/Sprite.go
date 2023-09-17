package gjfw

import (
	"github.com/habales/egj23/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	x, y, w, h, sx, sy float64
	rot                float64
	img                *ebiten.Image
	op                 *ebiten.DrawImageOptions
}

func NewSprite(file string) *Sprite {
	img := util.LoadImage(file)
	sp := &Sprite{
		img: img,
		w:   float64(img.Bounds().Dx()),
		h:   float64(img.Bounds().Dy()),
		op:  &ebiten.DrawImageOptions{},
	}
	return sp
}

func (sp *Sprite) Clone() *Sprite {
	panic("Not implemented")
}

func (sp *Sprite) SetPos(x, y int) {
	sp.x = float64(x)
	sp.y = float64(y)
	sp.op.GeoM.Reset()
	if sp.rot != 0 {
		sp.op.GeoM.Translate(-sp.w/2, -sp.h/2)
		sp.op.GeoM.Rotate(sp.rot)
		sp.op.GeoM.Translate(sp.w/2, sp.h/2)
	}
	sp.op.GeoM.Scale(sp.sx, sp.sy)
	sp.op.GeoM.Translate(sp.x, sp.y)
}

func (sp *Sprite) SetSize(w, h float64) {
	sp.sx = (100.0 / sp.w * w) / 100
	sp.sy = (100.0 / sp.h * h) / 100
	sp.op.GeoM.Scale(sp.sx, sp.sy)
}

func (sp *Sprite) SetAlpha(a float32) {
	sp.op.ColorScale.ScaleAlpha(a)
}

func (sp *Sprite) Render(suface *ebiten.Image) {
	suface.DrawImage(sp.img, sp.op)
}

func (sp *Sprite) Pos() (x, y int) {
	return int(sp.x), int(sp.y)
}

func (sp *Sprite) Rotate(rad float64) {
	sp.rot = rad
}
