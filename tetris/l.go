package tetris

import "fmt"

// A l represents a l piece.
type l struct {
	piece
}

var lRotMatrix = [][][]bool{
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
	piece.matrix = lRotMatrix[piece.rotationIdx]

	return &l{*piece}
}

func (l *l) rotate(board board) {
	tempBar := *l
	tempBar.rotationIdx++
	tempBar.rotationIdx %= len(lRotMatrix)
	tempBar.matrix = lRotMatrix[l.rotationIdx]
	if !tempBar.hasCollided(board) {
		l.piece.clear()
		*l = tempBar
		fmt.Println(l.piece)
	}
}
