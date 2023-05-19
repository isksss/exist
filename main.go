package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Version of the application
var Version = "1.0.0"

// Options struct to hold command line options
type Options struct {
    Help    bool
    Silent  bool
    All     bool
    Version bool
    Path    bool
    Cmd     string
}

// ParseOptions parses command line options
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

// FindCommandPath returns all the paths where the command can be found
func FindCommandPath(cmd string) ([]string, error) {
    path, err := exec.LookPath(cmd)
    if err != nil {
        return nil, err
    }
    return filepath.SplitList(path), nil
}

func main() {
    opts := ParseOptions()

    if opts.Help || opts.Cmd == "" {
        fmt.Fprintln(os.Stderr, "Usage: which [flags] <command>")
        fmt.Fprintln(os.Stderr, "Flags:")
        fmt.Fprintln(os.Stderr, "  --help: display this help message")
        fmt.Fprintln(os.Stderr, "  -s: silent mode")
        fmt.Fprintln(os.Stderr, "  -a: display all paths")
        fmt.Fprintln(os.Stderr, "  --version: display version")
        fmt.Fprintln(os.Stderr, "  -p: check if path is executable")
        os.Exit(1)
    }

    if opts.Version {
        fmt.Println("Version:", Version)
        os.Exit(0)
    }

    if opts.Path {
        info, err := os.Stat(opts.Cmd)
        if os.IsNotExist(err) {
            if !opts.Silent {
                fmt.Fprintf(os.Stderr, "%s: No such file or directory\n", opts.Cmd)
            }
            os.Exit(1)
        }
        if !info.Mode().IsRegular() || (info.Mode().Perm()&0111) == 0 {
            if !opts.Silent {
                fmt.Fprintf(os.Stderr, "%s: Not an executable file\n", opts.Cmd)
            }
            os.Exit(1)
        }
        fmt.Println(opts.Cmd)
        os.Exit(0)
    }

    paths, err := FindCommandPath(opts.Cmd)
    if err != nil {
        if !opts.Silent {
            fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        }
        os.Exit(1)
    }

    if opts.All {
        for _, path := range paths {
            fmt.Println(path)
        }
    } else {
        fmt.Println(paths[0])
    }
}
