package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)  

var Version = "1.0.0"  

type Options struct {  
  Help bool  
  Silent bool  
  All bool  
  Version bool  
  Path bool  
  Cmd string  
}  

func ParseOptions() Options {  
  opts := Options{}  

  flag.BoolVar(&opts.Help, "help", false, "display help message")  
  flag.BoolVar(&opts.Silent, "s", false, "silent mode")  
  flag.BoolVar(&opts.All, "a", false, "display all paths")  
  flag.BoolVar(&opts.Version, "version", false, "display version")  
  flag.BoolVar(&opts.Path, "p", false, "check if path is executable")  

  flag.Parse()  

  if len(flag.Args()) > 0 {  
    opts.Cmd = flag.Arg(0)  
  }  

  return opts  
}  

func FindCommandPath(cmd string) ([]string, error) {  
  path, err := exec.LookPath(cmd)  
  if err != nil {  
    return nil, err  
  }  
  return filepath.SplitList(path), nil  
}  

func HandleOptions(opts Options) error {
  if opts.Help || opts.Cmd == "" {  
    return fmt.Errorf("Usage: exist [flags] <command>\nFlags:\n --help: display this help message\n -s: silent mode\n -a: display all paths\n --version: display version\n -p: check if path is executable")  
  }  

  if opts.Version {  
    fmt.Println("Version:", Version)  
    return nil
  }  

  if opts.Path {  
    info, err := os.Stat(opts.Cmd)  
    if os.IsNotExist(err) {  
      if !opts.Silent {  
        return fmt.Errorf("%s: No such file or directory", opts.Cmd)  
      }  
      return nil
    }  
    if !info.Mode().IsRegular() || (info.Mode().Perm()&0111) == 0 {  
      if !opts.Silent {  
        return fmt.Errorf("%s: Not an executable file", opts.Cmd)  
      }  
      return nil
    }  
    fmt.Println(opts.Cmd)  
    return nil
  }  

  paths, err := FindCommandPath(opts.Cmd)  
  if err != nil {  
    if !opts.Silent {  
      return fmt.Errorf("Error: %v", err)  
    }  
    return nil
  }  

  if opts.All {  
    for _, path := range paths {  
      fmt.Println(path)  
    }  
  } else {  
    fmt.Println(paths[0])  
  }  

  return nil
}  

func main() {  
  opts := ParseOptions()  

  err := HandleOptions(opts)
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}  
