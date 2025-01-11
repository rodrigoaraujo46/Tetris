package tetris

// A j represents a j piece.
type j struct {
	*piece
}

var jMatrix = [][][]bool{
	{
		{false, false, false},
		{true, true, true},
		{false, false, true},
	},
	{
		{false, true, false},
		{false, true, false},
		{true, true, false},
	},
	{
		{true, false, false},
		{true, true, true},
		{false, false, false},
	},
	{
		{false, true, true},
		{false, true, false},
		{false, true, false},
	},
}

// Creates a j piece with a random colour, default starting position and a matrix representative of it's blocks.
func newJ() *j {
	piece := newPiece()
	piece.position = point{3, -1}
	piece.matrix = jMatrix[0]

	return &j{piece}
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
