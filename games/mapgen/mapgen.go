package mapgen

import (
	"fmt"
	"github.com/habales/egj23/gjfw"
	"github.com/habales/egj23/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

var SW = float32(1280.0)
var SH = float32(800.0)

type MapGen struct {
	logo       *gjfw.Sprite
	name       *gjfw.TextSprite
	spriteSize float32
}

func (m *MapGen) Name() *gjfw.TextSprite {
	return m.name
}

func (m *MapGen) Logo() *gjfw.Sprite {
	return m.logo
}

func (m *MapGen) Start() {
}

func (m *MapGen) Stop() {
}

func (m *MapGen) Update(d float64, ev map[int]float64) {
}

func (m *MapGen) Draw(surface *ebiten.Image) {
	rows := float32(int(SH / m.spriteSize))
	cols := float32(int(SW / m.spriteSize))

	ebitenutil.DebugPrint(surface, fmt.Sprintf("R: %f, C: %f", rows, cols))

	for r := float32(0.0); r < rows; r++ {
		vector.StrokeLine(surface, 0, r*m.spriteSize, SW, r*m.spriteSize, 1, colornames.Lightslategrey, false)
	}
	for c := float32(0.0); c < cols; c++ {
		vector.StrokeLine(surface, c*m.spriteSize, 0, c*m.spriteSize, SH, 1, colornames.Dimgrey, false)
	}

}

func New(spriteSize float32) *MapGen {
	fnt := util.LoadFont("mapgen/MorrisRomanBlack.ttf", 72)
	mg := &MapGen{
		logo:       gjfw.NewSprite("mapgen/logo.png"),
		name:       gjfw.NewTextSprite(fnt, "Map Gen", colornames.Darkgoldenrod),
		spriteSize: spriteSize,
	}
	mg.logo.SetSize(200, 200) //fixme start with a default size!!
	return mg
}
