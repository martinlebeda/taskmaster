package model

import "time"

type TimerDistance struct {
	Rowid    int
	Distance int
	Goal     time.Time
	Tag      string
	Note     string
}
