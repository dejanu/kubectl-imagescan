package main

import (
	"fmt"
	"os"

	cmd "github.com/dejanu/kubectl-imagescan/cmd/kubectl-imagescan"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
