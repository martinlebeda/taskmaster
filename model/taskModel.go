package model

import (
	"database/sql"
	"time"
)

type Task struct {
	Id       int
	ParentId sql.NullInt64
	Prio     sql.NullString
	Status   string
	Desc     string
	DateIn   time.Time
	DateDone time.Time
	Url      sql.NullString
	Note     sql.NullString
	Script   sql.NullString
}
