package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	arch := runtime.GOOS
	fmt.Fprintf(os.Stdout, "(v0.1.3) Hola, %s!\n", arch)
}
