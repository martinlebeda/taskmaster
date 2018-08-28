package model

import "time"

type WorkList struct {
	Id    int
	Desc  string
	Start time.Time
	Stop  time.Time
}

type WorkSum struct {
	Desc    string
	Seconds int
}
