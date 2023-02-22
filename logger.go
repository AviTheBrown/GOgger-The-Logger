package GOgger

import (
    "fmt"
    "io"
    "os"
)

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


    // output will implement the io.Writer so that the user can write to what ever they wish to write to
        // to log the messages that they have to a network, file, HTTP, etc.
    output io.Writer
}
// Debugf method-> takes no less that one argument
    // the Logger struct instance now can be used by
    // the Debugf method and manipulate its fields,
    // as its the first argument passed to the Debugf method
func(l *Logger) Debugf(format string, args ...any) {
    // This ensures that there is always a safe output stream to write to.
    if l.output == nil {
        l.output = os.Stdout
    }
    if l.threshold > LevelDebug {
        return
    }
    _, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

func(l *Logger) Infof(format string, args ...any) {
    // This ensures that there is always a safe output stream to write to.
    if l.output == nil {
        l.output = os.Stdout
    }
    if l.threshold > LevelInfo {
        return
    }
    _, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

func (l *Logger) Warnf(format string, args ...any) {
    // This ensures that there is always a safe output stream to write to.
    if l.output == nil {
        l.output = os.Stdout
    }
    if l.threshold > LevelWarn {
        return
    }
    _, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

func (l *Logger) Errorf(format string, args ...any) {
    // This ensures that there is always a safe output stream to write to.
    if l.output == nil {
        l.output = os.Stdout
    }
    if l.threshold > LevelError {
        return
    }
    _, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

func (l *Logger) Fatalf(format string, args ...any) {
    // This ensures that there is always a safe output stream to write to.
    if l.output == nil {
        l.output = os.Stdout
    }
    if l.threshold > LevelFatal {
        return
    }
    _, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

// log -> logs the message to the Stdout.
    // this is * not * exposed.
func (l *Logger) log(format string, args ...any) {
    _, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

// New -> takes one argument with of the type Level
    // and returns a pointer to a new Logger instance
    // with the threshold field filled with the threshold argument
    // and outout filled with the io.Writer
        // the default it Stdout.
    // this will correspond to the Level type.
    // *** this function invocation should be the first step in creating a Logger instance ***
func New(threshold Level, opts ...Option) *Logger {
    // lgr -> creates a new instance of Logger
    lgr := &Logger{threshold: threshold, output: os.Stdout}

    // the loop then iterates through the opts slice
        // and call each of the Option functions with the Logger instance lgr.
        // Each Option function modifies the lgr as needed, giving the user
        // the ablility to config the logger with as many args as needed.
    for _, configFunc := range opts {
        configFunc(lgr)
    }
    return lgr
}

func main(){

}