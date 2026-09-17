package ui

const (
	reset = "\033[0m"
	bold  = "\033[1m"

	cyan   = "\033[36m"
	green  = "\033[32m"
	yellow = "\033[33m"
	red    = "\033[31m"
	gray   = "\033[90m"
)

func Brand(text string) string {
	return cyan + bold + text + reset
}

func Accent(text string) string {
	return cyan + text + reset
}

func SuccessText(text string) string {
	return green + text + reset
}

func WarningText(text string) string {
	return yellow + text + reset
}

func ErrorText(text string) string {
	return red + text + reset
}

func Muted(text string) string {
	return gray + text + reset
}

func Bold(text string) string {
	return bold + text + reset
}
