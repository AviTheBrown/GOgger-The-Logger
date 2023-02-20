package GOgger

import "fmt"

// Logger struct will hold information needed
    // to log information
type Logger struct {
    // the Logger must have a strict imprtance level attributed
        // to its message. User subjectivity but there is a general
        // ideal of what should be of great and lesser importance.

    // should start at Debug/trace and increase in criticalness up to Fatal.
// --------------------------------------------------------------

    // threshold -> represents the level of the logger that the user wishes
        // to access.
        // no need to expose this field. The internl struct of the logger of the library callers.
    threshold Level
}
// Debugf method-> takes no less that one argument
    // the Logger struct instance now can be used by
    // the Debugf method and manipulate its fields,
    // as its the first argument passed to the Debugf method
func(l *Logger) Debugf(format string, args ...any) {
    if l.threshold > LevelDebug {
        return
    }
    _, _ = fmt.Printf(format+"\n", args...)
}

func(l *Logger) Infof(format string, args ...any) {
    if l.threshold > LevelInfo {
        return
    }
    _, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Warnf(format string, args ...any) {
    if l.threshold > LevelWarn {
        return
    }
    _, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Errorf(format string, args ...any) {
    if l.threshold > LevelError {
        return
    }
    _, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Fatalf(format string, args ...any) {
    if l.threshold > LevelFatal {
        return
    }
    _, _ = fmt.Printf(format+"\n", args...)
}


// New -> takes one argument with of the type Level
    // and returns a pointer to a new Logger instance
    // with the threshold field filled with the threshold argument
    // this will correspond to the Level type.
    // *** this function invocation should be the first step in creating a Logger instance ***
func New(threshold Level) *Logger {
    // creates a new instance of the Logger struct
    // and initializes its fields using a struct literal.
    return &Logger{
        threshold: threshold,
    }
}

func main(){

}