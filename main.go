package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func findCommandPath(cmd string) (string, error) {  
  return exec.LookPath(cmd)  
}  

func main() {  
  help := flag.Bool("help", false, "display help message")
  flag.Parse()
  
  if *help || len(flag.Args()) < 1 {  
    fmt.Fprintln(os.Stderr, "Usage: which <command>")  
    fmt.Fprintln(os.Stderr, "Flags:")
    fmt.Fprintln(os.Stderr, "  --help: display this help message")
    os.Exit(1)
  }  

  cmd := flag.Arg(0)

  path, err := findCommandPath(cmd)  
  if err != nil {  
    fmt.Fprintf(os.Stderr, "Error: %v\n", err)  
    os.Exit(1)  
  }  

  fmt.Printf("%s\n", path)  
}
