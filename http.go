package log

import "fmt"

func NewHttpLevel(code int) *httpLevel {
	lvName := HTTPStatusLevelName(code)
	lv, ok := Levels[lvName]
	if !ok {
		lv = LevelInfo
	}
	return &httpLevel{Code: code, Leveler: lv}
}

type httpLevel struct {
	Code int
	Leveler
}

func (h httpLevel) Tag() string {
	return `[` + fmt.Sprint(h.Code) + `]`
}
