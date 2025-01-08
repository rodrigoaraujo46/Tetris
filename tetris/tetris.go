package tetris

import (
	"bufio"
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

type keysPressed struct {
	up    bool
	right bool
	down  bool
	left  bool
	space bool
}

func readByte(readCh chan byte) {
	reader := bufio.NewReader(os.Stdin)
	for {
		byte, err := reader.ReadByte()
		if err != nil {
			panic(err)
		}
		readCh <- byte
	}
}

func getKeyfromBuffer(buffer *bytes.Buffer, byte byte) byte {
	const (
		escape      = 27
		squareBrack = 91
	)

	if buffer.Len() > 3 {
		buffer.Reset()

		return 0
	}
	if buffer.Len() == 1 {
		if byte != escape {
			buffer.Reset()
		}

		return byte
	}
	if buffer.Len() == 2 {
		if byte != squareBrack {
			buffer.Reset()
		}

		return 0
	}

	return byte
}

// This funtion handles user input and respectively set keysPressed during a frame.
func handleInput(keysChan chan keysPressed, keysRequest, quit chan struct{}) {
	const (
		leftKey  = 68
		rightKey = 67
		downKey  = 66
		upKey    = 65
		ctrlC    = 3
		space    = 32
	)

	keysP := keysPressed{}

	readCh := make(chan byte)
	go readByte(readCh)

	buffer := bytes.NewBuffer(nil)
	for {
		select {
		case <-keysRequest:
			keysChan <- keysP
			keysP = keysPressed{}
		case byte := <-readCh:
			buffer.WriteByte(byte)

			keyByte := getKeyfromBuffer(buffer, byte)

			switch keyByte {
			case ctrlC:
				close(quit)
			case space:
				keysP.space = true
			case rightKey:
				keysP.right = true
			case leftKey:
				keysP.left = true
			case downKey:
				keysP.down = true
			case upKey:
			}
		}
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
	piece := newPieceManager()
	fmt.Println(piece)

	keysChan := make(chan keysPressed)
	keysRequest := make(chan struct{})
	quit := make(chan struct{})
	go handleInput(keysChan, keysRequest, quit)

	gravity := time.NewTicker(500 * time.Millisecond)
	defer gravity.Stop()

	ticker := time.NewTicker(16 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-quit:
			return
		case <-gravity.C:
			if ok := piece.move(*board, south); !ok {
				piece.addToBoard(board)
				piece = newPieceManager()
			}
		default:
			keysRequest <- struct{}{}
			keysPressed := <-keysChan
			piece.applyMoves(keysPressed, *board)
		}

	}
}
