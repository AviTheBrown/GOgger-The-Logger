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

// testWriter ->is a struc that implements io.Writer
type testWriter struct {
        contents string
    }

    func (tw *testWriter) Write(p []byte) (n int, err error) {
        tw.contents = tw.contents + string(p)
        return len(p), nil
    }
