package pong

import (
	"fmt"
	"github.com/habales/egj23/gjfw"
	"github.com/habales/egj23/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

const (
	MAX_SPEED           = 8.0
	BALL_BASE_SPEED     = 5.0
	BALL_SPEEDUP_FACTOR = 1.1
	BALL_MAX_SPEED      = 10
	PADDLE_HEIGHT       = float32(100.0)
	BALL_HEIGHT         = 10
)

type Pong struct {
	upperBorder *gjfw.Element
	lowerBorder *gjfw.Element
	paddleA     *gjfw.Element
	paddleB     *gjfw.Element
	ball        *gjfw.Element
	scoreA      int
	scoreB      int
	background  *gjfw.Sprite
	ballVecX    float32
	ballVecY    float32
	bgm         *audio.Player
	sfxA        *audio.Player
	sfxB        *audio.Player
	sfxWin      *audio.Player
	sfxLoose    *audio.Player
	sfxBounce   *audio.Player
	font        font.Face

	//Stuff for overview for
	announceFont font.Face
	nameSprite   *gjfw.TextSprite
	logo         *gjfw.Sprite
}

func (g *Pong) Font() font.Face {
	return g.announceFont
}

func New() *Pong {
	p := &Pong{}
	paddleWidth := float32(15.0)

	lineHeight := float32(20.0)
	margin := float32(30.0)
	w := float32(gjfw.CFG.ScreenWidth) - 2*margin
	p.upperBorder = &gjfw.Element{X: margin, Y: margin, W: w, H: lineHeight, Color: colornames.Limegreen}

	lx := float32(gjfw.CFG.ScreenHeight) - (margin + lineHeight)
	p.lowerBorder = &gjfw.Element{X: margin, Y: lx, W: w, H: lineHeight, Color: colornames.Limegreen}

	p.ball = gjfw.NewElementRectC(300, float32(gjfw.CFG.ScreenHeight)/2, BALL_HEIGHT, BALL_HEIGHT, colornames.Darkred)

	p.paddleA = gjfw.NewElementRectC(margin, 720/2-0.5*PADDLE_HEIGHT, paddleWidth, PADDLE_HEIGHT, colornames.Greenyellow)
	p.paddleB = gjfw.NewElementRectC(1280-margin-paddleWidth, 720/2-0.5*PADDLE_HEIGHT, paddleWidth, PADDLE_HEIGHT, colornames.Greenyellow)

	p.background = gjfw.NewSprite("pong/01522-3155113783-spiral flash_ dream.png")
	p.background.SetSize(float64(gjfw.CFG.ScreenWidth), float64(gjfw.CFG.ScreenHeight))
	p.background.SetAlpha(0.3)

	p.ballVecY = 0.0
	p.ballVecX = 1.0

	p.bgm = util.LoadAudio("pong/Samurai-Sake-Showdown.mp3")
	p.sfxA = util.LoadAudio("pong/boom-geomorphism-cinematic-trailer-sound-effects-123876.mp3")
	p.sfxB = util.LoadAudio("pong/hit-brutal-puncher-cinematic-trailer-sound-effects-124760.mp3")

	p.sfxLoose = util.LoadAudio("pong/war-horn-blast-14760.mp3")
	p.sfxWin = util.LoadAudio("pong/ding-36029.mp3")
	p.sfxBounce = util.LoadAudio("pong/metal-hit-cartoon-7118.mp3")

	p.font = util.LoadFont("fonts/ka1.ttf", 4*24)
	p.announceFont = util.LoadFont("fonts/ka1.ttf", 24)
	p.nameSprite = gjfw.NewTextSprite(p.announceFont, "p0ng", colornames.Limegreen)

	p.logo = gjfw.NewSprite("pong/logo2.png")
	p.logo.SetSize(200, 200)

	return p

}

func (g *Pong) Start() {
	g.bgm.SetVolume(gjfw.CFG.Volume)
	g.bgm.SetVolume(0.3)
	g.bgm.Play()
}

func (g *Pong) Stop() {
	g.bgm.Pause()
	g.bgm.Seek(0)
}

func (g *Pong) Update(d float64, ev map[int]float64) {

	//Player Move
	moveFactor := 0.0

	if ev[util.UP] > 0.5 {
		moveFactor = -1
	}
	if ev[util.DOWN] > 0.5 {
		moveFactor = 1
	}

	speed := MAX_SPEED * moveFactor
	g.paddleA.Y += float32(speed)
	if g.paddleA.Intersect(g.upperBorder) || g.paddleA.Intersect(g.lowerBorder) {
		g.paddleA.Y -= float32(speed)
	}

	//AI Move
	//If ball is above us we move up else we move down
	paddleBMiddle := g.paddleB.Y + PADDLE_HEIGHT/2
	diff := paddleBMiddle - g.ball.Y
	if diff < 0 {
		if -1*diff > MAX_SPEED {
			g.paddleB.Y -= diff
		} else {
			g.paddleB.Y += MAX_SPEED
		}

	} else if diff > 0 {
		if diff > MAX_SPEED {
			g.paddleB.Y -= MAX_SPEED
		} else {
			g.paddleB.Y -= diff
		}

	}

	//Calc ball stuff

	//move Ball

	//FIXME !!! TUNNEL EFFECT!! --- iterate over speed in small chuncs!!
	if g.ballVecX > BALL_MAX_SPEED {
		g.ballVecX = BALL_MAX_SPEED
	}
	if g.ballVecY > BALL_MAX_SPEED {
		g.ballVecY = BALL_MAX_SPEED
	}

	g.ball.X += g.ballVecX * BALL_BASE_SPEED
	g.ball.Y += g.ballVecY * BALL_BASE_SPEED

	if g.ball.Intersect(g.paddleA) {
		g.ballVecX = -g.ballVecX
		g.ballVecX = BALL_SPEEDUP_FACTOR * g.ballVecX
		paddleMiddle := g.paddleA.Y + PADDLE_HEIGHT/2
		dy := g.ball.Y - paddleMiddle
		g.ballVecY = (100 / PADDLE_HEIGHT * dy) / 100 * BALL_BASE_SPEED
		g.sfxA.Seek(0)
		g.sfxA.Play()
	} else if g.ball.Intersect(g.paddleB) {
		g.ballVecX = -g.ballVecX
		g.ballVecX = BALL_SPEEDUP_FACTOR * g.ballVecX
		paddleMiddle := g.paddleB.Y + PADDLE_HEIGHT/2
		dy := g.ball.Y - paddleMiddle
		g.ballVecY = (100 / PADDLE_HEIGHT * dy) / 100 * BALL_BASE_SPEED
		g.sfxB.Seek(0)
		g.sfxB.Play()
	} else if g.ball.Intersect(g.lowerBorder) {
		g.ballVecY = -g.ballVecY
		g.sfxBounce.Seek(0)
		g.sfxBounce.Play()
	} else if g.ball.Intersect(g.upperBorder) {
		g.ballVecY = -g.ballVecY
		g.sfxBounce.Seek(0)
		g.sfxBounce.Play()
	} else if g.ball.X < 0 {
		g.scoreB += 1
		g.resetBall()
		g.sfxLoose.Seek(0)
		g.sfxLoose.Play()
	} else if g.ball.X > float32(gjfw.CFG.ScreenWidth) {
		g.scoreA += 1
		g.resetBall()
		g.sfxWin.Seek(0)
		g.sfxWin.Play()
	}

}

func (g *Pong) resetBall() {
	g.ball.X = 300
	g.ball.Y = float32(gjfw.CFG.ScreenHeight) / 2
	g.ballVecY = 0.0
	g.ballVecX = 1.0
}

func (g *Pong) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %d FPS: %f", ebiten.TPS(), ebiten.ActualFPS()))
	g.background.Render(screen)
	g.lowerBorder.Render(screen)
	g.upperBorder.Render(screen)
	text.Draw(screen, fmt.Sprintf("%d", g.scoreA), g.font, 400, 150, colornames.Limegreen)
	text.Draw(screen, fmt.Sprintf("%d", g.scoreB), g.font, 780, 150, colornames.Limegreen)

	g.paddleA.Render(screen)
	g.paddleB.Render(screen)

	g.ball.Render(screen)

}

func (g *Pong) Name() *gjfw.TextSprite {
	return g.nameSprite
}

func (g *Pong) Logo() *gjfw.Sprite {
	return g.logo
}
