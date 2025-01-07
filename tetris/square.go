package tetris

// A bar represents a bar piece.
type square struct {
	*piece
}

// Creates a bar piece with a random colour, default starting position and direction.
func newSquare() *bar {
	piece := newPiece()
	piece.positon = point{3, 0}
	piece.matrix = make([][]bool, 2)
	piece.matrix[0] = []bool{true, true}
	piece.matrix[1] = []bool{true, true}

	return &bar{piece}
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
