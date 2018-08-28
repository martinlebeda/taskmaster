package model

import (
	"database/sql"
	"time"
)

type Task struct {
	Id          int
	Status      string
	Desc        string
	DateIn      time.Time
	DateDoneRaw sql.NullString
	DateDone    time.Time
	Estimate    sql.NullString
}
