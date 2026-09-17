package prompt

import (
	"fmt"
	"os"

	"golang.org/x/term"

	"lutzkit/internal/ui"
)

const (
	keyUp    = "up"
	keyDown  = "down"
	keyEnter = "enter"
	keySpace = "space"
	keyCtrlC = "ctrl-c"
)

func Select(question string, options []string) (string, error) {
	if len(options) == 0 {
		return "", fmt.Errorf("no options available")
	}

	fd := int(os.Stdin.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return "", fmt.Errorf("failed to initialize terminal: %w", err)
	}

	defer term.Restore(fd, oldState)

	selected := 0

	render := func() {
		clearScreen()

		printLine(ui.Brand("◆ LutzKit") + ui.Muted(" / New Project"))
		printLine("")

		printLine(question)
		printLine("")

		for i, option := range options {
			if i == selected {
				printLine(fmt.Sprintf(
					"  %s %s",
					ui.Accent(ui.SymbolSelected),
					option,
				))
			} else {
				printLine(fmt.Sprintf("    %s", option))
			}
		}

		printLine("")
		printLine(fmt.Sprintf(
			"  %s",
			ui.Muted("↑↓ Navigate · Enter/Space Select · Ctrl+C Cancel"),
		))
	}

	render()

	buffer := make([]byte, 3)

	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			return "", err
		}

		key := parseKey(buffer[:n])

		switch key {
		case keyUp:
			selected--

			if selected < 0 {
				selected = len(options) - 1
			}

			render()

		case keyDown:
			selected++

			if selected >= len(options) {
				selected = 0
			}

			render()

		case keyEnter, keySpace:
			clearScreen()
			return options[selected], nil

		case keyCtrlC:
			clearScreen()
			return "", fmt.Errorf("cancelled")
		}
	}
}

func parseKey(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	if len(data) == 1 {
		switch data[0] {
		case 3:
			return keyCtrlC

		case 13, 10:
			return keyEnter

		case 32:
			return keySpace
		}
	}

	// Arrow keys:
	// Up   = ESC [ A
	// Down = ESC [ B
	if len(data) >= 3 &&
		data[0] == 27 &&
		data[1] == '[' {

		switch data[2] {
		case 'A':
			return keyUp

		case 'B':
			return keyDown
		}
	}

	return ""
}

func printLine(text string) {
	fmt.Print("\r")
	fmt.Print(text)
	fmt.Print("\r\n")
}

func clearScreen() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}
