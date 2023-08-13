// Copyright 2015 Qiang Xue. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/admpub/log"
	"github.com/stretchr/testify/assert"
)

func TestNewFileTarget(t *testing.T) {
	target := log.NewFileTarget()
	if target.MaxLevel != log.LevelDebug {
		t.Errorf("NewFileTarget.MaxLevel = %v, expected %v", target.MaxLevel, log.LevelDebug)
	}
	if target.Rotate != true {
		t.Errorf("NewFileTarget.Rotate = %v, expected %v", target.Rotate, true)
	}
	if target.BackupCount != log.DefaultFileBackupCount {
		t.Errorf("NewFileTarget.BackupCount = %v, expected %v", target.BackupCount, 10)
	}
	if target.MaxBytes != log.DefaultFileMaxBytes {
		t.Errorf("NewFileTarget.MaxBytes = %v, expected %v", target.MaxBytes, log.DefaultFileMaxBytes)
	}
}

func TestFileTarget(t *testing.T) {
	logFile := "app.log"
	//logFile := "app-{date:" + time.RFC3339Nano + "}.log"
	os.Remove(logFile)

	logger := log.NewLogger()
	target := log.NewFileTarget()
	defer func() {
		logger.Close()
		if e := recover(); e != nil {
			t.Log(e)
		}
		expectedFileCount := target.BackupCount
		if target.CountFiles() != expectedFileCount {
			t.Errorf("NewFileTarget.CountFiles() = %v, expected %v", target.CountFiles(), expectedFileCount)
		}
		target.ClearFiles()
	}()
	target.FileName = logFile
	target.BackupCount = 8
	target.MaxBytes = 500
	target.Categories = []string{"system.*"}
	logger.SetTarget(target)
	logger.Open()
	logger.Infof("t1: %v", 2)
	logger.GetLogger("system.db").Infof("t2: %v", 3)
	for i := 0; i < 100; i++ {
		logger.GetLogger("system.db").Infof(`async: %d`, i+1)
	}
	logger.GetLogger("system.db").Fatal(`fatal.file: `, time.Now())

	bytes, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("Unexpected error: %v: %v", logFile, err)
	}

	if strings.Contains(string(bytes), "t1: 2") {
		t.Errorf("Found unexpected %q", "t1: 2")
	}
	if !strings.Contains(string(bytes), "t2: 3") {
		t.Errorf("Expected %q not found", "t2: 3")
	}
}

func TestSymlink(t *testing.T) {
	source := fmt.Sprintf(`%d.test`, time.Now().UnixMilli())
	err := os.WriteFile(source, []byte(source+"\n"), os.ModePerm)
	assert.NoError(t, err)
	os.Remove(`latest.test`)
	err = os.Symlink(source, `latest.test`)
	assert.NoError(t, err)
	err = os.Symlink(source, `latest.test`)
	assert.ErrorIs(t, err, os.ErrExist)
	err = log.ForceCreateSymlink(source, `latest.test`)
	assert.NoError(t, err)

	for i := 0; i < 10; i++ {
		source := fmt.Sprintf(`%d.test`, time.Now().UnixMilli())
		err := os.WriteFile(source, []byte(source+"\n"), os.ModePerm)
		assert.NoError(t, err)
		err = log.ForceCreateSymlink(source, `latest.test`)
		assert.NoError(t, err)
		time.Sleep(time.Second * 3)
	}

	// tail -F ./latest.test
}
