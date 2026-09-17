package template

type Metadata struct {
	Name           string
	Description    string
	PackageManager string
	Variant        string
}

func GetMetadata(definition Definition, variant string) Metadata {
	return Metadata{
		Name:           definition.Name,
		Description:    definition.Description,
		PackageManager: definition.PackageManager,
		Variant:        variant,
	}
}
