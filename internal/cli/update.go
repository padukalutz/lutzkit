package cli

import (
	"fmt"
	"os"

	"lutzkit/internal/config"
	"lutzkit/internal/ui"
	"lutzkit/internal/version"
)

func updateProject() {
	ui.Header("Update")

	root, err := os.Getwd()

	if err != nil {
		ui.Error("Failed to determine current directory.")
		return
	}

	cfg, err := config.Load(root)

	if err != nil {
		ui.Error("LutzKit configuration not found.")
		ui.Hint("Run `lutzkit init` first.")
		return
	}

	fmt.Println(ui.Bold("LutzKit"))
	fmt.Println("  v" + version.Version)

	fmt.Println()
	fmt.Println(ui.Bold("Config"))
	fmt.Println("  " + config.ConfigFileName)

	fmt.Println()
	fmt.Println(ui.Bold("Config Version"))
	fmt.Printf("  %d\n", cfg.Version)

	ui.Blank()
	ui.Success("Project configuration is up to date.")
}

func updateHelp() {
	ui.Header("update")

	fmt.Println(ui.Bold("Check LutzKit project configuration."))
	fmt.Println()

	fmt.Println(ui.Bold("Usage"))
	fmt.Println("  lutzkit update")
	fmt.Println("  lutzkit update --help")

	ui.Blank()
}
