package main

import (
	"fmt"

	ls "github.com/pushking812/loggy/logging/simple"
	ex1 "github.com/pushking812/loggy/use/simple/exemplar1"
	ex2 "github.com/pushking812/loggy/use/simple/exemplar2"
)

func main() {
	// Bug: variable 'simple.debug' is shared by different packages
	fmt.Println("initial: ", ls.IsDebug())
	ex1.Run(true)
	fmt.Println("after ex1: ", ls.IsDebug())
	ex2.Run(false)
	fmt.Println("after ex2: ", ls.IsDebug())
}
