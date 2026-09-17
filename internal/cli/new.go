package cli

import (
	"fmt"

	"github.com/padukalutz/lutzkit/internal/generator"
	"github.com/padukalutz/lutzkit/internal/prompt"
	"github.com/padukalutz/lutzkit/internal/template"
	"github.com/padukalutz/lutzkit/internal/ui"
)

func newProject() {
	ui.Header("New Project")

	definitions := template.All()

	if len(definitions) == 0 {
		ui.Error("No templates available.")
		return
	}

	templateNames := make([]string, 0, len(definitions))

	for _, definition := range definitions {
		templateNames = append(templateNames, definition.Name)
	}

	selectedTemplate, err := prompt.Select(
		"What do you want to build?",
		templateNames,
	)
	if err != nil {
		ui.Error(err.Error())
		return
	}

	definition, ok := template.Find(selectedTemplate)
	if !ok {
		ui.Error(fmt.Sprintf(
			"Template %q not found.",
			selectedTemplate,
		))
		return
	}

	variantNames := make([]string, 0, len(definition.Variants))

	for _, variant := range definition.Variants {
		variantNames = append(variantNames, variant.Name)
	}

	selectedVariant := ""

	if len(variantNames) > 0 {
		selectedVariant, err = prompt.Select(
			"Choose a variant:",
			variantNames,
		)
		if err != nil {
			ui.Error(err.Error())
			return
		}
	}

	name, err := prompt.Input("Project name:")
	if err != nil {
		ui.Error(err.Error())
		return
	}

	fmt.Println()
	ui.Processing("Creating project...")

	result, err := generator.Generate(generator.Project{
		Name:     name,
		Template: selectedTemplate,
		Variant:  selectedVariant,
	})
	if err != nil {
		ui.Blank()
		ui.Error(err.Error())
		return
	}

	ui.Blank()
	ui.Success("Project created successfully.")

	fmt.Println()
	fmt.Println(ui.Bold("Project"))
	fmt.Println("  " + name)

	fmt.Println()
	fmt.Println(ui.Bold("Template"))
	fmt.Println("  " + selectedTemplate)

	if selectedVariant != "" {
		fmt.Println()
		fmt.Println(ui.Bold("Variant"))
		fmt.Println("  " + selectedVariant)
	}

	if len(result.Files) > 0 {
		fmt.Println()
		fmt.Println(ui.Bold("Files"))

		for _, file := range result.Files {
			fmt.Printf(
				"  %s %s\n",
				ui.SuccessText(ui.SymbolSuccess),
				file,
			)
		}
	}

	ui.Blank()
	ui.Hint(
		fmt.Sprintf(
			"Next: cd %s",
			name,
		),
	)
}
