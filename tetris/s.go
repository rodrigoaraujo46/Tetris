package tetris

// A s represents a s piece.
type s struct {
	*piece
}

var sMatrix = [][][]bool{
	{
		{false, false, false},
		{false, true, true},
		{true, true, false},
	},
	{
		{false, true, false},
		{false, true, true},
		{false, false, true},
	},
}

// Creates a s piece with a random colour, default starting position and a matrix representative of it's blocks.
func newS() *s {
	piece := newPiece()
	piece.position = point{3, -1}
	piece.matrix = sMatrix[0]

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
