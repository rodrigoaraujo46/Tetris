package tetris

// A l represents a l piece.
type l struct {
	*piece
}

// Creates a l piece with a random colour, default starting position and a matrix representative of it's blocks.
func newL() *l {
	piece := newPiece()
	piece.position = point{3, -2}
	piece.matrix = make([][]bool, 3)
	piece.matrix[0] = []bool{false, false, true}
	piece.matrix[1] = []bool{true, true, true}

	return &l{piece}
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
