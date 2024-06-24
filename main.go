package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	arch := runtime.GOOS
	fmt.Fprintf(os.Stdout, "(v0.2.2) Hola, %s!\n", arch)
}
