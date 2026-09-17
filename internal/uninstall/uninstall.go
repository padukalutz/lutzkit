package uninstall

import (
	"fmt"
	"os"
	"path/filepath"

	"lutzkit/internal/project"
)

type Result struct {
	Type           project.Type
	PackageManager string
	Removed        []string
}

func Run(root string) (Result, error) {
	info := project.Detect(root)

	switch info.Type {
	case project.TypeNodeJS,
		project.TypeReact,
		project.TypeNextJS,
		project.TypeExpress:

		return runNode(root, info.PackageManager)

	case project.TypeGo:
		return runGo(root)

	default:
		return Result{}, fmt.Errorf(
			"unable to detect a supported project",
		)
	}
}

func runNode(root string, packageManager string) (Result, error) {
	nodeModules := filepath.Join(root, "node_modules")

	if _, err := os.Stat(nodeModules); err != nil {
		if os.IsNotExist(err) {
			return Result{
				Type:           project.TypeNodeJS,
				PackageManager: packageManager,
			}, nil
		}

		return Result{}, fmt.Errorf(
			"failed to check node_modules: %w",
			err,
		)
	}

	if err := os.RemoveAll(nodeModules); err != nil {
		return Result{}, fmt.Errorf(
			"failed to remove node_modules: %w",
			err,
		)
	}

	return Result{
		Type:           project.TypeNodeJS,
		PackageManager: packageManager,
		Removed: []string{
			"node_modules",
		},
	}, nil
}

func runGo(root string) (Result, error) {
	if _, err := os.Stat(filepath.Join(root, "go.mod")); err != nil {
		if os.IsNotExist(err) {
			return Result{}, fmt.Errorf("go.mod not found")
		}

		return Result{}, fmt.Errorf(
			"failed to check go.mod: %w",
			err,
		)
	}

	return Result{
		Type: project.TypeGo,
	}, nil
}
