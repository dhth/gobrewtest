package main

import (
	"fmt"
	"os"
	"runtime"
)

const version = "0.3.0"

func main() {
	fmt.Fprintf(os.Stdout, "(%s) Hola, %s!\n", version, runtime.GOOS)
}
