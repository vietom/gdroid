package main

import (
	"os"
	"github.com/vietom/gdroid/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
