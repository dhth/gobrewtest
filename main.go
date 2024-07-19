package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
)

var (
	version = "dev"
)

func main() {
	v := version
	if version == "dev" {
		info, ok := debug.ReadBuildInfo()
		if ok {
			v = info.Main.Version
		}
	}
	fmt.Fprintf(os.Stdout, "(%s) Hola, %s!\n", v, runtime.GOOS)
}
