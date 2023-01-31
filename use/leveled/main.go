package main

import (
	"time"

	ll "github.com/pushking812/loggy/logging/leveled"
)

func main() {
	logger := ll.New(time.RFC3339, true)

	logger.Log("info", "starting up service")
	logger.Log("warning", "no tasks found")
	logger.Log("error", "exiting: no work performed")

}
