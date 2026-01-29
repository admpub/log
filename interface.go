package log

import (
	"fmt"

	"github.com/admpub/color"
)

type LevelSetter interface {
	SetLevel(string)
}

// Leveler 日志等级接口
type Leveler interface {
	fmt.Stringer
	IsEnabled(level Level) bool
	Int() int
	Tag() string
	Color() *color.Color
	Level() Level
}

var _ LevelSetter = (*Logger)(nil)
