package main

import (
	"time"

	la "github.com/pushking812/loggy/logging/advanced"
)

func main() {
	logger := la.New(time.RFC3339, true)

	logger.Log("This is a debug statement...")
}
