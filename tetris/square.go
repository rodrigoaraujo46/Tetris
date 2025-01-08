package tetris

// A square represents a square piece.
type square struct {
	*piece
}

// Creates a square piece with a random colour, default starting position and a matrix representative of it's blocks.
func newSquare() *square{
	piece := newPiece()
	piece.position = point{3, -2}
	piece.matrix = make([][]bool, 2)
	piece.matrix[0] = []bool{true, true}
	piece.matrix[1] = []bool{true, true}

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
