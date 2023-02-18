package GOgger

// Level represents an available logging level
type Level byte

const (
    // LevelDebug -> represents the lowest level of the log,
    // mostly used for debugging purposes.
    LevelDebug Level = iota

    // LevelInfo -> represents a logging level that contains
    // information deem valuable.
    LevelInfo

    // LevelError -> represents the highest logging level,
    // only used to trace errors.
    LevelError
)

func main() {
    
}