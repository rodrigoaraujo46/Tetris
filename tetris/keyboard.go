package tetris

import (
	"bufio"
	"os"
)

// keysPressed represents the keys pressed in each frame.
type keysPressed struct {
	up    bool
	right bool
	down  bool
	left  bool
	ctrlC bool
}

// Writes to a channel every byte read from stdin.
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

// Writes to a channel a byte representing that is the last byte of an important key that was pressed.
// A key is important if it used by our tetris game, these are:
//
//	leftKey  = 68
//	rightKey = 67
//	downKey  = 66
//	upKey    = 65
//	ctrlC    = 3
func getKeys(keyCh chan byte) {
	const (
		escape  = 27 // First byte of an escape sequence.
		escape2 = 91 // Second byte of an escape sequence.
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

// Sets keys pressed during every frame and writes the struct to the given channel
// when it is ready to receive.
func handleInput(keysChan chan keysPressed) {
	// Last byte of an important key.
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
