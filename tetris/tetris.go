package tetris

import (
	"fmt"
	"time"
)

// Prints tetris ascii art.
func printLogo() {
	fmt.Println("                    ___________     __         .__\r")
	fmt.Println("                    \\__    ___/____/  |________|__| ______\r")
	fmt.Println("                      |    |_/ __ \\   __\\_  __ \\  |/  ___/\r")
	fmt.Println("                      |    |\\  ___/|  |  |  | \\/  |\\___ \\\r")
	fmt.Println("                      |____| \\___  >__|  |__|  |__/____  >\r")
	fmt.Println("                                 \\/                    \\/\r")
}

func LockAndNew(piece, nextPiece *piece, board *board) (*piece, *piece) {
	piece.lock(board)

	nextPiece.clear()
	piece = nextPiece
	piece.moveToBoard()

	nextPiece = newPiece()

	return piece, nextPiece
}

// Sets up game and starts the game loop.
func StartGame() {
	const (
		frameTime = 16
		ticksPerG = 48
	)

	oldState, err := prepTerminal()
	if err != nil {
		panic(err.Error())
	}
	defer resetTerminal(oldState)

	printLogo()

	board := makeBoard()
	fmt.Println(board)
	printNextWindow()

	piece := newPiece()
	piece.moveToBoard()
	fmt.Println(piece)

	nextPiece := newPiece()
	fmt.Println(nextPiece)

	keysChan := make(chan keysPressed)
	go handleInput(keysChan)

	ticker, ticks := time.NewTicker(frameTime*time.Millisecond), 0
	for range ticker.C {
		ticks++
		if ticks == ticksPerG {
			ticks = 0
			if piece.move(*board, south) {
				piece, nextPiece = LockAndNew(piece, nextPiece, board)
				if piece.hasCollided(*board) {
					break
				}
				fmt.Println(piece, nextPiece)
				continue
			}
		}

		keysPressed := <-keysChan
		if keysPressed.ctrlC {
			return
		}
		if piece.applyMoves(keysPressed, *board) {
			piece, nextPiece = LockAndNew(piece, nextPiece, board)
			if piece.hasCollided(*board) {
				break
			}
			fmt.Println(piece, nextPiece)
		}

	}
}
