// Copyright 2015 Qiang Xue. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/admpub/log"
)

func TestNewLogger(t *testing.T) {
	logger := log.NewLogger()
	if logger.MaxLevel != log.LevelDebug {
		t.Errorf("NewLogger().MaxLevel = %v, expected %v", logger.MaxLevel, log.LevelDebug)
	}
	if logger.Category != "app" {
		t.Errorf("NewLogger().Category = %v, expected %v", logger.Category, "app")
	}
	if logger.CallStackDepth != 0 {
		t.Errorf("NewLogger().CallStackDepth = %v, expected %v", logger.CallStackDepth, 0)
	}
}

func TestGetLogger(t *testing.T) {
	formatter := func(*log.Logger, *log.Entry) string {
		return "test"
	}
	logger := log.NewLogger()
	logger1 := logger.GetLogger("testing")
	if logger1.Category != "testing" {
		t.Errorf("logger1.Category = %v, expected %v", logger1.Category, "testing")
	}
	logger2 := logger.GetLogger("routing", formatter)
	if logger2.Category != "routing" {
		t.Errorf("logger2.Category = %v, expected %v", logger2.Category, "routing")
	}
	if logger2.Formatter(logger2, nil) != "test" {
		t.Errorf("logger2.Formatter has an unexpected value")
	}
}

type MemoryTarget struct {
	*log.Filter
	entries []*log.Entry
	open    bool
	ready   chan bool
	Option1 string
	Option2 bool
}

func (m *MemoryTarget) Open(io.Writer) error {
	m.open = true
	m.entries = make([]*log.Entry, 0)
	return nil
}

func (m *MemoryTarget) Process(e *log.Entry) {
	if e == nil {
		m.ready <- true
	} else {
		m.entries = append(m.entries, e)
	}
}

func (t *MemoryTarget) Close() {
	<-t.ready
}

func TestLoggerLog(t *testing.T) {
	logger := log.NewLogger().Sync()
	target := &MemoryTarget{
		Filter: &log.Filter{MaxLevel: log.LevelDebug},
		ready:  make(chan bool, 0),
	}
	logger.SetTarget()
	if target.open {
		t.Errorf("target.open = %v, expected %v", target.open, false)
	}
	logger.SetTarget(target)
	logger.Open()
	if !target.open {
		t.Errorf("target.open = %v, expected %v", target.open, true)
	}

	logger.Logf(log.LevelInfo, "t0: %v", 1)
	logger.Debugf("t1: %v", 2)
	logger.Info("t2")
	logger.Warn("t3")
	logger.Error("t4")
	logger.Fatal("t5")

	logger.Close()

	if len(target.entries) != 6 {
		for i, v := range target.entries {
			fmt.Printf("%v.\t%#v\n", i, *v)
		}
		t.Errorf("len(target.entries) = %v, expected %v", len(target.entries), 6)
	}
	levels := ""
	messages := ""
	for i := 0; i < 6; i++ {
		levels += target.entries[i].Level.String() + ","
		messages += target.entries[i].Message + ","
	}
	expectedLevels := "Info,Debug,Info,Warn,Error,Fatal,"
	expectedMessages := "t0: 1,t1: 2,t2,t3,t4,t5,"
	if levels != expectedLevels {
		t.Errorf("levels = %v, expected %v", levels, expectedLevels)
	}
	if messages != expectedMessages {
		t.Errorf("messages = %v, expected %v", messages, expectedMessages)
	}
	//*
	consoleTarget := log.NewConsoleTarget()
	logger.Async().SetTarget(consoleTarget)
	for i := 0; i < 10; i++ {
		logger.Infof(`async: %d`, i+1)
	}
	logger.Fatal(`fatal.`)
	logger.Sync().Error(`end.`)
	logger.Close()
	//*/
}
