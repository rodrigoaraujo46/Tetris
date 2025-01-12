package tetris

import (
	"fmt"
)

// A j represents a j piece.
type j struct {
	piece
}

var jRotMatrix = [][][]bool{
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
	piece.matrix = jRotMatrix[piece.rotationIdx]

	return &j{*piece}
}

func (j *j) rotate(board board) {
	tempBar := *j
	tempBar.rotationIdx++
	tempBar.rotationIdx %= len(jRotMatrix)
	tempBar.matrix = jRotMatrix[j.rotationIdx]
	if !tempBar.hasCollided(board) {
		j.piece.clear()
		*j = tempBar
		fmt.Println(j.piece)
	}
}
