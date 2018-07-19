package service

import (
	. "github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/termout"
	"github.com/martinlebeda/taskmaster/tools"
	"strconv"
	"strings"
	"time"
)

func WrkStop(before, timeOpt, dateOpt string) {
	db := OpenDB()
	stmt, err := db.Prepare("update work set stop = ? where stop is null")
	tools.CheckErr(err)
	dateTime := getDateTime(before, timeOpt, dateOpt)
	result, err := stmt.Exec(dateTime)
	tools.CheckErr(err)
	count, err := result.RowsAffected()
	tools.CheckErr(err)
	termout.Verbose("Count of stopped task: " + strconv.FormatInt(count, 10))
}

func getDateTime(before, timeOpt, dateOpt string) time.Time {
	result, err := time.Parse("2006-01-02 15:04", dateOpt+" "+timeOpt)
	tools.CheckErr(err)

	if before != "" {
		duration, err := time.ParseDuration(before)
		tools.CheckErr(err)
		result = result.Add(duration)
	}

	return result
}

func WrkStart(taskName string, category, code, before, timeOpt, dateOpt string) {
	WrkStop(before, timeOpt, dateOpt)

	db := OpenDB()
	stmt, err := db.Prepare("insert into work (desc, start, category, code) values (?, ?, ?, ?)")
	tools.CheckErr(err)
	dateTime := getDateTime(before, timeOpt, dateOpt)
	result, err := stmt.Exec(taskName, dateTime, category, code)
	tools.CheckErr(err)
	count, err := result.RowsAffected()
	tools.CheckErr(err)
	termout.Verbose("Task inserted: " + strconv.FormatInt(count, 10))
}

func WrkGetWork(timeFrom, timeTo time.Time, onlyOpen bool) []WorkList {
	db := OpenDB()
	sql := "select rowid, " +
		" CASE WHEN category IS NULL THEN '' ELSE category END," +
		" CASE WHEN code IS NULL THEN '' ELSE code END," +
		" CASE WHEN desc IS NULL THEN '' ELSE desc END," +
		" start, stop from work where 1=1 "

	if !timeFrom.IsZero() {
		sql += " and start >= ? and start <= ? "
	}

	if onlyOpen {
		sql += " and stop is null "
	}

	sql += " order by start "

	rows, err := db.Query(sql, timeFrom, timeTo)
	tools.CheckErr(err)

	var result []WorkList
	for rows.Next() {
		var workList WorkList
		rows.Scan(&workList.Rowid, &workList.Category, &workList.Code, &workList.Desc, &workList.Start, &workList.Stop)
		result = append(result, workList)
	}

	return result
}

func WrkDel(args []string) {
	// sql by field
	sql := ""
	sql = "delete from work where rowid in (" + strings.Join(args, ",") + ")"

	// execute delete
	db := OpenDB()
	stmt, err := db.Prepare(sql)
	tools.CheckErr(err)
	_, err = stmt.Exec()
	tools.CheckErr(err)
	termout.Verbose("Worklog deleted: ", strings.Join(args, ","))
}

func WrkUpdate(code, category, desc string, start, stop time.Time, ids []string) {
	sql := "update work set"

	// add parameters
	var setSql []string
	var argSql []interface{}
	if code != "" {
		setSql = append(setSql, "code = ?")
		argSql = append(argSql, code)
	}
	if category != "" {
		setSql = append(setSql, "category = ?")
		argSql = append(argSql, category)
	}
	if desc != "" {
		setSql = append(setSql, "desc = ?")
		argSql = append(argSql, desc)
	}

	if !start.IsZero() {
		setSql = append(setSql, "start = ?")
		argSql = append(argSql, start)
	}
	if !stop.IsZero() {
		setSql = append(setSql, "stop = ?")
		argSql = append(argSql, stop)
	}

	sql += "set " + strings.Join(setSql, ", ")
	sql += " where rowid in (" + strings.Join(ids, ",") + ")"

	// execute update
	db := OpenDB()
	stmt, err := db.Prepare(sql)
	tools.CheckErr(err)
	_, err = stmt.Exec(argSql...)
	tools.CheckErr(err)
	termout.Verbose("Worklog updated: ", strings.Join(ids, ","))
}

func WrkGetWorkSum(timeFrom, timeTo time.Time, sumByField string) []WorkSum {
	db := OpenDB()
	sql := "select " + "CASE WHEN " + sumByField + " IS NULL THEN '' ELSE " + sumByField + " END," +
		" sum(strftime('%s', CASE WHEN stop IS NULL THEN 'now' ELSE stop END, 'localtime') - strftime('%s', start, 'localtime')) as distance " +
		" from work "

	sql += " where start >= ? and start <= ? "

	sql += " group by " + sumByField
	sql += " order by " + sumByField

	rows, err := db.Query(sql, timeFrom, timeTo)
	tools.CheckErr(err)

	var result []WorkSum
	for rows.Next() {
		var workSum WorkSum
		rows.Scan(&workSum.Desc, &workSum.Seconds)
		result = append(result, workSum)
	}

	return result
}
