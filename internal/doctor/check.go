package doctor

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/padukalutz/lutzkit/internal/template"
)

type Status string

const (
	StatusOK      Status = "ok"
	StatusWarning Status = "warning"
	StatusError   Status = "error"
)

type Check struct {
	Name    string
	Status  Status
	Version string
	Message string
}

func Run() []Check {
	return []Check{
		checkCommand("Go", "go", "version"),
		checkCommand("Node.js", "node", "--version"),
		checkCommand("npm", "npm", "--version"),
		checkCommand("Git", "git", "--version"),
		checkTemplates(),
	}
}

func checkCommand(name string, command string, args ...string) Check {
	path, err := exec.LookPath(command)

	if err != nil {
		return Check{
			Name:    name,
			Status:  StatusError,
			Message: "Not installed",
		}
	}

	cmd := exec.Command(path, args...)

	output, err := cmd.Output()

	if err != nil {
		return Check{
			Name:    name,
			Status:  StatusWarning,
			Message: "Installed but version check failed",
		}
	}

	version := strings.TrimSpace(string(output))

	return Check{
		Name:    name,
		Status:  StatusOK,
		Version: version,
	}
}

func checkTemplates() Check {
	count := len(template.All())

	if count == 0 {
		return Check{
			Name:    "Templates",
			Status:  StatusError,
			Message: "No templates available",
		}
	}

	return Check{
		Name:    "Templates",
		Status:  StatusOK,
		Message: formatTemplateCount(count),
	}
}

func formatTemplateCount(count int) string {
	if count == 1 {
		return "1 template available"
	}

	return strings.Join(
		[]string{
			strconv.Itoa(count),
			"templates available",
		},
		" ",
	)
}
