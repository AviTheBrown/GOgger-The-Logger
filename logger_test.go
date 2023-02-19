package GOgger_test

import (
	GOgger "Gogger"
)

func ExampleLogger_Debug() {
	dubugLogger := GOgger.New(GOgger.LevelDebug)
	dubugLogger.Debugf("Hello, %s", "world")
	// Output : Hello, world
}
