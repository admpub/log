package log

var defaultLog = &defaultLogger{Logger: New()}

type defaultLogger struct {
	*Logger
}

func GetLogger(category string, formatter ...Formatter) *Logger {
	return defaultLog.GetLogger(category, formatter...)
}

func SetTarget(targets ...Target) {
	defaultLog.SetTarget(targets...)
}

func AddTarget(targets ...Target) {
	defaultLog.AddTarget(targets...)
}

func SetLevel(level string) {
	defaultLog.SetLevel(level)
}

func Fatalf(format string, a ...interface{}) {
	defaultLog.Fatalf(format, a...)
}

func Errorf(format string, a ...interface{}) {
	defaultLog.Errorf(format, a...)
}

func Warnf(format string, a ...interface{}) {
	defaultLog.Warnf(format, a...)
}

func Infof(format string, a ...interface{}) {
	defaultLog.Infof(format, a...)
}

func Debugf(format string, a ...interface{}) {
	defaultLog.Debugf(format, a...)
}

func Fatal(a ...interface{}) {
	defaultLog.Fatal(a...)
}

func Error(a ...interface{}) {
	defaultLog.Error(a...)
}

func Warn(a ...interface{}) {
	defaultLog.Warn(a...)
}

func Info(a ...interface{}) {
	defaultLog.Info(a...)
}

func Debug(a ...interface{}) {
	defaultLog.Debug(a...)
}
