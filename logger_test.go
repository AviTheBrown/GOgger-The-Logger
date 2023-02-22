package GOgger_test

import (
    GOgger "Gogger"
    "testing"
)

type testWriter struct {
    contents string }
// Write implements the io.Writer interface.
func (tw *testWriter) Write(p []byte) (n int, err error) {
tw.contents = tw.contents + string(p)
return len(p), nil }

const (
    debugMessage = "Why write I still all one, ever the same,"
    infoMessage = "And keep invention in a noted weed,"
    warnMessage = "Then take the chance by the thorns"
    errorMessage = "That every word doth almost tell my name,"
    fatalMessage = "Why thou"
)


func TestLogger_DebugfInforfErrorf(t *testing.T) {
    type testCase struct {
        threshold GOgger.Level
        expectedString string
    }

    tt := map[string]testCase{
        "debug": {
            threshold: GOgger.LevelDebug,
            expectedString: debugMessage +"\n" + infoMessage + "\n" + warnMessage + "\n" + errorMessage + "\n" + fatalMessage + "\n",
        },
        "info" : {
            threshold: GOgger.LevelInfo,
            expectedString: infoMessage + "\n" + warnMessage + "\n" + errorMessage + "\n" + fatalMessage + "\n",
        },
        "warn" : {
            threshold: GOgger.LevelWarn,
            expectedString: warnMessage + "\n" + errorMessage + "\n" + fatalMessage + "\n",
        },
        "error" : {
            threshold: GOgger.LevelError,
            expectedString: errorMessage + "\n" + fatalMessage + "\n",
        },
        "fatal" : {
            threshold: GOgger.LevelFatal,
            expectedString: fatalMessage +"\n",
        },
    }

    for name ,tc := range tt {
        t.Run(name, func(t *testing.T) {
            tw := &testWriter{}

            testedLogger := GOgger.New(tc.threshold, GOgger.WithOutput(tw))

            testedLogger.Debugf(debugMessage)
            testedLogger.Infof(infoMessage)
            testedLogger.Warnf(warnMessage)
            testedLogger.Errorf(errorMessage)
            testedLogger.Fatalf(fatalMessage)


            if tw.contents != tc.expectedString {
                t.Errorf("invlaid contents, expected: %q, got: %q", tc.expectedString, tw.contents)
            }
        })
    }
}