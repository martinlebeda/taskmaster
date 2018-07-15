package model

import "time"

type WorkList struct {
	Rowid    int
	Category string
	Code     string
	Desc     string
	Start    time.Time
	Stop     time.Time
}

type WorkSum struct {
	Desc    string
	Seconds int
}
