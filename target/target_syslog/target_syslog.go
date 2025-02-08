//go:build !windows && !plan9
// +build !windows,!plan9

package target_syslog

import (
	"errors"
	"fmt"
	"io"
	"log/syslog"
	"runtime"

	"github.com/admpub/log"
)

type SyslogTarget struct {
	*log.Filter
	Writer *syslog.Writer
	close  chan bool
}

func NewSyslogTarget(prefix string) (*SyslogTarget, error) {
	w, err := syslog.New(syslog.LOG_CRIT, prefix)
	return &SyslogTarget{
		Filter: &log.Filter{MaxLevel: log.LevelDebug},
		Writer: w,
		close:  make(chan bool),
	}, err
}

func (t *SyslogTarget) Open(io.Writer) error {
	t.Filter.Init()
	if t.Writer == nil {
		return errors.New("SyslogTarget.Writer cannot be nil")
	}
	if runtime.GOOS == "windows" {
		return errors.New("SyslogTarget not supported on Windows")
	}
	if runtime.GOOS == "plan9" {
		return errors.New("SyslogTarget not supported on plan9")
	}
	return nil
}

func (t *SyslogTarget) Process(e *log.Entry) {
	if e == nil {
		t.close <- true
		return
	}
	if !t.Allow(e) {
		return
	}
	msg := e.String()
	_, err := t.Writer.Write([]byte(msg))
	if err != nil {
		fmt.Println("Failed to write syslog message:", err)
	}
}

func (t *SyslogTarget) Close() {
	<-t.close
}
