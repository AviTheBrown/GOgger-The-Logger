package  GOgger_test

import (
    GOgger "Gogger"
)

func ExampleLogger_Debug() {
	debugLogger := GOgger.New(GOgger.LevelDebug)
    debugLogger.Debugf("Hello, %s", "world")
    // Output: Hello, world
}
