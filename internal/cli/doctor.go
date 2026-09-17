package cli

import (
	"fmt"

	"github.com/padukalutz/lutzkit/internal/doctor"
	"github.com/padukalutz/lutzkit/internal/ui"
)

func doctorCheck() {
	ui.Header("Doctor")

	fmt.Println(ui.Bold("Environment"))
	fmt.Println()

	checks := doctor.Run()

	hasError := false
	hasWarning := false

	for _, check := range checks {
		symbol := ui.SymbolSuccess
		text := check.Message

		switch check.Status {
		case doctor.StatusOK:
			symbol = ui.SymbolSuccess

			if check.Version != "" {
				text = check.Version
			}

		case doctor.StatusWarning:
			symbol = ui.SymbolWarning
			hasWarning = true

		case doctor.StatusError:
			symbol = ui.SymbolError
			hasError = true
		}

		fmt.Printf(
			"  %s %-14s %s\n",
			symbol,
			check.Name,
			text,
		)
	}

	fmt.Println()

	switch {
	case hasError:
		ui.Error("Environment has problems that need attention.")

	case hasWarning:
		ui.Warning("Environment is usable but has warnings.")

	default:
		ui.Success("Environment looks good.")
	}

	ui.Blank()
}

func doctorHelp() {
	ui.Header("doctor")

	fmt.Println(ui.Bold("Check the development environment."))
	fmt.Println()

	fmt.Println(ui.Bold("Usage"))
	fmt.Println("  lutzkit doctor")
	fmt.Println("  lutzkit doctor --help")

	ui.Blank()
}
