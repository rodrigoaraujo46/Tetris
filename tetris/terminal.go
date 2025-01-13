package tetris

import (
	"fmt"
	"os"

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
