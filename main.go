package main

import (
	"fmt"
	"os"
	"os/exec"
)

func findCommandPath(cmd string) (string, error) {
	return exec.LookPath(cmd)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: which <command>")
		os.Exit(1)
	}
	
	cmd := os.Args[1]

	path, err := findCommandPath(cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", path)
}
