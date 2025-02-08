package target

import (
	"strings"

	"github.com/admpub/log"
	"github.com/admpub/log/target/target_mail"
	"github.com/admpub/log/target/target_network"
)

func UseCommonTargets(levelName string, targetNames ...string) *log.Logger {
	log.DefaultLog.SetLevel(levelName)
	targets := []log.Target{}

	for _, targetName := range targetNames {
		ti := strings.SplitN(targetName, `:`, 2)
		var categories []string
		if len(ti) == 2 {
			targetName = ti[0]
			if len(ti[1]) > 0 {
				categories = strings.Split(ti[1], `,`)
			}
		}
		switch targetName {
		case "console":
			//输出到命令行
			consoleTarget := log.NewConsoleTarget()
			consoleTarget.ColorMode = log.DefaultConsoleColorize
			consoleTarget.Categories = categories
			targets = append(targets, consoleTarget)

		case "file":
			//输出到文件
			if log.DefaultLog.MaxLevel.Int() >= log.LevelInfo.Int() {
				fileTarget := log.NewFileTarget()
				fileTarget.FileName = `logs/{date:20060102}_info.log`
				fileTarget.Levels = map[log.Leveler]bool{log.LevelInfo: true}
				fileTarget.Categories = categories
				fileTarget.MaxBytes = log.DefaultFileMaxBytes
				targets = append(targets, fileTarget)
			}
			if log.DefaultLog.MaxLevel.Int() >= log.LevelWarn.Int() {
				fileTarget := log.NewFileTarget()
				fileTarget.FileName = `logs/{date:20060102}_warn.log` //按天分割日志
				fileTarget.Levels = map[log.Leveler]bool{log.LevelWarn: true}
				fileTarget.Categories = categories
				fileTarget.MaxBytes = log.DefaultFileMaxBytes
				targets = append(targets, fileTarget)
			}
			if log.DefaultLog.MaxLevel.Int() >= log.LevelError.Int() {
				fileTarget := log.NewFileTarget()
				fileTarget.FileName = `logs/{date:20060102}_error.log` //按天分割日志
				fileTarget.MaxLevel = log.LevelError
				fileTarget.Categories = categories
				fileTarget.MaxBytes = log.DefaultFileMaxBytes
				targets = append(targets, fileTarget)
			}
			if log.DefaultLog.MaxLevel == log.LevelDebug {
				fileTarget := log.NewFileTarget()
				fileTarget.FileName = `logs/{date:20060102}_debug.log`
				fileTarget.Levels = map[log.Leveler]bool{log.LevelDebug: true}
				fileTarget.Categories = categories
				fileTarget.MaxBytes = log.DefaultFileMaxBytes
				targets = append(targets, fileTarget)
			}

		case "mail":
			if log.DefaultLog.MaxLevel.Int() == log.LevelFatal.Int() || log.DefaultLog.MaxLevel.Int() >= log.LevelError.Int() {
				mailTarget := target_mail.NewMailTarget()
				mailTarget.MaxLevel = log.LevelError
				mailTarget.Categories = categories
				targets = append(targets, mailTarget)
			}

		case "network":
			netTarget := target_network.NewNetworkTarget()
			netTarget.Categories = categories
			targets = append(targets, netTarget)
		}
	}
	log.SetTarget(targets...)
	log.SetFatalAction(log.ActionExit)
	return log.DefaultLog.Logger
}
