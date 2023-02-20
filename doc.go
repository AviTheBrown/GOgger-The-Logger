// Gogger is a logging library that can be use to log information that the user deems nessary. LevelDebug is of the less severaity and LevelFatal being that of the most significant.
// All Logger instances should be init with the New function as this acts a starting point and also returns a new Logger instance.
// The New function sets the Debug Level with the threshold field, this can be changed later on but this should not be done to prevent unwarranted effects.
// If a new log needs logging it would be best to create a new instance.
// The Level type is of type byte, this is used instead of int32 to save memory as the user system or production enviorment might be memory signifigant.
package GOgger

