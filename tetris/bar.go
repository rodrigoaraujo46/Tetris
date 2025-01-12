package tetris

import "fmt"

// A bar represents a bar piece.
type bar struct {
	piece
}

var barRotMatrix = [][][]bool{
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
	piece.matrix = barRotMatrix[piece.rotationIdx]

	return &bar{*piece}
}

func (b *bar) rotate(board board) {
	tempBar := *b
	tempBar.rotationIdx++
	tempBar.rotationIdx %= len(barRotMatrix)
	tempBar.matrix = barRotMatrix[tempBar.rotationIdx]

	if !tempBar.hasCollided(board) {
		b.piece.clear()
		*b = tempBar
		fmt.Println(b.piece)
	}
}
