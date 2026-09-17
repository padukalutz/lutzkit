package ui

import "fmt"

func Header(section string) {
	fmt.Println()
	fmt.Println(Brand("◆ LutzKit") + Muted(" / "+section))
	fmt.Println()
}

func Section(title string) {
	fmt.Println()
	fmt.Println(Bold(title))
	fmt.Println()
}

func Success(message string) {
	fmt.Printf("%s %s\n", SuccessText(SymbolSuccess), message)
}

func Error(message string) {
	fmt.Printf("%s %s\n", ErrorText(SymbolError), message)
}

func Warning(message string) {
	fmt.Printf("%s %s\n", WarningText(SymbolWarning), message)
}

func Processing(message string) {
	fmt.Printf("%s %s\n", Accent(SymbolProcessing), message)
}

func Hint(message string) {
	fmt.Println()
	fmt.Println(Muted(message))
}

func Blank() {
	fmt.Println()
}
