package install

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"lutzkit/internal/project"
)

type Result struct {
	PackageManager string
	Command        string
}

func Run(root string) (Result, error) {
	info := project.Detect(root)

	switch info.Type {
	case project.TypeGo:
		return runGo(root)

	case project.TypeNodeJS,
		project.TypeReact,
		project.TypeNextJS,
		project.TypeExpress:

		return runNode(root, info.PackageManager)

	default:
		return Result{}, fmt.Errorf(
			"unable to detect a supported project",
		)
	}
}

func runGo(root string) (Result, error) {
	if !fileExists(filepath.Join(root, "go.mod")) {
		return Result{}, fmt.Errorf(
			"go.mod not found",
		)
	}

	if err := runCommand(
		root,
		"go",
		"mod",
		"tidy",
	); err != nil {
		return Result{}, err
	}

	return Result{
		PackageManager: "Go",
		Command:        "go mod tidy",
	}, nil
}

func runNode(root string, packageManager string) (Result, error) {
	if !fileExists(filepath.Join(root, "package.json")) {
		return Result{}, fmt.Errorf(
			"package.json not found",
		)
	}

	switch packageManager {
	case "pnpm":
		if err := runCommand(root, "pnpm", "install"); err != nil {
			return Result{}, err
		}

		return Result{
			PackageManager: "pnpm",
			Command:        "pnpm install",
		}, nil

	case "yarn":
		if err := runCommand(root, "yarn", "install"); err != nil {
			return Result{}, err
		}

		return Result{
			PackageManager: "yarn",
			Command:        "yarn install",
		}, nil

	case "bun":
		if err := runCommand(root, "bun", "install"); err != nil {
			return Result{}, err
		}

		return Result{
			PackageManager: "bun",
			Command:        "bun install",
		}, nil

	case "npm":
		if err := runCommand(root, "npm", "install"); err != nil {
			return Result{}, err
		}

		return Result{
			PackageManager: "npm",
			Command:        "npm install",
		}, nil

	default:
		return Result{}, fmt.Errorf(
			"no supported Node.js package manager found",
		)
	}
}

func runCommand(root string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = root
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf(
			"command %q failed: %w",
			command,
			err,
		)
	}

	return nil
}

func fileExists(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	return !info.IsDir()
}
