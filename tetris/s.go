package tetris

import (
	"fmt"
)

// A s represents a s piece.
type s struct {
	piece
}

var sRotMatrix = [][][]bool{
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
	piece.matrix = sRotMatrix[piece.rotationIdx]

	return &s{*piece}
}

func (s *s) rotate(board board) {
	tempBar := *s
	tempBar.rotationIdx++
	tempBar.rotationIdx %= len(sRotMatrix)
	tempBar.matrix = sRotMatrix[s.rotationIdx]
	if !tempBar.hasCollided(board) {
		s.piece.clear()
		*s = tempBar
		fmt.Println(s.piece)
	}
}
