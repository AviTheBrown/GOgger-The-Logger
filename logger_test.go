package GOgger_test

import GOgger "Gogger"

func ExampleLogger_Debugf() {
    debugLogger := GOgger.New(GOgger.LevelDebug)
    debugLogger.Debugf("Hello &s", "world")
    // Output:
    // nooo
}
