package service

import (
	. "github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/termout"
	"strconv"
	"time"
)

func WrkStop() {
	db := OpenDB()
	stmt, err := db.Prepare("update work set stop = ? where stop is null")
	CheckErr(err)
	result, err := stmt.Exec(time.Now()) // TODO Lebeda - add time shift
	CheckErr(err)
	count, err := result.RowsAffected()
	CheckErr(err)
	termout.Verbose("Count of stopped task: " + strconv.FormatInt(count, 10))
}

func WrkStart(taskName string) {
	WrkStop()

	db := OpenDB()
	stmt, err := db.Prepare("insert into work (desc, start) values (?, ?)")
	CheckErr(err)
	result, err := stmt.Exec(taskName, time.Now()) // TODO Lebeda - add time shift
	CheckErr(err)
	count, err := result.RowsAffected()
	CheckErr(err)
	termout.Verbose("Task inserted: " + strconv.FormatInt(count, 10))
}

func WrkGetWork() []WorkList {
	db := OpenDB()
	sql := "select rowid, " +
		" CASE WHEN category IS NULL THEN '' ELSE category END," +
		" CASE WHEN code IS NULL THEN '' ELSE code END," +
		" CASE WHEN desc IS NULL THEN '' ELSE desc END," +
		" start, stop from work order by start "

	rows, err := db.Query(sql)
	CheckErr(err)

	var result []WorkList
	for rows.Next() {
		var workList WorkList
		rows.Scan(&workList.Rowid, &workList.Category, &workList.Code, &workList.Desc, &workList.Start, &workList.Stop)
		result = append(result, workList)
	}

	return result
}
