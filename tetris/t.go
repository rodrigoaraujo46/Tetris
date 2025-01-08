package tetris

// A t represents a bar piece.
type t struct {
	*piece
}

// Creates a t piece with a random colour, default starting position and a matrix representative of it's blocks.
func newT() *t {
	piece := newPiece()
	piece.position = point{3, -2}
	piece.matrix = make([][]bool, 3)
	piece.matrix[0] = []bool{false, true, false}
	piece.matrix[1] = []bool{true, true, true}

	return &t{piece}
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
