package advanced

import (
	"fmt"
	"time"

	"github.com/pushking812/loggy/logging"
)

type AdvancedLogger struct {
	logging.Logger
}

func New(timeFormat string, debug bool) *AdvancedLogger {
	l := logging.New(timeFormat, debug)
	return &AdvancedLogger{*l}
}

func (l *AdvancedLogger) Log(s string) {
	debug := l.IsDebug()
	if !debug {
		return
	}

	timeFormat := l.GetTimeFormat()
	fmt.Printf("Advanced: %s %s\n", time.Now().Format(timeFormat), s)
}

// type debugVar bool
// type timeFormatVar bool

// func NewVar(timeFormat string, debug bool) (*bool, *string) {
// 	return &debug, &timeFormat
// }
