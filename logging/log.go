package logging

type Logger struct {
	timeFormat string
	debug      bool
}

func New(timeFormat string, debug bool) *Logger {
	return &Logger{
		timeFormat: timeFormat,
		debug:      debug,
	}
}

func (l Logger) IsDebug() bool {
	return l.debug
}

func (l Logger) GetTimeFormat() string {
	return l.timeFormat
}
