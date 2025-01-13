package tetris

import (
	"bufio"
	"os"
)

type keysPressed struct {
	up    bool
	right bool
	down  bool
	left  bool
	ctrlC bool
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
