// Copyright 2015 Qiang Xue. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/admpub/log"
)

func TestNewFileTarget(t *testing.T) {
	target := log.NewFileTarget()
	if target.MaxLevel != log.LevelDebug {
		t.Errorf("NewFileTarget.MaxLevel = %v, expected %v", target.MaxLevel, log.LevelDebug)
	}
	if target.Rotate != true {
		t.Errorf("NewFileTarget.Rotate = %v, expected %v", target.Rotate, true)
	}
	if target.BackupCount != 10 {
		t.Errorf("NewFileTarget.BackupCount = %v, expected %v", target.BackupCount, 10)
	}
	if target.MaxBytes != (1 << 20) {
		t.Errorf("NewFileTarget.MaxBytes = %v, expected %v", target.MaxBytes, 1<<20)
	}
}

func TestFileTarget(t *testing.T) {
	logFile := "app.log"
	os.Remove(logFile)

	logger := log.NewLogger()
	defer logger.Close()
	defer func() {
		if e := recover(); e != nil {
			t.Log(e)
		}
	}()
	logger.SetFatalAction(log.ActionPanic)
	target := log.NewFileTarget()
	target.FileName = logFile
	target.Categories = []string{"system.*"}
	logger.SetTarget(target)
	logger.Open()
	logger.Infof("t1: %v", 2)
	logger.GetLogger("system.db").Infof("t2: %v", 3)
	for i := 0; i < 100; i++ {
		logger.GetLogger("system.db").Infof(`async: %d`, i+1)
	}
	logger.GetLogger("system.db").Fatal(`fatal.file: `, time.Now())

	bytes, err := ioutil.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if strings.Contains(string(bytes), "t1: 2") {
		t.Errorf("Found unexpected %q", "t1: 2")
	}
	if !strings.Contains(string(bytes), "t2: 3") {
		t.Errorf("Expected %q not found", "t2: 3")
	}
}
