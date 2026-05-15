package main

import (
	"fmt"
	"os"
	"runtime"
)

const version = "0.4.1"

func main() {
	fmt.Fprintf(os.Stdout, "(%s) Hola, %s!\n", version, runtime.GOOS)
}
