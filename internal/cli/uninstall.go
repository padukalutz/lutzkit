package cli

import (
	"fmt"
	"os"

	"lutzkit/internal/ui"
	"lutzkit/internal/uninstall"
)

func uninstallProject() {
	ui.Header("Uninstall")

	root, err := os.Getwd()
	if err != nil {
		ui.Error("Failed to determine current directory.")
		return
	}

	result, err := uninstall.Run(root)
	if err != nil {
		ui.Error(err.Error())
		return
	}

	switch result.Type {
	case "Go":
		ui.Warning("Go dependencies are managed outside the project.")
		ui.Hint("LutzKit does not remove the global Go module cache.")
		return
	}

	if len(result.Removed) == 0 {
		ui.Warning("No installed dependencies found.")
		return
	}

	fmt.Println(ui.Bold("Detected Project"))
	fmt.Println("  " + string(result.Type))

	if result.PackageManager != "" {
		fmt.Println()
		fmt.Println(ui.Bold("Package Manager"))
		fmt.Println("  " + result.PackageManager)
	}

	fmt.Println()
	fmt.Println(ui.Bold("Removed"))

	for _, item := range result.Removed {
		fmt.Printf(
			"  %s %s\n",
			ui.SuccessText(ui.SymbolSuccess),
			item,
		)
	}

	ui.Blank()
	ui.Success("Project dependencies removed.")
}

func uninstallHelp() {
	ui.Header("uninstall")

	fmt.Println(ui.Bold("Remove installed project dependencies."))
	fmt.Println()

	fmt.Println(ui.Bold("Usage"))
	fmt.Println("  lutzkit uninstall")
	fmt.Println("  lutzkit uninstall --help")

	ui.Blank()

	ui.Hint("For Node.js projects, this removes node_modules.")
}
