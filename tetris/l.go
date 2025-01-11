package tetris

// A l represents a l piece.
type l struct {
	*piece
}

var lMatrix = [][][]bool{
	{
		{false, false, false},
		{true, true, true},
		{true, false, false},
	},
	{
		{true, true, false},
		{false, true, false},
		{false, true, false},
	},
	{
		{false, false, true},
		{true, true, true},
		{false, false, false},
	},
	{
		{false, true, false},
		{false, true, false},
		{false, true, true},
	},
}

// Creates a l piece with a random colour, default starting position and a matrix representative of it's blocks.
func newL() *l {
	piece := newPiece()
	piece.position = point{3, -1}
	piece.matrix = lMatrix[0]

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
