package tetris

import (
	"math/rand/v2"
)

const (
	block  = '█'
	yTempl = "\033[%dH" //Preped string to move cursor in y axis.
	xTempl = "\033[%dC" //Preped string to move cursor in x axis.
)

// Direction is an int type used with ther directions iota.
type direction int

// Cardinal directions enum.
const (
	north direction = iota
	east
	south
	west
)

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

// Interface that enables functions to work with all different kinds of pieces.
type pieceManager interface {
	String() string
	clear()
	move(board, direction)
	hasCollided(board) bool
}

// Randomly selects a color from the following:
//
//	Yellow, red, blue, green, purple, orange, pink.
//
// Returns a string that can set the terminal ouput to the given color.
func newColor() string {

	// Strings that set terminal ouput to the given color.
	const (
		yellow = "\033[38;2;255;255;0m"
		red    = "\033[38;2;255;0;0m"
		blue   = "\033[38;2;0;0;255m"
		green  = "\033[38;2;0;255;255m"
		purple = "\033[38;2;128;0;128m"
		orange = "\033[38;2;255;165;0m"
		pink   = "\033[38;2;255;192;203m"
	)

	switch rand.IntN(7) {
	case 0:
		return yellow
	case 1:
		return red
	case 2:
		return blue
	case 3:
		return green
	case 4:
		return purple
	case 5:
		return orange
	default:
		return pink
	}
}

// Generates a random piece.
// Returns pieceManager interface
func newPiece() pieceManager {
	switch rand.IntN(7) {
	default:
		return newBar()
	}
}
