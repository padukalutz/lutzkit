package template

import "strings"

type Definition struct {
	Name           string
	Path           string
	Description    string
	PackageManager string
	Variants       []Variant
}

type Variant struct {
	Name string
	Path string
}

var definitions = []Definition{
	{
		Name:           "HTML / CSS / JS",
		Path:           "html",
		Description:    "Static website with HTML, CSS, and JavaScript",
		PackageManager: "",
		Variants: []Variant{
			{Name: "Basic", Path: "basic"},
			{Name: "Landing", Path: "landing"},
			{Name: "Portfolio", Path: "portfolio"},
			{Name: "Blank", Path: "blank"},
		},
	},
	{
		Name:           "React",
		Path:           "react",
		Description:    "React application powered by Vite",
		PackageManager: "npm",
		Variants: []Variant{
			{Name: "Vite", Path: "vite"},
		},
	},
	{
		Name:           "Next.js",
		Path:           "next",
		Description:    "Next.js application",
		PackageManager: "npm",
		Variants: []Variant{
			{Name: "Basic", Path: "basic"},
		},
	},
	{
		Name:           "Express",
		Path:           "express",
		Description:    "Express.js backend application",
		PackageManager: "npm",
		Variants: []Variant{
			{Name: "Basic", Path: "basic"},
		},
	},
	{
		Name:           "Node.js",
		Path:           "node",
		Description:    "Node.js application",
		PackageManager: "npm",
		Variants: []Variant{
			{Name: "Basic", Path: "basic"},
		},
	},
	{
		Name:           "Go",
		Path:           "go",
		Description:    "Go application",
		PackageManager: "",
		Variants: []Variant{
			{Name: "Basic", Path: "basic"},
			{Name: "CLI", Path: "cli"},
			{Name: "API", Path: "api"},
		},
	},
}

func All() []Definition {
	result := make([]Definition, len(definitions))
	copy(result, definitions)

	return result
}

func Find(name string) (Definition, bool) {
	for _, definition := range definitions {
		if definition.Name == name {
			return definition, true
		}
	}

	return Definition{}, false
}

func FindVariant(definition Definition, name string) (Variant, bool) {
	for _, variant := range definition.Variants {
		if variant.Name == name {
			return variant, true
		}
	}

	return Variant{}, false
}

func Search(query string) []Definition {
	query = strings.ToLower(strings.TrimSpace(query))

	if query == "" {
		return nil
	}

	var results []Definition

	for _, definition := range definitions {
		if matchesDefinition(definition, query) {
			results = append(results, definition)
		}
	}

	return results
}

func matchesDefinition(definition Definition, query string) bool {
	if strings.Contains(
		strings.ToLower(definition.Name),
		query,
	) {
		return true
	}

	if strings.Contains(
		strings.ToLower(definition.Description),
		query,
	) {
		return true
	}

	if strings.Contains(
		strings.ToLower(definition.PackageManager),
		query,
	) {
		return true
	}

	for _, variant := range definition.Variants {
		if strings.Contains(
			strings.ToLower(variant.Name),
			query,
		) {
			return true
		}
	}

	return false
}
