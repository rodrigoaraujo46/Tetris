package tetris

// A bar represents a bar piece.
type bar struct {
	*piece
}

var barMatrix = [][][]bool{
	{
		{false, false, false, false},
		{false, false, false, false},
		{true, true, true, true},
		{false, false, false, false},
	},
	{
		{false, false, true, false},
		{false, false, true, false},
		{false, false, true, false},
		{false, false, true, false},
	},
}

// Creates a bar piece with a random colour, default starting position and a matrix representative of it's blocks.
func newBar() *bar {
	piece := newPiece()
	piece.position = point{3, -2}
	piece.matrix = barMatrix[0]

	return &bar{piece}
}

/*
func (b *bar) rotate(board board) {
	b.piece.clear()

	tempBar := *b
	tempBar.isVertical = !tempBar.isVertical
	if !tempBar.hasCollided(board) {
		*b = tempBar
	}

	fmt.Println(b.piece)
}
*/
