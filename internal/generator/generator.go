package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/padukalutz/lutzkit/internal/template"
)

type Project struct {
	Name     string
	Template string
	Variant  string
}

type Result struct {
	Files []string
}

func Generate(project Project) (Result, error) {
	definition, ok := template.Find(project.Template)
	if !ok {
		return Result{}, fmt.Errorf(
			"template %q not found",
			project.Template,
		)
	}

	templatePath, err := template.Resolve(
		definition,
		project.Variant,
	)
	if err != nil {
		return Result{}, err
	}

	files, err := generateFiles(
		project,
		templatePath,
	)
	if err != nil {
		return Result{}, err
	}

	return Result{
		Files: files,
	}, nil
}

func generateFiles(
	project Project,
	templatePath string,
) ([]string, error) {
	if _, err := template.Files.ReadDir(templatePath); err != nil {
		return nil, fmt.Errorf(
			"template directory does not exist: %s",
			templatePath,
		)
	}

	if _, err := os.Stat(project.Name); err == nil {
		return nil, fmt.Errorf(
			"directory %q already exists",
			project.Name,
		)
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf(
			"failed to check project directory: %w",
			err,
		)
	}

	if err := os.MkdirAll(project.Name, 0755); err != nil {
		return nil, fmt.Errorf(
			"failed to create project directory: %w",
			err,
		)
	}

	files, err := copyTemplate(
		templatePath,
		project.Name,
		project.Name,
	)

	if err != nil {
		_ = os.RemoveAll(project.Name)
		return nil, err
	}

	return files, nil
}

func copyTemplate(
	sourceDir string,
	targetDir string,
	projectName string,
) ([]string, error) {
	entries, err := template.Files.ReadDir(sourceDir)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to read template: %w",
			err,
		)
	}

	var files []string

	for _, entry := range entries {
		sourcePath := filepath.Join(
			sourceDir,
			entry.Name(),
		)

		fileName := entry.Name()

		if strings.HasSuffix(fileName, ".template") {
			fileName = strings.TrimSuffix(fileName, ".template")
		}

		targetPath := filepath.Join(
			targetDir,
			fileName,
		)

		if entry.IsDir() {
			if err := os.MkdirAll(
				targetPath,
				0755,
			); err != nil {
				return nil, fmt.Errorf(
					"failed to create directory %s: %w",
					targetPath,
					err,
				)
			}

			childFiles, err := copyTemplate(
				sourcePath,
				targetPath,
				projectName,
			)

			if err != nil {
				return nil, err
			}

			for _, file := range childFiles {
				files = append(
					files,
					filepath.Join(
						entry.Name(),
						file,
					),
				)
			}

			continue
		}

		content, err := template.Files.ReadFile(
			sourcePath,
		)

		if err != nil {
			return nil, fmt.Errorf(
				"failed to read %s: %w",
				sourcePath,
				err,
			)
		}

		contentString := strings.ReplaceAll(
			string(content),
			"{{PROJECT_NAME}}",
			projectName,
		)

		if err := os.WriteFile(
			targetPath,
			[]byte(contentString),
			0644,
		); err != nil {
			return nil, fmt.Errorf(
				"failed to create %s: %w",
				targetPath,
				err,
			)
		}

		relativePath, err := filepath.Rel(
			targetDir,
			targetPath,
		)

		if err != nil {
			return nil, fmt.Errorf(
				"failed to determine generated file path: %w",
				err,
			)
		}

		files = append(files, relativePath)
	}

	return files, nil
}
