package tetris

import "fmt"

// A t represents a bar piece.
type t struct {
	piece
}

var tRotMatrix = [][][]bool{
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

// Creates a t piece with a random colour, default starting position and a matrix representative of it't blocks.
func newT() *t {
	piece := newPiece()
	piece.position = point{3, -1}
	piece.matrix = tRotMatrix[piece.rotationIdx]

	return &t{*piece}
}

func (t *t) rotate(board board) {
	tempBar := *t
	tempBar.rotationIdx++
	tempBar.rotationIdx %= len(tRotMatrix)
	tempBar.matrix = tRotMatrix[t.rotationIdx]
	if !tempBar.hasCollided(board) {
		t.piece.clear()
		*t = tempBar
		fmt.Println(t.piece)
	}
}
