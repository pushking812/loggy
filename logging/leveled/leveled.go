package leveled

import (
	"fmt"
	"strings"
	"time"

	"github.com/pushking812/loggy/logging"
)

type LeveledLogger struct {
	logging.Logger
}

func New(timeFormat string, debug bool) *LeveledLogger {
	l := logging.New(timeFormat, debug)
	return &LeveledLogger{*l}
}

func (l *LeveledLogger) Log(level string, s string) {
	level = strings.ToLower(level)
	debug := l.IsDebug()
	switch level {
	case "info", "warning":
		if debug {
			l.write(level, s)
		}
	default:
		l.write(level, s)
	}
}

func (l *LeveledLogger) write(level string, s string) {
	timeFormat := l.GetTimeFormat()
	fmt.Printf("Leveled: [%s] %s %s\n", level, time.Now().Format(timeFormat), s)
}
