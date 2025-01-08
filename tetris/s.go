package tetris

// A s represents a s piece.
type s struct {
	*piece
}

// Creates a s piece with a random colour, default starting position and a matrix representative of it's blocks.
func newS() *s {
	piece := newPiece()
	piece.position = point{3, -2}
	piece.matrix = make([][]bool, 3)
	piece.matrix[0] = []bool{false, true, true}
	piece.matrix[1] = []bool{true, true, false}

	return &s{piece}
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
