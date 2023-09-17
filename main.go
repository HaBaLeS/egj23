package main

import (
	"fmt"
	"github.com/habales/egj23/games"
	"github.com/habales/egj23/gjfw"
	"github.com/habales/egj23/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
	_ "image/png"
	"log"
	"os"
)

var version string
var buildtime string
var deckBuild string
var gameMode bool
var currentGame games.JamGame
var gameList []games.JamGame

const UNDERLAY_BORDER = 5

type Game struct {
	bgm          *audio.Player
	selectionIdx int
	highlight    vector.Path
	underlay     *gjfw.Element
}

func (g *Game) Update() error {

	evl := util.GetInputEvents()
	//fmt.Printf("Got %d events: %v\n", len(evl), evl)

	if !gameMode && !g.bgm.IsPlaying() {
		g.bgm.SetVolume(gjfw.CFG.Volume)
		g.bgm.Seek(0)
		g.bgm.Play()
	}

	if evl[util.RIGHT] == 1 {
		g.selectionIdx++
		if g.selectionIdx >= len(gameList) {
			g.selectionIdx = len(gameList) - 1
		}
	}
	if evl[util.LEFT] == 1 {
		g.selectionIdx--
		if g.selectionIdx < 0 {
			g.selectionIdx = 0
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) || evl[util.BTN_B] == 1 {
		if !gameMode {
			os.Exit(1)
		} else {
			currentGame.Stop()
			currentGame = nil
			gameMode = false
		}
	}
	if !gameMode {
		if inpututil.IsKeyJustReleased(ebiten.KeySpace) || evl[util.BTN_A] == 1 {
			g.bgm.Pause()
			currentGame = gameList[g.selectionIdx]
			currentGame.Start()
			gameMode = true
		}

	} else {
		currentGame.Update(ebiten.ActualTPS(), evl)
	}
	return nil
}

func (g *Game) Draw(surface *ebiten.Image) {

	if !gameMode {
		g.underlay.Render(surface)
		for i, game := range gameList {
			if g.selectionIdx == i {
				g.underlay.X = float32(50 + i*250 - UNDERLAY_BORDER)
				g.underlay.Y = 50 - UNDERLAY_BORDER
			}
			game.Logo().SetPos(50+i*250, 50)
			game.Logo().Render(surface)

			w := game.Name().Bounds().Dx()
			d := (200 - w) / 2
			game.Name().SetPos(50+d+250*i, 275)
			game.Name().Render(surface)

		}
	} else {
		currentGame.Draw(surface)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gjfw.CFG.ScreenWidth, gjfw.CFG.ScreenHeight
}

func main() {

	gameList = games.Init()

	fmt.Printf("Version:%s -- BuildTime %s\n", version, buildtime)
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Private Easter Game Jam 2023")

	game := &Game{}
	game.bgm = util.LoadAudio("bgm/Otjanbird-Pt.-I.mp3")
	game.selectionIdx = 0
	game.underlay = &gjfw.Element{X: 0, Y: 0, H: 200 + 2*UNDERLAY_BORDER, W: 200 + 2*UNDERLAY_BORDER, Color: colornames.Whitesmoke}
	ebiten.SetScreenClearedEveryFrame(true)
	if deckBuild != "" {
		ebiten.SetFullscreen(true)
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
