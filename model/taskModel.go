package model

import (
	"database/sql"
	"time"
)

type Task struct {
	Id          int
	Prio        sql.NullString
	Code        sql.NullString
	Category    sql.NullString
	Status      string
	Desc        string
	DateIn      time.Time
	DateDoneRaw sql.NullString
	DateDone    time.Time
	Url         sql.NullString
	Note        sql.NullString
	Script      sql.NullString
	Estimate    sql.NullString
}
