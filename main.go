package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	cmd := os.Args[1]

	path, err := exec.LookPath(cmd)
	if err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("%s\n", path)
	}
}
