package GOgger

type Logger struct {
    threshold Level
}

// (helper to Logger struct ) (the format of the string ex. %s, %d, Variadic arg)
func (l *Logger) Debugf(format string, args ...any) {
	// impl
}

// (helper to Logger struct ) (the format of the string ex. %s, %d, Variadic arg)
func (l *Logger) Infof(format string, args ...any) {
	// impl
}

// (helper to Logger struct ) (the format of the string ex. %s, %d, Variadic arg)
func (l *Logger) Errorf(format string, args ...any){
    // impl
}

// New returns a logger, ready to log at the required threshold.
func New(threshold Level) *Logger {
    return &Logger{
        threshold: threshold,
    }
}

func main() {


}
