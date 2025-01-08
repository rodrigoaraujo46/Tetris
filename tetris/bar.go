package tetris

// A bar represents a bar piece.
type bar struct {
	*piece
}

// Creates a bar piece with a random colour, default starting position and a matrix representative of it's blocks.
func newBar() *bar {
	piece := newPiece()
	piece.position = point{3, -2}
	piece.matrix = make([][]bool, 4)
	piece.matrix[1] = []bool{true, true, true, true}

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
