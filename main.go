package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"

	"github.com/firefart/nonamedreturns/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	a := analyzer.Analyzer
	a.Flags.Var(versionFlag{}, "V", "print version and exit")

	singlechecker.Main(a)
}

type versionFlag struct{}

func (versionFlag) IsBoolFlag() bool { return true }
func (versionFlag) Get() any         { return nil }
func (versionFlag) String() string   { return "" }
func (versionFlag) Set(s string) error {
	info, ok := debug.ReadBuildInfo()
	if ok {
		fmt.Fprintf(os.Stderr, "%s version %s built with %s (%s/%s)\n",
			filepath.Base(os.Args[0]), info.Main.Version, info.GoVersion, runtime.GOOS, runtime.GOARCH)
	}

	os.Exit(0)
	return nil
}
