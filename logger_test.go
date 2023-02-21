package GOgger_test

import (
    GOgger "Gogger"
    "os"
)

func ExampleLogger_Debugf() {
    debugLogger := GOgger.New(
        GOgger.LevelFatal,
        GOgger.WithOutput(os.Stdout))
    debugLogger.Debugf("Hello %s", "world")
    // Output:
    // Hello world
}
