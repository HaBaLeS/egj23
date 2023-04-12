package pacman

import (
	"github.com/habales/egj23/gjfw"
	"github.com/habales/egj23/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"image/color"
)

const (
	BG_TIME    = 30000.0
	MAZE_PAD   = 5
	BRICK_SIZE = 26
)

var C_SOLID = color.RGBA{118, 66, 138, 255}
var C_OPEN = color.RGBA{0, 0, 0, 255}

type PacMan struct {
	logo         *gjfw.Sprite
	announceFont font.Face
	nameSprite   *gjfw.TextSprite

	bgList     []*gjfw.Sprite
	currentIdx int
	nextBg     float64
	testimgage *gjfw.Sprite

	bgm *audio.Player

	mazeData []*gjfw.Element
}

func New() *PacMan {
	p := &PacMan{}

	p.announceFont = util.LoadFont("pacman/MB-ThinkTwice_Regular.ttf", 72)

	p.logo = gjfw.NewSprite("pacman/img_4.png")
	p.logo.SetSize(200, 200)

	p.nameSprite = gjfw.NewTextSprite(p.announceFont, "pacman", colornames.Orangered)

	p.bgList = append(p.bgList, gjfw.NewSprite("pacman/yellow.png"))
	p.bgList = append(p.bgList, gjfw.NewSprite("pacman/pink.png"))
	p.bgList = append(p.bgList, gjfw.NewSprite("pacman/red.png"))
	p.bgList = append(p.bgList, gjfw.NewSprite("pacman/orange.png"))
	p.bgList = append(p.bgList, gjfw.NewSprite("pacman/cyan.png"))
	p.nextBg = BG_TIME

	p.testimgage = gjfw.NewSprite("pacman/test.png")
	mazeSize := float64(gjfw.CFG.ScreenHeight - 2*MAZE_PAD)
	p.testimgage.SetSize(mazeSize, mazeSize)
	p.testimgage.SetPos(int(gjfw.CFG.ScreenWidth/2-int(mazeSize)/2), MAZE_PAD)

	p.bgm = util.LoadAudio("pacman/pratzapp-sakura-hz-thar-3.mp3")

	return p
}

func (p *PacMan) Name() *gjfw.TextSprite {
	return p.nameSprite
}

func (p *PacMan) Logo() *gjfw.Sprite {
	return p.logo
}

func (p *PacMan) Start() {
	p.bgm.Seek(0)
	p.bgm.Play()

	mapImg := util.LoadImage("pacman/play_field.png")
	for i := 0; i < mapImg.Bounds().Dy(); i++ {
		for j := 0; j < mapImg.Bounds().Dx(); j++ {
			c := mapImg.At(j, i)
			if c == C_SOLID {
				p.mazeData = append(p.mazeData, gjfw.NewSquareC(float32(300+j*BRICK_SIZE), float32(MAZE_PAD+i*BRICK_SIZE), BRICK_SIZE, colornames.Cornflowerblue))
			} else if c == C_OPEN {
				p.mazeData = append(p.mazeData, gjfw.NewSquareC(float32(300+j*BRICK_SIZE), float32(MAZE_PAD+i*BRICK_SIZE), BRICK_SIZE, color.Black))
			}
		}
	}
}

func (p *PacMan) Stop() {
	p.bgm.Pause()
}

func (p *PacMan) Update(d float64, ev map[int]float64) {
	p.nextBg = p.nextBg - d
	if p.nextBg < 0 {
		p.nextBg = BG_TIME
		p.currentIdx++
		if p.currentIdx >= len(p.bgList) {
			p.currentIdx = 0
		}
	}
}

func (p *PacMan) Draw(surface *ebiten.Image) {
	p.bgList[p.currentIdx].Render(surface)

	for _, b := range p.mazeData {
		b.Render(surface)
	}

}
