package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("{{PROJECT_NAME}} CLI")

	if len(os.Args) > 1 {
		fmt.Println("Arguments:", os.Args[1:])
	}
}
