package tetris

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const (
	block   = '█'
	yTempl  = "\033[%dH" // Preped string to move cursor in y axis.
	xTempl  = "\033[%dC" // Preped string to move cursor in x axis.
	forward = "\033[C" // Moves the cursor forward once
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

// Piece represents a tetris piece.
type piece struct {
	position    point
	color       string
	matrix      [][]bool
	rotationIdx int
}

// Creates a new piece and set its color.
func newPiece() *piece {
	color := newColor()

	return &piece{color: color}
}

// String implements fmt's Stringer interface.
// Return a string that if printed shows the piece on the board.
func (p piece) String() string {
	var sBuilder strings.Builder

	sBuilder.WriteString(p.color) //Sets terminal ouput color.

	// Moves cursor to printing positon.
	cursorY := fmt.Sprintf(yTempl, initY+p.position.y*scaleY)
	cursorX := fmt.Sprintf(xTempl, initX+p.position.x*scaleX)
	sBuilder.WriteString(cursorY + cursorX)

	//Sets the bar with the given scale.
	for y := range p.matrix {
		for range scaleY {
			for x := range p.matrix[y] {
				for range scaleX {
					if p.matrix[y][x] {
						sBuilder.WriteRune(block)
					} else {
						sBuilder.WriteString(forward)
					}
				}
			}
			sBuilder.WriteString("\n\r" + cursorX) //Sets cursor to new line's first position.
		}
	}

	return sBuilder.String()
}

// Clears the piece from the terminal.
func (p piece) clear() {
	var sBuilder strings.Builder

	// Moves cursor to printing positon.
	cursorY := fmt.Sprintf(yTempl, initY+p.position.y*scaleY)
	cursorX := fmt.Sprintf(xTempl, initX+p.position.x*scaleX)
	sBuilder.WriteString(cursorY + cursorX)

	//Sets the bar with the given scale to whitespace.
	for y := range p.matrix {
		for range scaleY {
			for x := range p.matrix[y] {
				for range scaleX {
					if p.matrix[y][x] {
						sBuilder.WriteRune(' ')
					} else {
						sBuilder.WriteString(forward)
					}
				}
			}
			sBuilder.WriteString("\n\r" + cursorX) //Sets cursor to new line's first position.
		}
	}
	fmt.Println(sBuilder.String())
}

// Checks if any point in a piece has collided with the given board.
func (p piece) hasCollided(board board) bool {
	pos := p.position
	for y := range p.matrix {
		for x := range p.matrix[y] {
			if p.matrix[y][x] {
				if board.hasCollided(point{pos.x + x, pos.y + y}) {
					return true
				}
			}
		}
	}
	return false
}

// Moves a piece in the given direction, if piece collides with the board, nothing happens.
func (p *piece) move(board board, dir direction) bool {
	tempBar := *p
	tempBar.position.move(dir)
	if !tempBar.hasCollided(board) {
		p.clear()
		*p = tempBar
		fmt.Println(p)
		return false
	}

	return true
}

func (p *piece) applyMoves(keysP keysPressed, b board) bool {
	if keysP.right {
		p.move(b, east)
	}
	if keysP.left {
		p.move(b, west)
	}
	if keysP.down {
		if p.move(b, south) {
			return true
		}
	}

	return false
}

func (p piece) lock(b *board) {
	for y := range p.matrix {
		for x := range p.matrix[y] {
			if p.matrix[y][x] {
				b[p.position.y+y][p.position.x+x] = true
			}
		}
	}
}

// Randomly selects a color from the following:
//
//	Yellow, red, blue, green, purple, orange, pink.
//
// Returns a string that can set the terminal ouput to the given color.
func newColor() string {

	// Strings that set terminal ouput to the given color.
	colors := [7]string{
		"\033[38;2;255;255;0m",
		"\033[38;2;255;0;0m",
		"\033[38;2;0;0;255m",
		"\033[38;2;0;255;255m",
		"\033[38;2;128;0;128m",
		"\033[38;2;255;165;0m",
		"\033[38;2;255;192;203m",
	}

	return colors[rand.IntN(7)]
}
