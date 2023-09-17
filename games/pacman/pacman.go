package pacman

import (
	"fmt"
	"github.com/habales/egj23/egmath"
	"github.com/habales/egj23/gjfw"
	"github.com/habales/egj23/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"image/color"
	"math"
)

const (
	BG_TIME    = 30000.0
	MAZE_PAD   = 5
	BRICK_SIZE = 26
)

var C_SOLID = color.RGBA{118, 66, 138, 255}
var C_OPEN = color.RGBA{0, 0, 0, 255}
var C_TELEPORT = color.RGBA{255, 0, 0, 255}
var C_POWERUP = color.RGBA{255, 242, 0, 255}
var C_EMPTY = color.RGBA{0, 0, 0, 0}

var PINKY = color.RGBA{255, 184, 255, 50}
var BLINKY = color.RGBA{255, 0, 0, 50}
var INKY = color.RGBA{0, 255, 255, 50}
var CLYDE = color.RGBA{255, 184, 82, 50}

var wallSprite *gjfw.Sprite
var floorSprite *gjfw.Sprite
var mapWidth int
var mapHeight int

type PacMan struct {
	logo         *gjfw.Sprite
	announceFont font.Face
	nameSprite   *gjfw.TextSprite

	bgList     []*gjfw.Sprite
	currentIdx int
	nextBg     float64

	bgm *audio.Player

	mazeMap map[egmath.Vec2]*MazeField

	pacman *MazeObject
	inky   *MazeObject
	blinky *MazeObject
	pinky  *MazeObject
	clyde  *MazeObject

	//physics
	space *cp.Space
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

	p.bgm = util.LoadAudio("pacman/pratzapp-sakura-hz-thar-3.mp3")

	wallSprite = gjfw.NewSprite("pacman/wall_01.png")
	wallSprite.SetSize(BRICK_SIZE, BRICK_SIZE)

	floorSprite = gjfw.NewSprite("pacman/floor_01.png")
	floorSprite.SetSize(BRICK_SIZE, BRICK_SIZE)

	return p
}

func (p *PacMan) Name() *gjfw.TextSprite {
	return p.nameSprite
}

func (p *PacMan) Logo() *gjfw.Sprite {
	return p.logo
}

func (p *PacMan) Start() {
	p.bgm.SetVolume(gjfw.CFG.Volume)
	p.bgm.Seek(0)
	p.bgm.Play()

	p.space = cp.NewSpace()
	p.space.Iterations = 1

	mapImg := util.LoadImage("pacman/play_field.png")
	mapWidth = mapImg.Bounds().Dx()
	mapHeight = mapImg.Bounds().Dy()
	p.mazeMap = make(map[egmath.Vec2]*MazeField, mapWidth*mapWidth)
	for i := 0; i < mapHeight; i++ {
		for j := 0; j < mapWidth; j++ {
			pos := egmath.Vec2{j, i}
			mf := NewMazeField(pos)
			p.mazeMap[pos] = mf

			c := mapImg.At(pos.X, pos.Y)
			if c == C_SOLID {
				mf.mainSprite = wallSprite
				wall := cp.NewStaticBody()
				wall.SetPosition(cp.Vector{float64(float32(300 + j*BRICK_SIZE)), float64(float32(MAZE_PAD + i*BRICK_SIZE))})
				wallShape := cp.NewBox(wall, BRICK_SIZE, BRICK_SIZE, 0)
				p.space.AddBody(wall)
				p.space.AddShape(wallShape)
				mf.moveCost = math.MaxInt
			} else if c == C_OPEN {
				mf.mainSprite = floorSprite
				//TODO add to some overlay a collectable sprite
			} else if c == C_TELEPORT {
				//we need to tell this  element it's special, if i collide with it, i trigger a teleport (for the thing that runs into it)
			} else if c == C_POWERUP {
				//TODO add to some overlay a powerup sprite
			} else if c == C_EMPTY {
				mf.moveCost = math.MaxInt
				//Noop we dont to things here
			} else {
				panic(fmt.Errorf("unhandled error %v", c))
			}
		}
	}

	//build the neighbors!
	for _, mf := range p.mazeMap {
		mf.north = p.mazeMap[egmath.Vec2{mf.pos.X, mf.pos.Y - 1}]
		mf.south = p.mazeMap[egmath.Vec2{mf.pos.X, mf.pos.Y + 1}]
		mf.east = p.mazeMap[egmath.Vec2{mf.pos.X + 1, mf.pos.Y}]
		mf.west = p.mazeMap[egmath.Vec2{mf.pos.X - 1, mf.pos.Y}]
	}

	p.pacman = NewMazeObject("pacman/pm_01.png", egmath.Vec2f{375, 342})
	p.SetUpPhysic(p.pacman)

	p.pinky = NewMazeObject("pacman/pinky_01.png", egmath.Vec2f{741, 135})
	p.pinky.mainColor = PINKY
	p.SetUpPhysic(p.pinky)

	p.inky = NewMazeObject("pacman/inky_01.png", egmath.Vec2f{326, 630})
	p.inky.mainColor = INKY
	p.SetUpPhysic(p.inky)

	p.blinky = NewMazeObject("pacman/blinky_01.png", egmath.Vec2f{585, 343})
	p.blinky.mainColor = BLINKY
	p.SetUpPhysic(p.blinky)

	p.clyde = NewMazeObject("pacman/clyde_01.png", egmath.Vec2f{749, 653})
	p.clyde.mainColor = CLYDE
	p.SetUpPhysic(p.clyde)

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

	cur := p.pacman.playerBody.Position()
	if ev[util.RIGHT] != 0 { //FIXM RL MESSED up!?!?!?
		cur.X += 50
		p.pacman.facing = FACING_EAST //FIXME i might be going east west .. where am i facing?
	}
	if ev[util.LEFT] != 0 {
		cur.X -= 50
		p.pacman.facing = FACING_WEST
	}
	if ev[util.UP] != 0 {
		cur.Y -= 50
		p.pacman.facing = FACING_NORTH
	}
	if ev[util.DOWN] != 0 {
		cur.Y += 50
		p.pacman.facing = FACING_SOUTH
	}
	if ev[util.AXIS_LEFT_HZN] != 0 {
		val := ev[util.AXIS_LEFT_HZN]
		if val > 0 {
			p.pacman.facing = FACING_EAST
			cur.X += 50 * math.Pow(math.Abs(val), 2.7)
		} else {
			p.pacman.facing = FACING_WEST
			cur.X -= 50 * math.Pow(math.Abs(val), 2.7)
		}

	}
	if ev[util.AXIS_LEFT_VERT] != 0 {
		val := ev[util.AXIS_LEFT_VERT]
		if val > 0 {
			p.pacman.facing = FACING_SOUTH
			cur.Y += 50 * math.Pow(math.Abs(val), 2.7)
		} else {
			p.pacman.facing = FACING_NORTH
			cur.Y -= 50 * math.Pow(math.Abs(val), 2.7)
		}
	}

	p.pacman.targetBody.SetPosition(cur)
	p.space.Step(1.0 / float64(ebiten.TPS()))

	p.pacman.Update(p)
	p.pinky.Update(p)
	p.inky.Update(p)
	p.clyde.Update(p)
	p.blinky.Update(p)

}

