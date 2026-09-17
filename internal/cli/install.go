package cli

import (
	"fmt"
	"os"

	"lutzkit/internal/install"
	"lutzkit/internal/ui"
)

func installProject() {
	ui.Header("Install")

	root, err := os.Getwd()

	if err != nil {
		ui.Error("Failed to determine current directory.")
		return
	}

	ui.Processing("Installing project dependencies...")

	result, err := install.Run(root)

	if err != nil {
		ui.Blank()
		ui.Error(err.Error())

		return
	}

	ui.Blank()
	ui.Success("Dependencies installed successfully.")

	fmt.Println()
	fmt.Println(ui.Bold("Package Manager"))
	fmt.Println("  " + result.PackageManager)

	fmt.Println()
	fmt.Println(ui.Bold("Command"))
	fmt.Println("  " + result.Command)

	ui.Blank()
}

func installHelp() {
	ui.Header("install")

	fmt.Println(ui.Bold("Install project dependencies."))
	fmt.Println()

	fmt.Println(ui.Bold("Usage"))
	fmt.Println("  lutzkit install")
	fmt.Println("  lutzkit install --help")

	ui.Blank()
}
