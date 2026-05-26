package ui

import (
	"text/template"
	"time"
)

var (
	funcMap = template.FuncMap{
		"duration": Duration,
		"unix":     Unix,
	}
)

func Duration(stop, start int64) string { _ = "STUB: not implemented"; return "" }

func Unix(t int64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
