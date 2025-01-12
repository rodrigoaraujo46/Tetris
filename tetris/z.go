package tetris

import "fmt"

// A z represents a z piece.
type z struct {
	piece
}

var zRotMatrix = [][][]bool{
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
	piece.matrix = zRotMatrix[piece.rotationIdx]

	return &z{*piece}
}

func (z *z) rotate(board board) {
	tempBar := *z
	tempBar.rotationIdx++
	tempBar.rotationIdx %= len(zRotMatrix)
	tempBar.matrix = zRotMatrix[z.rotationIdx]
	if !tempBar.hasCollided(board) {
		z.piece.clear()
		*z = tempBar
		fmt.Println(z.piece)
	}
}
