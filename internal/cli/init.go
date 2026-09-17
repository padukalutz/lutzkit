package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"lutzkit/internal/config"
	"lutzkit/internal/project"
	"lutzkit/internal/prompt"
	"lutzkit/internal/ui"
)

func initProject() {
	ui.Header("Init")

	root, err := os.Getwd()

	if err != nil {
		ui.Error("Failed to determine current directory.")
		return
	}

	info := project.Detect(root)

	if info.Type == project.TypeUnknown {
		ui.Warning("No supported project detected.")
		ui.Hint("Run `lutzkit new` to create a new project.")
		return
	}

	fmt.Println(ui.Bold("Project"))
	fmt.Println("  " + filepath.Base(root))

	fmt.Println()
	fmt.Println(ui.Bold("Type"))
	fmt.Println("  " + string(info.Type))

	if info.PackageManager != "" {
		fmt.Println()
		fmt.Println(ui.Bold("Package Manager"))
		fmt.Println("  " + info.PackageManager)
	}

	if len(info.Files) > 0 {
		fmt.Println()
		fmt.Println(ui.Bold("Detected Files"))

		for i, file := range info.Files {
			prefix := "├──"

			if i == len(info.Files)-1 {
				prefix = "└──"
			}

			fmt.Printf("  %s %s\n", prefix, file)
		}
	}

	ui.Blank()

	options := []string{
		"Git configuration",
		"Environment files",
		"Editor configuration",
		"LutzKit configuration",
	}

	selected, err := prompt.Select(
		"? What do you want to add?",
		options,
	)

	if err != nil {
		if err.Error() == "cancelled" {
			ui.Warning("Operation cancelled.")
			return
		}

		ui.Error(err.Error())
		return
	}

	switch selected {
	case "Git configuration":
		initGit(root)

	case "Environment files":
		initEnvironment(root)

	case "Editor configuration":
		initEditor(root)

	case "LutzKit configuration":
		initConfig(root)
	}
}

func initGit(root string) {
	path := filepath.Join(root, ".gitignore")

	if fileExists(path) {
		ui.Warning(".gitignore already exists.")
		return
	}

	content := `node_modules/
.env
.env.local
dist/
build/
.next/
coverage/
bin/
*.log
`

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		ui.Error("Failed to create .gitignore: " + err.Error())
		return
	}

	ui.Success("Created .gitignore")
}

func initEnvironment(root string) {
	path := filepath.Join(root, ".env.example")

	if fileExists(path) {
		ui.Warning(".env.example already exists.")
		return
	}

	content := `# Environment variables
# DATABASE_URL=
# PORT=
`

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		ui.Error("Failed to create .env.example: " + err.Error())
		return
	}

	ui.Success("Created .env.example")
}

func initEditor(root string) {
	path := filepath.Join(root, ".editorconfig")

	if fileExists(path) {
		ui.Warning(".editorconfig already exists.")
		return
	}

	content := `root = true

[*]
charset = utf-8
end_of_line = lf
insert_final_newline = true
indent_style = space
indent_size = 2

[*.go]
indent_style = tab
`

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		ui.Error("Failed to create .editorconfig: " + err.Error())
		return
	}

	ui.Success("Created .editorconfig")
}

func initConfig(root string) {
	path := filepath.Join(root, config.ConfigFileName)

	if fileExists(path) {
		ui.Warning(config.ConfigFileName + " already exists.")
		return
	}

	if err := config.Save(root, config.Default()); err != nil {
		ui.Error(err.Error())
		return
	}

	ui.Success("Created " + config.ConfigFileName)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func initHelp() {
	ui.Header("init")

	fmt.Println(ui.Bold("Initialize LutzKit configuration in an existing project."))
	fmt.Println()

	fmt.Println(ui.Bold("Usage"))
	fmt.Println("  lutzkit init")
	fmt.Println("  lutzkit init --help")

	ui.Blank()
}
