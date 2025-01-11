package tetris

// A z represents a z piece.
type z struct {
	*piece
}

var zMatrix = [][][]bool{
	{
		{false, false, false},
		{true, true, false},
		{false, true, true},
	},
	{
		{false, false, true},
		{false, true, true},
		{false, true, false},
	},
}

// Creates a z piece with a random colour, default starting position and a matrix representative of it's blocks.
func newZ() *z {
	piece := newPiece()
	piece.position = point{3, -1}
	piece.matrix = zMatrix[0]

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
