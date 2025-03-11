package main

import (
	"fmt"
	"os"

	"github.com/nokamoto/yf/internal/cmd"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
