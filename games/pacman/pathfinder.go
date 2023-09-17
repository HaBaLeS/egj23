package pacman

import (
	"github.com/habales/egj23/egmath"
	"reflect"
)

func (p *PacMan) Path(from, to egmath.Vec2) []*MazeField {
	//optimisation hint: do not reallocate the map app the time
	start := p.mazeMap[from]
	cameFrom := make(map[*MazeField]egmath.Vec2)
	frontier := make([]*MazeField, 0)
	reached := make(map[egmath.Vec2]struct{})
	frontier = append(frontier, start)

	cameFrom[p.mazeMap[from]] = from

	for len(frontier) > 0 {
		f := frontier[0]
		frontier = frontier[1:]
		if f.pos == to {
			break
		}
		for _, n := range f.Neighbours() {
			if _, exists := reached[n.pos]; !exists {
				frontier = append(frontier, n)
				reached[n.pos] = struct{}{}
				cameFrom[n] = f.pos
			}
		}
	}

	current := p.mazeMap[to]
	path := make([]*MazeField, 0)
	for current != start {

		cf := cameFrom[current]
		if cf.X == 0 && cf.Y == 0 { //this is a special case with the NULL vector, as there is always a wall at 0,0 we probably can do this hack
			break
		}
		path = append(path, current)
		current = p.mazeMap[cf]
	}
	reverse(path)
	return path
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
