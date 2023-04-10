package pacman

import (
	"github.com/habales/egj23/gjfw"
	"github.com/habales/egj23/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

const (
	BG_TIME  = 30000.0
	MAZE_PAD = 10
)

type PacMan struct {
	logo         *gjfw.Sprite
	announceFont font.Face
	nameSprite   *gjfw.TextSprite

	bgList     []*gjfw.Sprite
	currentIdx int
	nextBg     float64
	testimgage *gjfw.Sprite

	bgm *audio.Player
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

func (p *PacMan) Draw(screen *ebiten.Image) {
	p.bgList[p.currentIdx].Render(screen)

	p.testimgage.Render(screen)
}
