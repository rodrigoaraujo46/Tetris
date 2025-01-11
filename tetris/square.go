package tetris

// A square represents a square piece.
type square struct {
	*piece
}

var squareMatrix = [][]bool{
	{true, true},
	{true, true},
}

// Creates a square piece with a random colour, default starting position and a matrix representative of it's blocks.
func newSquare() *square {
	piece := newPiece()
	piece.position = point{3, 0}
	piece.matrix = squareMatrix

	return &square{piece}
}

/*
func (b *bar) rotate(board board) {
	b.clear()

	tempBar := *b
	tempBar.isVertical = !tempBar.isVertical
	if !tempBar.hasCollided(board) {
		*b = tempBar
	}

	fmt.Println(b)
}
*/
