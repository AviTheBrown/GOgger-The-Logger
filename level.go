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

	// LevelWarn -> represents a log between that of basic infomation
	// and that of of and error.
	LevelWarn

	// LevelError -> represents the penultimate highest logging level,
	// only used to trace errors.
	LevelError

	// LevelFatal -> represents a log that will cause signifigant
	// bugs to the program
	LevelFatal
)
