package GOgger

import "io"

type Option func(*Logger)

// WithOutput -> creates a Option function that takes a pointer of Logger
    // and uses closure to fill the output field of the Logger instance.
    // This can be used with the New function as variadic argumets of the
    // same type.
func WithOutput(output io.Writer) Option {
    return func(l *Logger) {
        l.output = output
    }
}