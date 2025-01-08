package tetris

import "math/rand/v2"

// Interface that enables functions to work with all different kinds of pieces.
type pieceManager interface {
	String() string
	clear()
	move(board, direction) bool
	hasCollided(board) bool
	applyMoves(keysPressed, board)
	addToBoard(*board)
	//rotate(board)
}

// Generates a random piece.
// Returns pieceManager interface
func newPieceManager() pieceManager {
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
