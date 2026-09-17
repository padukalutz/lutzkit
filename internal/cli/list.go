package cli

import (
	"fmt"
	"strings"

	"github.com/padukalutz/lutzkit/internal/template"
	"github.com/padukalutz/lutzkit/internal/ui"
)

func listTemplates() {
	ui.Header("Templates")

	definitions := template.All()

	if len(definitions) == 0 {
		ui.Warning("No templates available.")
		return
	}

	fmt.Println(ui.Bold("Templates"))
	fmt.Println()

	for i, definition := range definitions {
		fmt.Println("  " + ui.Accent(definition.Name))
		fmt.Println("    " + ui.Muted(definition.Description))

		if definition.PackageManager != "" {
			fmt.Println(
				"    " +
					ui.Muted("Package manager: ") +
					definition.PackageManager,
			)
		}

		if len(definition.Variants) > 0 {
			variants := make([]string, 0, len(definition.Variants))

			for _, variant := range definition.Variants {
				variants = append(variants, variant.Name)
			}

			fmt.Println(
				"    " +
					ui.Muted("Variants: ") +
					strings.Join(variants, ", "),
			)
		}

		if i < len(definitions)-1 {
			fmt.Println()
		}
	}

	ui.Blank()
	ui.Hint("Run `lutzkit new` to create a project.")
}

func listHelp() {
	ui.Header("list")

	fmt.Println(ui.Bold("List available templates."))
	fmt.Println()

	fmt.Println(ui.Bold("Usage"))
	fmt.Println("  lutzkit list")
	fmt.Println("  lutzkit list --help")

	ui.Blank()
}
