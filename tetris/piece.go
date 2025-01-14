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
	forward = "\033[C"   // Moves the cursor forward once
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
	position      point // The position of the top-left corner of the piece on the board.
	boardPosition point
	color         string
	rotMatrix     [][][]bool // Matrix representing the possible states of the piece when rotating.
	rotationIdx   int        // Matrix id for the current rotation state.
}

// Generates and returns a random piece.
func newPiece() *piece {
	switch rand.IntN(7) {
	case 0:
		return newSquare()
	case 1:
		return newBar()
	case 2:
		return newT()
	case 3:
		return newS()
	case 4:
		return newL()
	case 5:
		return newJ()
	case 6:
		return newZ()
	}
	panic("Not a valid piece type")
}

// Returns a piece representing the bar piece with its default values.
func newBar() *piece {
	piece := &piece{}
	piece.position = point{11, 8}
	piece.boardPosition = point{3, -2}
	piece.color = newColor()
	piece.rotMatrix = [][][]bool{
		{
			{false, false, false, false},
			{false, false, false, false},
			{true, true, true, true},
			{false, false, false, false},
		},
		{
			{false, false, true, false},
			{false, false, true, false},
			{false, false, true, false},
			{false, false, true, false},
		},
	}

	return piece
}

// Returns a piece representing the J piece with its default values.
func newJ() *piece {
	piece := &piece{}
	piece.position = point{11, 9}
	piece.boardPosition = point{3, -1}
	piece.color = newColor()
	piece.rotMatrix = [][][]bool{
		{
			{false, false, false},
			{true, true, true},
			{false, false, true},
		},
		{
			{false, true, false},
			{false, true, false},
			{true, true, false},
		},
		{
			{true, false, false},
			{true, true, true},
			{false, false, false},
		},
		{
			{false, true, true},
			{false, true, false},
			{false, true, false},
		},
	}

	return piece
}

// Returns a piece representing the L piece with its default values.
func newL() *piece {
	piece := &piece{}
	piece.position = point{11, 9}
	piece.boardPosition = point{3, -1}
	piece.color = newColor()
	piece.rotMatrix = [][][]bool{
		{
			{false, false, false},
			{true, true, true},
			{true, false, false},
		},
		{
			{true, true, false},
			{false, true, false},
			{false, true, false},
		},
		{
			{false, false, true},
			{true, true, true},
			{false, false, false},
		},
		{
			{false, true, false},
			{false, true, false},
			{false, true, true},
		},
	}

	return piece
}

// Returns a piece representing the S piece with its default values.
func newS() *piece {
	piece := &piece{}
	piece.position = point{11, 9}
	piece.boardPosition = point{3, -1}
	piece.color = newColor()
	piece.rotMatrix = [][][]bool{
		{
			{false, false, false},
			{false, true, true},
			{true, true, false},
		},
		{
			{false, true, false},
			{false, true, true},
			{false, false, true},
		},
	}

	return piece
}

// Returns a piece representing the square piece with its default values.
func newSquare() *piece {
	piece := &piece{}
	piece.position = point{12, 10}
	piece.boardPosition = point{3, 0}
	piece.color = newColor()
	piece.rotMatrix = [][][]bool{
		{
			{true, true},
			{true, true},
		},
	}

	return piece
}

// Returns a piece representing the T piece with its default values.
func newT() *piece {
	piece := &piece{}
	piece.position = point{11, 9}
	piece.boardPosition = point{3, -1}
	piece.color = newColor()
	piece.rotMatrix = [][][]bool{
		{
			{false, false, false},
			{true, true, true},
			{false, true, false},
		},
		{
			{false, true, false},
			{true, true, false},
			{false, true, false},
		},
		{
			{false, true, false},
			{true, true, true},
			{false, false, false},
		},
		{
			{false, true, false},
			{false, true, true},
			{false, true, false},
		},
	}

	return piece
}

// Returns a piece representing the Z piece with its default values.
func newZ() *piece {
	piece := &piece{}
	piece.position = point{11, 9}
	piece.boardPosition = point{3, -1}
	piece.color = newColor()
	piece.rotMatrix = [][][]bool{
		{
			{false, false, false},
			{true, true, false},
			{false, true, true},
		},
		{
			{false, false, true},
			{false, true, true},
			{false, true, false},
		},
	}

	return piece
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
	for y := range p.rotMatrix[p.rotationIdx] {
		for range scaleY {
			for x := range p.rotMatrix[p.rotationIdx][y] {
				for range scaleX {
					isVisible := p.position.y+y >= 0
					if p.rotMatrix[p.rotationIdx][y][x] && isVisible {
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
	for y := range p.rotMatrix[p.rotationIdx] {
		for range scaleY {
			for x := range p.rotMatrix[p.rotationIdx][y] {
				for range scaleX {
					isVisible := p.position.y+y >= 0
					if p.rotMatrix[p.rotationIdx][y][x] && isVisible {
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
	for y := range p.rotMatrix[p.rotationIdx] {
		for x := range p.rotMatrix[p.rotationIdx][y] {
			if p.rotMatrix[p.rotationIdx][y][x] {
				if board.hasCollided(point{pos.x + x, pos.y + y}) {
					return true
				}
			}
		}
	}
	return false
}

// Moves a piece in the given direction, if piece collides with the board.
// Returns true if piece has collided ans can't be moved.
func (p *piece) move(board board, dir direction) bool {
	tempPiece := *p
	tempPiece.position.move(dir)
	if !tempPiece.hasCollided(board) {
		p.clear()
		*p = tempPiece
		fmt.Println(p)
		return false
	}

	return true
}

func (p *piece) moveToBoard() {
	p.position = p.boardPosition
}

// Apply moves to the piece based on the keys pressed during the frame.
// Returns true if a piece collides when moving down.
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

// Locks piece in the board, sets position in board to true representing a
// new colision point.
func (p piece) lock(b *board) {
	for y := range p.rotMatrix[p.rotationIdx] {
		for x := range p.rotMatrix[p.rotationIdx][y] {
			if p.rotMatrix[p.rotationIdx][y][x] {
				b[p.position.y+y][p.position.x+x] = true
			}
		}
	}
}

// Rotates piece clockwise, if possible.
func (p *piece) rotate(board board) {
	tempPiece := *p
	tempPiece.rotationIdx++
	tempPiece.rotationIdx %= len(tempPiece.rotMatrix)
	if !tempPiece.hasCollided(board) {
		p.clear()
		*p = tempPiece
		fmt.Println(p)
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
