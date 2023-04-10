package games

import (
	"github.com/habales/egj23/games/pacman"
	"github.com/habales/egj23/games/pong"
	"github.com/habales/egj23/gjfw"
	"github.com/hajimehoshi/ebiten/v2"
)

var games []JamGame

type JamGame interface {
	Name() *gjfw.TextSprite
	Logo() *gjfw.Sprite
	Start()
	Stop()
	Update(d float64, ev map[int]float64)
	Draw(screen *ebiten.Image)
}

func Init() []JamGame {
	games = make([]JamGame, 0)

	games = append(games, pong.New())
	games = append(games, pacman.New())
	return games
}
