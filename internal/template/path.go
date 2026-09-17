package template

import "fmt"

func Resolve(definition Definition, variant string) (string, error) {
	if variant == "" {
		return "templates/" + definition.Path, nil
	}

	selected, ok := FindVariant(definition, variant)
	if !ok {
		return "",
			fmt.Errorf(
				"variant %q not found for template %q",
				variant,
				definition.Name,
			)
	}

	return "templates/" +
		definition.Path +
		"/" +
		selected.Path, nil
}