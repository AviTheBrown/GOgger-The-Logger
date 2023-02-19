package GOgger

import "fmt"

type Logger struct {
	threshold Level
}

// Debugf -> the format of the string ex. %s, %d, Variadic arg.
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

// Infof -> the format of the string ex. %s, %d, Variadic arg.
func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

// Warnf -> the format of the string ex. %s, %d, Variadic arg.
func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

// Errorf -> the format of the string ex. %s, %d, Variadic arg.
func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

// Fatalf -> the format of the string ex. %s, %d, Variadic arg.
func (l *Logger) Fatalf(format string, args ...any) {
	if l.threshold > LevelFatal {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

// New -> returns a logger, ready to log at the required threshold.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}

func main() {

}
