package tetris

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

// Sets terminal to raw mode, hides cursor and clears terminal.
func prepTerminal() (*term.State, error) {
	const (
		clearTerminal = "\033[H\033[J"
		hideCursor    = "\033[?25l"
	)

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return nil, err
	}

	fmt.Print(clearTerminal)
	fmt.Print(hideCursor)

	return oldState, nil
}

// Sets terminal to the original state, un-hides cursor and clears terminal.
func resetTerminal(oldState *term.State) {
	const (
		clearTerminal = "\033[H\033[J"
		showCursor    = "\033[?25h"
	)

	fmt.Print(showCursor)
	fmt.Print(clearTerminal)

	term.Restore(int(os.Stdin.Fd()), oldState)
}

// Prints tetris ascii art.
func printLogo() {
	fmt.Println("      ___________     __         .__\r")
	fmt.Println("      \\__    ___/____/  |________|__| ______\r")
	fmt.Println("        |    |_/ __ \\   __\\_  __ \\  |/  ___/\r")
	fmt.Println("        |    |\\  ___/|  |  |  | \\/  |\\___ \\\r")
	fmt.Println("        |____| \\___  >__|  |__|  |__/____  >\r")
	fmt.Println("                   \\/                    \\/\r")
}

// This funtion handles user input and respectively calls fuctions.
func handleInput(board board, piece pieceManager, exit chan struct{}) {
	// Expected bytes and what keys the represent
	const (
		leftKey     = 67
		rightKey    = 68
		downKey     = 66
		escape      = 27
		squareBrack = 91
		ctrlC       = 3
		space       = 32
	)

	buffer := bytes.NewBuffer(nil)
	for {
		buf := make([]byte, 1)
		_, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}

		byte := buf[0]
		buffer.WriteByte(byte)

		if buffer.Len() > 3 {
			buffer.Reset()
			continue
		}

		if buffer.Len() == 1 {
			if byte != escape {
				buffer.Reset()
			}

			switch byte {
			case ctrlC:
				close(exit)
				return
			case space:
			}

			continue
		}
		if buffer.Len() == 2 {
			if byte != squareBrack {
				buffer.Reset()
			}
			continue
		}

		switch byte {
		case rightKey:
			piece.move(board, west)
		case leftKey:
			piece.move(board, east)
		case downKey:
			piece.move(board, south)
		}
		buffer.Reset()
	}
}

// Sets up game and starts the game loop.
func StartGame() {
	oldState, err := prepTerminal()
	if err != nil {
		panic(err.Error())
	}
	defer resetTerminal(oldState)
	printLogo()

	board := makeBoard()
	fmt.Println(board)
	piece := newPiece()
	fmt.Println(piece)

	exit := make(chan struct{})
	go handleInput(*board, piece, exit)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			piece.move(*board, south)
		case <-exit:
			return
		}
	}
}
