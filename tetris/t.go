package tetris

// A t represents a bar piece.
type t struct {
	*piece
}

var tMatrix = [][][]bool{
	{
		{false, false, false},
		{true, true, true},
		{false, true, false},
	},
	{
		{false, true, false},
		{true, true, false},
		{false, true, false},
	},
	{
		{false, true, false},
		{true, true, true},
		{false, false, false},
	},
	{
		{false, true, false},
		{false, true, true},
		{false, true, false},
	},
}

// Creates a t piece with a random colour, default starting position and a matrix representative of it's blocks.
func newT() *t {
	piece := newPiece()
	piece.position = point{3, -1}
	piece.matrix = tMatrix[0]

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
