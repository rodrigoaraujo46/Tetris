package tetris

import "strings"

const (
	rows   = 20
	cols   = 10
	scaleX = 5 // Scale at which the board is printed in the X axis.
	scaleY = 2 // Scale at which the board is printed in the Y axis.
	initX  = 1 // X positon for top left corner of the board.
	initY  = 8 // Y positon for top left corner of the board.
)

// Represents the board, true represents a colision point.
type board [rows][cols]bool

// Creates a new board and returns a pointer to it.
func makeBoard() *board {
	return new(board)
}

// String implements fmt's Stringer interface.
// Returns a string that if printed displayed the board.
func (b board) String() string {
	const (
		boardStrSize   = 2304 // Byte size of board string builder.
		horizontalWall = '═'
		verticalWall   = '║'
		topLeft        = '╔'
		topRight       = '╗'
		bottomLeft     = '╚'
		bottomRight    = '╝'
	)

	var sBuilder strings.Builder
	sBuilder.Grow(boardStrSize)

	//Top border
	sBuilder.WriteRune(topLeft)
	for range cols * scaleX {
		sBuilder.WriteRune(horizontalWall)
	}
	sBuilder.WriteRune(topRight)
	sBuilder.WriteString("\n\r")

	//Left and right border
	for range rows {
		for range scaleY {
			sBuilder.WriteRune(verticalWall)
			for range cols {
				for range scaleX {
					sBuilder.WriteRune(' ')
				}
			}
			sBuilder.WriteRune(verticalWall)
			sBuilder.WriteString("\n\r")
		}
	}

	//Bottom border
	sBuilder.WriteRune(bottomLeft)
	for range cols * scaleX {
		sBuilder.WriteRune(horizontalWall)
	}
	sBuilder.WriteRune(bottomRight)
	sBuilder.WriteString("\n\n\r")
	return sBuilder.String()
}

// Checks if point is in bounds and if it doesn't collide with other blocks.
func (b board) hasCollided(p point) bool {
	if p.x < 0 || p.x >= cols {
		return true
	}
	if p.y >= rows {
		return true
	}
	if p.y >= 0 && b[p.y][p.x] {
		return true
	}

	return false
}
