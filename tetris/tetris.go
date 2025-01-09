package tetris

import (
	"bufio"
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
	fmt.Println(clearTerminal)

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

func getKeys(keyCh chan byte) {
	const (
		escape  = 27
		escape2 = 91
		ctrlC   = 3
	)

	byteCh := make(chan byte)
	go readByte(byteCh)

	byteNum := 0
	for byte := range byteCh {
		byteNum++
		switch byteNum {
		case 1:
			if byte != escape {
				byteNum = 0
				if byte == ctrlC {
					keyCh <- byte
				}
			}
		case 2:
			if byte != escape2 {
				byteNum = 0
			}
		case 3:
			keyCh <- byte
			byteNum = 0
		default:
			byteNum = 0
		}
	}
}

type keysPressed struct {
	up    bool
	right bool
	down  bool
	left  bool
	ctrlC bool
}

// This funtion handles user input and respectively set keysPressed during a frame.
func handleInput(keysChan chan keysPressed) {
	const (
		leftKey  = 68
		rightKey = 67
		downKey  = 66
		upKey    = 65
		ctrlC    = 3
	)

	keysP := keysPressed{}

	keyCh := make(chan byte)
	go getKeys(keyCh)

	for {
		select {
		case keysChan <- keysP:
			keysP = keysPressed{}
		case keyByte := <-keyCh:
			switch keyByte {
			case ctrlC:
				keysP.ctrlC = true
			case rightKey:
				keysP.right = true
			case leftKey:
				keysP.left = true
			case downKey:
				keysP.down = true
			case upKey:
				keysP.up = true
			}
		}
	}
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
	piece := newPieceManager()
	fmt.Println(piece)

	keysChan := make(chan keysPressed)
	go handleInput(keysChan)

	ticker, ticks := time.NewTicker(frameTime*time.Millisecond), 0
	for range ticker.C {
		ticks++
		if ticks == ticksPerG {
			ticks = 0
			if piece.move(*board, south) {
				piece.lock(board)
				piece = newPieceManager()
			}
		}
		keysPressed := <-keysChan
		if keysPressed.ctrlC {
			return
		}
		if piece.applyMoves(keysPressed, *board) {
			piece.lock(board)
			piece = newPieceManager()
		}

	}
}
