package simple

import (
	"fmt"
	"time"
)

var debug bool

func Debug(b bool) {
	debug = b
}

func IsDebug() bool {
	return debug
}

func Log(statement string) {
	if !debug {
		return
	}

	fmt.Printf("Simple: %s %s\n", time.Now().Format(time.RFC3339), statement)
}