func (p *PacMan) Draw(surface *ebiten.Image) {
	p.bgList[p.currentIdx].Render(surface)

	for _, me := range p.mazeMap {
		me.Render(surface)
	}

	p.pinky.Render(surface)
	p.inky.Render(surface)
	p.clyde.Render(surface)
	p.blinky.Render(surface)
	p.pacman.Render(surface)

	ebitenutil.DebugPrint(surface, fmt.Sprintf("P:(%d:%d) -> (%.2f:%2.f)", p.pacman.gridPos.X, p.pacman.gridPos.Y, p.pacman.screenPos.X, p.pacman.screenPos.Y))
}

func (p *PacMan) SetUpPhysic(obj *MazeObject) {
	obj.playerBody = cp.NewBody(1.0, cp.INFINITY)
	obj.playerBody.SetPosition(cp.Vector(obj.screenPos))

	playsShape := cp.NewCircle(obj.playerBody, BRICK_SIZE/2, cp.Vector{0, 0})
	playsShape.SetFriction(0.1)
	playsShape.Filter = cp.NewShapeFilter(1, 1, 1)
	obj.targetBody = cp.NewBody(cp.INFINITY, cp.INFINITY)
	obj.targetBody.SetPosition(cp.Vector(obj.screenPos))

	joint := cp.NewPivotJoint(obj.targetBody, obj.playerBody, cp.Vector{0, 0})
	joint.SetMaxBias(200.0)
	joint.SetMaxForce(3000.0)

	p.space.AddBody(obj.playerBody)
	p.space.AddShape(playsShape)
	p.space.AddConstraint(joint)
}

type MazeField struct {
	pos            egmath.Vec2
	teleportTarget *MazeField

	north      *MazeField
	east       *MazeField
	south      *MazeField
	west       *MazeField
	neighbours []*MazeField

	collectable *Collectable
	mainSprite  *gjfw.Sprite
	//todo add List of some sprites like blood and etc...

	moveCost     int
	overlayColor color.Color
}

func NewMazeField(pos egmath.Vec2) *MazeField {
	return &MazeField{pos: pos}
}

func (n *MazeField) Neighbours() []*MazeField {
	//FIXME add caching
	retList := make([]*MazeField, 0)
	if n.north != nil && n.moveCost != math.MaxInt {
		retList = append(retList, n.north)
	}
	if n.south != nil && n.moveCost != math.MaxInt {
		retList = append(retList, n.south)
	}
	if n.west != nil && n.moveCost != math.MaxInt {
		retList = append(retList, n.west)
	}
	if n.east != nil && n.moveCost != math.MaxInt {
		retList = append(retList, n.east)
	}
	return retList
}

func (n *MazeField) Render(surface *ebiten.Image) {
	if n.mainSprite != nil {
		n.mainSprite.SetPos(300+n.pos.X*BRICK_SIZE, MAZE_PAD+n.pos.Y*BRICK_SIZE)
		n.mainSprite.Render(surface)
	}
	if n.overlayColor != nil {
		if n.mainSprite != nil {
			x, y := n.mainSprite.Pos()
			vector.DrawFilledRect(surface, float32(x), float32(y), BRICK_SIZE, BRICK_SIZE, n.overlayColor, true)
		}
		n.overlayColor = nil
	}
}

type Collectable struct {
	image    *gjfw.Sprite
	itemType int
}
