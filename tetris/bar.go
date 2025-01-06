package tetris

import (
	"fmt"
	"strings"
)

// A bar represents a bar piece.
type bar struct {
	start      point // Bar leftmost block.
	color      string
	isVertical bool
}

// Creates a bar piece with a random colour, default starting position and direction.
func newBar() *bar {
	start := point{3, 0}
	color := newColor()

	return &bar{start, color, false}
}

// String implements fmt's Stringer interface.
// Return a string that if printed shows the piece on the board.
func (b bar) String() string {
	var sBuilder strings.Builder
	sBuilder.Grow(160) // First sBuilder size cappable of holding the entire string.

	sBuilder.WriteString(b.color) //Sets terminal ouput color.

	// Moves cursor to printing positon.
	cursorY := fmt.Sprintf(yTempl, initY+b.start.y*scaleY)
	cursorX := fmt.Sprintf(xTempl, initX+b.start.x*scaleX)
	sBuilder.WriteString(cursorY + cursorX)

	//Sets the bar with the given scale.
	for range scaleY {
		if b.isVertical {
			for range 4 {
				for range scaleX {
					sBuilder.WriteRune(block)
				}
				sBuilder.WriteString("\n\r" + cursorX) //Sets cursor to new line's first position.
			}
		} else {
			for range scaleX {
				for range 4 {
					sBuilder.WriteRune(block)
				}
			}
			sBuilder.WriteString("\n\r" + cursorX) //Sets cursor to new line's first position.
		}
	}

	return sBuilder.String()
}

// Clears the piece from the terminal.
func (b bar) clear() {
	var sBuilder strings.Builder
	sBuilder.Grow(64) // First sBuilder size cappable of holding the entire string.

	// Moves cursor to printing positon.
	cursorY := fmt.Sprintf(yTempl, initY+b.start.y*scaleY)
	cursorX := fmt.Sprintf(xTempl, initX+b.start.x*scaleX)
	sBuilder.WriteString(cursorY + cursorX)

	//Sets the bar with the given scale to whitespace.
	for range scaleY {
		if b.isVertical {
			for range 4 {
				for range scaleX {
					sBuilder.WriteRune(' ')
				}
				sBuilder.WriteString("\n\r" + cursorX) //Sets cursor to new line's first position.
			}
		} else {
			for range scaleX {
				for range 4 {
					sBuilder.WriteRune(' ')
				}
			}
			sBuilder.WriteString("\n\r" + cursorX) //Sets cursor to new line's first position.
		}
	}
	fmt.Println(sBuilder.String())
}

// Checks if any point in bor have collided with the given board.
func (b bar) hasCollided(board board) bool {
	p := b.start
	for range 4 {
		if board.hasCollided(p) {
			return true
		}
		if b.isVertical {
			p.move(south)
		} else {
			p.move(east)
		}
	}
	return false
}

// Moves a bar in the given direction, if bar collides with the board, nothing happens.
func (b *bar) move(board board, dir direction) {
	b.clear()

	tempBar := *b
	tempBar.start.move(dir)
	if !tempBar.hasCollided(board) {
		*b = tempBar
	}

	fmt.Println(b)
}
