package log

type Option func(*Logger)

// OptionLevel sets the severity level of a logger.
// The level parameter is a string which must match one of the following values:
// "Fatal", "Error", "Warn", "Okay", "Info", "Progress", "Debug".
// If a value other than the above is passed, the logger will not be configured.
// If the logger is not configured, it will not log any messages.
func OptionLevel(level string) Option {
	return func(logger *Logger) {
		logger.SetLevel(level)
	}
}

// OptionEmoji sets the emoji status of a logger.
// If true is passed, logger will attach emojis to log messages.
// If false is passed, logger will not attach emojis to log messages.
func OptionEmoji(on bool) Option {
	return func(logger *Logger) {
		logger.SetEmoji(on)
	}
}

// OptionFormatter sets the formatter of a logger.
// The formatter is used to format log messages.
// If nil is passed, the logger will use its default formatter.
func OptionFormatter(formatter Formatter) Option {
	return func(logger *Logger) {
		logger.SetFormatter(formatter)
	}
}

// OptionTarget sets the targets of a logger.
// The targets are used to process log messages.
// Each target can filter log messages by their severity level and category.
// The targets are also used to format log messages and write them to their destination storage.
func OptionTarget(targets ...Target) Option {
	return func(logger *Logger) {
		logger.SetTarget(targets...)
	}
}

// OptionFatalAction sets the action to be taken when a fatal log message is encountered.
// The action must be one of the following values:
// - ActionNothing: no action will be taken
// - ActionPanic: a panic will be triggered
// - ActionExit: the program will be terminated with an exit code of 1
func OptionFatalAction(action Action) Option {
	return func(logger *Logger) {
		logger.SetFatalAction(action)
	}
}

// OptionCategory sets the category of a logger.
// The category is used to categorize log messages.
// It can be used to filter log messages by their category.
// The category is also included in the formatted log message.
// The default category is "app".
func OptionCategory(category string) Option {
	return func(logger *Logger) {
		logger.Category = category
	}
}

// OptionCallStack sets the call stack configuration for a logger.
// It takes two parameters: level, which specifies the log level to
// configure the call stack for, and callStack, which specifies the
// call stack configuration.
//
// A nil callStack parameter will disable call stack logging for the
// specified level.
func OptionCallStack(level Level, callStack *CallStack) Option {
	return func(logger *Logger) {
		logger.SetCallStack(level, callStack)
	}
}
