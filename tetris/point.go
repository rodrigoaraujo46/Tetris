package tetris

// Represents a block's position.
type point struct {
	x int
	y int
}

// Stores the way to move a point based on a fiven direction.
var directionDelta = [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

// Moves a point to a new position based on a direction
func (p *point) move(dir direction) {
	p.x += directionDelta[dir][0]
	p.y += directionDelta[dir][1]
}
