package pacman

import (
	"github.com/habales/egj23/egmath"
	"github.com/habales/egj23/gjfw"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"image/color"
	"math"
)

const (
	STRAT_NONE = iota
	STRAT_CHASE
	STRAT_SCATTER
	STRAT_PANIC
)

const INKY_AHEAD_DIST = 4

var CLYDE_START_POS = egmath.Vec2{}

const (
	FACING_NORTH = iota
	FACING_SOUTH
	FACING_EAST
	FACING_WEST
)

const (
	ROT_NORTH = -0.5 * math.Pi
	ROT_EAST  = 0
	ROT_SOUTH = 0.5 * math.Pi
	ROT_WEST  = math.Pi
)

type MazeObject struct {
	//Player, Ghost, Collectible
	gridPos   egmath.Vec2
	screenPos egmath.Vec2f
	sprite    *gjfw.Sprite
	mainColor color.Color

	//cp physics
	playerBody *cp.Body
	targetBody *cp.Body

	strategy int
	facing   int
}

func NewMazeObject(sprite string, pos egmath.Vec2f) *MazeObject {
	obj := &MazeObject{}
	obj.sprite = gjfw.NewSprite(sprite)
	obj.screenPos = pos
	obj.sprite.SetSize(BRICK_SIZE, BRICK_SIZE)
	obj.facing = -1
	obj.strategy = STRAT_SCATTER
	return obj
}

func (mo *MazeObject) Update(world *PacMan) {
	//Teleport
	if mo.gridPos.X == 20 && mo.gridPos.Y == 13 {
		mo.gridPos.X = 1
		//calc screenPos
		//Set physics Target
	}
	//TODO do the same for the opposte direction
	//do not teleport on the teleporter teleport 1 field next to it
	//FIXME manually add the teleport to the pathfinder graph !!!

	//Calc gridPos from ScreenPos
	mo.gridPos.X = ((int(mo.screenPos.X) + BRICK_SIZE/2 - 300) / BRICK_SIZE) % mapWidth
	mo.gridPos.Y = ((int(mo.screenPos.Y) + BRICK_SIZE/2 - MAZE_PAD) / BRICK_SIZE) % mapHeight

	if mo.playerBody != nil {
		mo.screenPos.X = mo.playerBody.Position().X //FIXME move this into the maze Objects update function as soo as we know how to handle the player body and target
		mo.screenPos.Y = mo.playerBody.Position().Y //FIXME move this into the maze Objects
	}

	targetPos := world.pacman.gridPos
	var path []*MazeField
	if mo.strategy == STRAT_CHASE {
		if mo.mainColor == PINKY {
			switch world.pacman.facing {
			case FACING_NORTH:
				targetPos.Y -= INKY_AHEAD_DIST
			case FACING_SOUTH:
				targetPos.Y += INKY_AHEAD_DIST
			case FACING_EAST:
				targetPos.X += INKY_AHEAD_DIST
			case FACING_WEST:
				targetPos.X -= INKY_AHEAD_DIST
			}
			path = world.Path(mo.gridPos, targetPos)
		} else if mo.mainColor == CLYDE {
			path = world.Path(mo.gridPos, targetPos)
			if len(path) <= 8 {
				mo.strategy = STRAT_SCATTER
			}
		} else if mo.mainColor == BLINKY {
			path = world.Path(mo.gridPos, targetPos)
		} else if mo.mainColor == INKY {
			//Calc target pos by blaaaa
			path = world.Path(mo.gridPos, targetPos)
		}
	}

	if mo.strategy == STRAT_SCATTER {
		var startPos egmath.Vec2
		if mo.mainColor == CLYDE {
			startPos = egmath.Vec2{1, 24}
		}
		path = world.Path(mo.gridPos, startPos)
		if len(path) == 1 {
			mo.strategy = STRAT_CHASE
		}
	}

	if gjfw.CFG.Debug {
		for _, w := range path {
			if w != nil { //Can happen the path shows outside of the maze if pinky is trying to get ahead of pacman
				w.overlayColor = mo.mainColor
			}
		}
	}

	if len(path) > 0 {
		tx := 300 + path[0].pos.X*BRICK_SIZE
		ty := MAZE_PAD + path[0].pos.Y*BRICK_SIZE
		mo.targetBody.SetPosition(cp.Vector{float64(tx), float64(ty)})
	}

}

func (mo *MazeObject) Render(surface *ebiten.Image) {
	mo.sprite.SetPos(int(mo.screenPos.X), int(mo.screenPos.Y))
	switch mo.facing {
	case FACING_NORTH:
		mo.sprite.Rotate(ROT_NORTH)
	case FACING_SOUTH:
		mo.sprite.Rotate(ROT_SOUTH)
	case FACING_WEST:
		mo.sprite.Rotate(ROT_WEST)
	case FACING_EAST:
		mo.sprite.Rotate(ROT_EAST)
	}

	mo.sprite.Render(surface)
}
