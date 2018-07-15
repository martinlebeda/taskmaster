package service

import (
	. "github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/termout"
	"strconv"
	"time"
)

func WrkStop(before, timeOpt, dateOpt string) {
	db := OpenDB()
	stmt, err := db.Prepare("update work set stop = ? where stop is null")
	CheckErr(err)
	dateTime := getDateTime(before, timeOpt, dateOpt)
	result, err := stmt.Exec(dateTime)
	CheckErr(err)
	count, err := result.RowsAffected()
	CheckErr(err)
	termout.Verbose("Count of stopped task: " + strconv.FormatInt(count, 10))
}

func getDateTime(before, timeOpt, dateOpt string) time.Time {
	result, err := time.Parse("2006-01-02 15:04", dateOpt+" "+timeOpt)
	CheckErr(err)

	if before != "" {
		duration, err := time.ParseDuration(before)
		CheckErr(err)
		result = result.Add(duration)
	}

	return result
}

func WrkStart(taskName string, category, code, before, timeOpt, dateOpt string) {
	WrkStop(before, timeOpt, dateOpt)

	db := OpenDB()
	stmt, err := db.Prepare("insert into work (desc, start, category, code) values (?, ?, ?, ?)")
	CheckErr(err)
	dateTime := getDateTime(before, timeOpt, dateOpt)
	result, err := stmt.Exec(taskName, dateTime, category, code)
	CheckErr(err)
	count, err := result.RowsAffected()
	CheckErr(err)
	termout.Verbose("Task inserted: " + strconv.FormatInt(count, 10))
}

func WrkGetWork(timeFrom, timeTo time.Time) []WorkList {
	db := OpenDB()
	sql := "select rowid, " +
		" CASE WHEN category IS NULL THEN '' ELSE category END," +
		" CASE WHEN code IS NULL THEN '' ELSE code END," +
		" CASE WHEN desc IS NULL THEN '' ELSE desc END," +
		" start, stop from work "

	sql += " where start >= ? and start <= ? "

	sql += " order by start "

	rows, err := db.Query(sql, timeFrom, timeTo)
	CheckErr(err)

	var result []WorkList
	for rows.Next() {
		var workList WorkList
		rows.Scan(&workList.Rowid, &workList.Category, &workList.Code, &workList.Desc, &workList.Start, &workList.Stop)
		result = append(result, workList)
	}

	return result
}
