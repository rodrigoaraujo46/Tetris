package tetris

// A j represents a j piece.
type j struct {
	*piece
}

// Creates a j piece with a random colour, default starting position and a matrix representative of it's blocks.
func newJ() *j {
	piece := newPiece()
	piece.position = point{3, -2}

	piece.matrix = make([][]bool, 3)
	piece.matrix[0] = []bool{true, false, false}
	piece.matrix[1] = []bool{true, true, true}

	return &j{piece}
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
