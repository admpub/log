// Copyright 2015 Qiang Xue. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import "github.com/admpub/log"

func ExampleLogger_Error() {
	logger := log.NewLogger()

	logger.Targets = append(logger.Targets, log.NewConsoleTarget())

	logger.Open()

	// log without formatting
	logger.Error("a plain message")
	// log with formatting
	logger.Errorf("the value is: %v", 100)
}

func ExampleNewConsoleTarget() {
	logger := log.NewLogger()

	// creates a ConsoleTarget with color mode being disabled
	target := log.NewConsoleTarget()
	target.ColorMode = false

	logger.Targets = append(logger.Targets, target)

	logger.Open()

	// ... logger is ready to use ...
}

func ExampleNewFileTarget() {
	logger := log.NewLogger()

	// creates a FileTarget which keeps log messages in the app.log file
	target := log.NewFileTarget()
	target.FileName = "app.log"

	logger.Targets = append(logger.Targets, target)

	logger.Open()

	// ... logger is ready to use ...
}
