package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln (os.Stdout, "This is an STDOUT message.")
	fmt.Fprintln (os.Stderr, "This is an STDERR message.")
}
