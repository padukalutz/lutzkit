package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/padukalutz/lutzkit/internal/ui"
)

func Input(question string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println(question)
		fmt.Printf("  %s ", ui.Accent("❯"))

		value, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		value = strings.TrimSpace(value)

		if err := ValidateProjectName(value); err != nil {
			ui.Error(err.Error())
			continue
		}

		return value, nil
	}
}

func ValidateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	if name == "." || name == ".." {
		return fmt.Errorf("invalid project name")
	}

	for _, r := range name {
		if !isProjectNameChar(r) {
			return fmt.Errorf(
				"invalid project name: use only letters, numbers, `-`, or `_`",
			)
		}
	}

	return nil
}

func isProjectNameChar(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9') ||
		r == '-' ||
		r == '_'
}
