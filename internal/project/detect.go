package project

import (
	"os"
	"path/filepath"
)

func Detect(root string) Info {
	info := Info{
		Type: TypeUnknown,
	}

	info.Files = detectFiles(root)

	switch {
	case fileExists(filepath.Join(root, "next.config.js")),
		fileExists(filepath.Join(root, "next.config.mjs")),
		fileExists(filepath.Join(root, "next.config.ts")):

		info.Type = TypeNextJS
		info.PackageManager = detectPackageManager(root)

	case fileExists(filepath.Join(root, "vite.config.js")),
		fileExists(filepath.Join(root, "vite.config.ts")),
		fileExists(filepath.Join(root, "vite.config.mjs")):

		info.Type = TypeReact
		info.PackageManager = detectPackageManager(root)

	case fileExists(filepath.Join(root, "go.mod")):

		info.Type = TypeGo

	case fileExists(filepath.Join(root, "package.json")):

		info.Type = detectNodeProject(root)
		info.PackageManager = detectPackageManager(root)
	}

	return info
}

func detectNodeProject(root string) Type {
	data, err := os.ReadFile(filepath.Join(root, "package.json"))

	if err != nil {
		return TypeNodeJS
	}

	content := string(data)

	switch {
	case containsDependency(content, "express"):
		return TypeExpress

	default:
		return TypeNodeJS
	}
}

func detectPackageManager(root string) string {
	switch {
	case fileExists(filepath.Join(root, "pnpm-lock.yaml")):
		return "pnpm"

	case fileExists(filepath.Join(root, "yarn.lock")):
		return "yarn"

	case fileExists(filepath.Join(root, "bun.lockb")),
		fileExists(filepath.Join(root, "bun.lock")):
		return "bun"

	case fileExists(filepath.Join(root, "package-lock.json")):
		return "npm"

	case fileExists(filepath.Join(root, "package.json")):
		return "npm"

	default:
		return ""
	}
}

func detectFiles(root string) []string {
	candidates := []string{
		"package.json",
		"package-lock.json",
		"pnpm-lock.yaml",
		"yarn.lock",
		"bun.lock",
		"bun.lockb",
		"go.mod",
		"go.sum",
		"vite.config.js",
		"vite.config.ts",
		"vite.config.mjs",
		"next.config.js",
		"next.config.mjs",
		"next.config.ts",
	}

	var result []string

	for _, name := range candidates {
		if fileExists(filepath.Join(root, name)) {
			result = append(result, name)
		}
	}

	return result
}

func containsDependency(content string, dependency string) bool {
	needle := `"` + dependency + `"`

	for i := 0; i+len(needle) <= len(content); i++ {
		if content[i:i+len(needle)] == needle {
			return true
		}
	}

	return false
}

func fileExists(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	return !info.IsDir()
}
