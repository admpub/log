package log

var DefaultLog = &defaultLogger{Logger: New()}

type defaultLogger struct {
	*Logger
}

func GetLogger(category string, formatter ...Formatter) *Logger {
	return DefaultLog.GetLogger(category, formatter...)
}

func Sync(args ...bool) {
	DefaultLog.Sync(args...)
}

func SetTarget(targets ...Target) {
	DefaultLog.SetTarget(targets...)
}

func AddTarget(targets ...Target) {
	DefaultLog.AddTarget(targets...)
}

func SetLevel(level string) {
	DefaultLog.SetLevel(level)
}

func Fatalf(format string, a ...interface{}) {
	DefaultLog.Fatalf(format, a...)
}

func Errorf(format string, a ...interface{}) {
	DefaultLog.Errorf(format, a...)
}

func Warnf(format string, a ...interface{}) {
	DefaultLog.Warnf(format, a...)
}

func Infof(format string, a ...interface{}) {
	DefaultLog.Infof(format, a...)
}

func Debugf(format string, a ...interface{}) {
	DefaultLog.Debugf(format, a...)
}

func Fatal(a ...interface{}) {
	DefaultLog.Fatal(a...)
}

func Error(a ...interface{}) {
	DefaultLog.Error(a...)
}

func Warn(a ...interface{}) {
	DefaultLog.Warn(a...)
}

func Info(a ...interface{}) {
	DefaultLog.Info(a...)
}

func Debug(a ...interface{}) {
	DefaultLog.Debug(a...)
}
