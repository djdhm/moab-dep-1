package lib

import (
	"fmt"

	display "github.com/djdhm/external-dep/lib"
)

const (
	Version = "v1.0.0"
)

func PrintVersion() {
	fmt.Println("I am moab-dep-1 version " + Version)
	fmt.Printf("External Lib version = %v\n", display.Version())
	display.Display()
}
