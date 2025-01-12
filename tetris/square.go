package tetris

// A square represents a square piece.
type square struct {
	piece
}

var squareRotMatrix = [][]bool{
	{true, true},
	{true, true},
}

// Creates a square piece with a random colour, default starting position and a matrix representative of it'square blocks.
func newSquare() *square {
	piece := newPiece()
	piece.position = point{3, 0}
	piece.matrix = squareRotMatrix

	return &square{*piece}
}

func (square *square) rotate(board board) {}
