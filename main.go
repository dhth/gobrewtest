package main

import (
	"fmt"
	"os"
	"runtime"
)

const version = "0.4.2"

func main() {
	fmt.Fprintf(os.Stdout, "(%s) Hola, %s!\n", version, runtime.GOOS)
}
