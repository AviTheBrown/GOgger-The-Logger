package GOgger_test

import GOgger "Gogger"

func ExampleLogger_Debugf() {
    debugLogger := GOgger.New(GOgger.LevelFatal)
    debugLogger.Fatalf("Hello %s", "world")
    // Output:
    // Hello world
}
