// Copyright 2015 Qiang Xue. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
)

type consoleBrush func(string) string

func newConsoleBrush(format string) consoleBrush {
	return func(text string) string {
		return "\033[" + format + "m" + text + "\033[0m"
	}
}

var brushes = map[Level]consoleBrush{
	LevelDebug: newConsoleBrush("34"), // blue
	LevelInfo:  newConsoleBrush("32"), // green
	LevelWarn:  newConsoleBrush("33"), // yellow
	LevelError: newConsoleBrush("31"), // red
	LevelFatal: newConsoleBrush("35"), // magenta
}

// ConsoleTarget writes filtered log messages to console window.
type ConsoleTarget struct {
	*Filter
	ColorMode    bool      // whether to use colors to differentiate log levels
	Writer       io.Writer // the writer to write log messages
	close        chan bool
	ColorStrFunc func(Level) string
}

// NewConsoleTarget creates a ConsoleTarget.
// The new ConsoleTarget takes these default options:
// MaxLevel: LevelDebug, ColorMode: true, Writer: os.Stdout
func NewConsoleTarget() *ConsoleTarget {
	return &ConsoleTarget{
		Filter:    &Filter{MaxLevel: LevelDebug},
		ColorMode: true,
		Writer:    os.Stdout,
		close:     make(chan bool, 0),
		ColorStrFunc: func(_ Level) string {
			return `●`
		},
	}
}

// Open prepares ConsoleTarget for processing log messages.
func (t *ConsoleTarget) Open(io.Writer) error {
	t.Filter.Init()
	if t.Writer == nil {
		return errors.New("ConsoleTarget.Writer cannot be nil")
	}
	if runtime.GOOS == `windows` {
		t.ColorMode = false
	}
	return nil
}

// Process writes a log message using Writer.
func (t *ConsoleTarget) Process(e *Entry) {
	if e == nil {
		t.close <- true
		return
	}
	if !t.Allow(e) {
		return
	}
	var msg string
	if t.ColorMode {
		msg = t.Colored(e.Level, e.String())
	} else {
		msg = e.String()
	}
	fmt.Fprintln(t.Writer, msg)
}

func (t *ConsoleTarget) Colored(level Level, msg string) string {
	brush, ok := brushes[level]
	if ok {
		if t.ColorStrFunc != nil {
			return brush(t.ColorStrFunc(level)) + msg
		}
		return brush(msg)
	}
	return msg
}

// Close closes the console target.
func (t *ConsoleTarget) Close() {
	<-t.close
}
