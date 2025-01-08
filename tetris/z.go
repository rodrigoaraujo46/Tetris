package tetris

// A z represents a z piece.
type z struct {
	*piece
}

// Creates a z piece with a random colour, default starting position and a matrix representative of it's blocks.
func newZ() *z {
	piece := newPiece()
	piece.position = point{3, -2}
	piece.matrix = make([][]bool, 3)
	piece.matrix[0] = []bool{true, true, false}
	piece.matrix[1] = []bool{false, true, true}

	return &z{piece}
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
