package cli

import (
	"fmt"
	"strings"

	"github.com/padukalutz/lutzkit/internal/template"
	"github.com/padukalutz/lutzkit/internal/ui"
)

func searchTemplates(query string) {
	query = strings.TrimSpace(query)

	if query == "" {
		ui.Error("Search query cannot be empty.")
		ui.Hint("Usage: lutzkit search <query>")
		return
	}

	results := template.Search(query)

	ui.Header("Search")

	fmt.Println(ui.Bold("Search"))
	fmt.Println("  " + query)
	fmt.Println()

	if len(results) == 0 {
		ui.Warning("No templates found.")
		return
	}

	fmt.Println(
		ui.Muted(fmt.Sprintf("%d template(s) found.", len(results))),
	)
	fmt.Println()

	for i, result := range results {
		fmt.Println("  " + ui.Accent(result.Name))
		fmt.Println("    " + ui.Muted(result.Description))

		if result.PackageManager != "" {
			fmt.Println(
				"    " +
					ui.Muted("Package manager: ") +
					result.PackageManager,
			)
		}

		if len(result.Variants) > 0 {
			variants := make([]string, 0, len(result.Variants))

			for _, variant := range result.Variants {
				variants = append(variants, variant.Name)
			}

			fmt.Println(
				"    " +
					ui.Muted("Variants: ") +
					strings.Join(variants, ", "),
			)
		}

		if i < len(results)-1 {
			fmt.Println()
		}
	}

	ui.Blank()
	ui.Hint("Run `lutzkit new` to create a project.")
}

func searchHelp() {
	ui.Header("search")

	fmt.Println(ui.Bold("Search available templates."))
	fmt.Println()

	fmt.Println(ui.Bold("Usage"))
	fmt.Println("  lutzkit search <query>")
	fmt.Println("  lutzkit search --help")

	ui.Blank()

	ui.Hint("Example: lutzkit search react")
}
