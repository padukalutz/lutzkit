package main

import (
	"os"

	"github.com/padukalutz/lutzkit/internal/cli"
)

func main() {
	cli.Run(os.Args[1:])
}
