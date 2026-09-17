package cli

import (
	"fmt"

	"lutzkit/internal/ui"
	"lutzkit/internal/version"
)

func Run(args []string) {
	if len(args) == 0 {
		printHelp()
		return
	}

	switch args[0] {
	case "--help", "-h", "help":
		printHelp()

	case "--version", "-v", "version":
		printVersion()

	case "new":
		if hasHelp(args[1:]) {
			newHelp()
			return
		}
		newProject()

	case "init":
		if hasHelp(args[1:]) {
			initHelp()
			return
		}
		initProject()

	case "list":
		if hasHelp(args[1:]) {
			listHelp()
			return
		}
		listTemplates()

	case "search":
		if hasHelp(args[1:]) {
			searchHelp()
			return
		}

		if len(args) < 2 {
			searchHelp()
			return
		}

		searchTemplates(args[1])

	case "doctor":
		if hasHelp(args[1:]) {
			doctorHelp()
			return
		}
		doctorCheck()

	case "install":
		if hasHelp(args[1:]) {
			installHelp()
			return
		}
		installProject()

	case "uninstall":
		if hasHelp(args[1:]) {
			uninstallHelp()
			return
		}
		uninstallProject()

	case "update":
		if hasHelp(args[1:]) {
			updateHelp()
			return
		}
		updateProject()

	default:
		ui.Error(fmt.Sprintf("Unknown command: %s", args[0]))
		ui.Hint("Run `lutzkit --help` to see available commands.")
	}
}

func hasHelp(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}

	return false
}

func printVersion() {
	fmt.Printf("%s v%s\n", version.Name, version.Version)
}

func newHelp() {
	ui.Header("new")

	fmt.Println(ui.Bold("Create a new project from a LutzKit template."))
	fmt.Println()

	fmt.Println(ui.Bold("Usage"))
	fmt.Println("  lutzkit new")
	fmt.Println("  lutzkit new --help")

	ui.Blank()

	ui.Hint("Use the interactive selector to choose a template and variant.")
}

func printHelp() {
	ui.Header("Help")

	fmt.Println(ui.Bold("Usage"))
	fmt.Println("  lutzkit <command> [options]")

	fmt.Println()
	fmt.Println(ui.Bold("Commands"))
	fmt.Println()

	fmt.Printf(
		"  %-12s %s\n",
		"new",
		"Create a new project",
	)

	fmt.Printf(
		"  %-12s %s\n",
		"init",
		"Initialize LutzKit in an existing project",
	)

	fmt.Printf(
		"  %-12s %s\n",
		"list",
		"List available templates",
	)

	fmt.Printf(
		"  %-12s %s\n",
		"search",
		"Search available templates",
	)

	fmt.Printf(
		"  %-12s %s\n",
		"doctor",
		"Check development environment",
	)

	fmt.Printf(
		"  %-12s %s\n",
		"install",
		"Install project dependencies",
	)

	fmt.Printf(
		"  %-12s %s\n",
		"uninstall",
		"Remove installed project dependencies",
	)

	fmt.Printf(
		"  %-12s %s\n",
		"update",
		"Check LutzKit project configuration",
	)

	fmt.Println()
	fmt.Println(ui.Bold("Options"))
	fmt.Println()

	fmt.Println("  -h, --help       Show help")
	fmt.Println("  -v, --version    Show version")

	ui.Blank()
}
