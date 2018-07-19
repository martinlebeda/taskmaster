package service

import (
	"fmt"
	. "github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/termout"
	"github.com/martinlebeda/taskmaster/tools"
	"strconv"
	"strings"
	"time"
)

func TskAdd(task Task) {

	task = nullTask(task)

	fmt.Println(task)

	db := OpenDB()
	stmt, err := db.Prepare("insert into task (desc, status, date_in, prio, url, note, script) values (?, ?, ?, ?, ?, ?, ?)")
	tools.CheckErr(err)
	result, err := stmt.Exec(task.Desc, "N", time.Now(), task.Prio, task.Url, task.Note, task.Script)
	tools.CheckErr(err)
	count, err := result.RowsAffected()
	tools.CheckErr(err)
	termout.Verbose("Task inserted: " + strconv.FormatInt(count, 10))

	/*
	   tkAddCmd.Flags().StringVarP(&taskOpt.Prio.String, "prio", "p", "", "task priority")
	       tkAddCmd.Flags().StringVarP(&taskOpt.Url.String, "url", "u", "", "url for this task (ie. sources on internet)")
	       tkAddCmd.Flags().StringVarP(&taskOpt.Note.String, "note", "n", "", "path to file with note")
	       tkAddCmd.Flags().StringVarP(&taskOpt.Script.String, "script", "s", "", "path to file with script")
	*/
}

// set null values for empty string or int is 0
func nullTask(task Task) Task {
	task.Prio.Valid = task.Prio.String != ""
	task.ParentId.Valid = task.ParentId.Int64 > 0
	task.Url.Valid = task.Url.String != ""
	task.Note.Valid = task.Note.String != ""
	task.Script.Valid = task.Script.String != ""

	return task
}

// delete tasks by list ID
func TskDel(ids []string) {
	// sql by field
	sql := "delete from task where id in (" + strings.Join(ids, ",") + ")"

	// execute delete
	db := OpenDB()
	stmt, err := db.Prepare(sql)
	tools.CheckErr(err)
	_, err = stmt.Exec()
	tools.CheckErr(err)
	termout.Verbose("Task deleted: ", strings.Join(ids, ","))
}

// update task non empty values in pattern for ids
func TskUpdate(task Task, ids []string) {
	sql := "update task set"

	// add parameters
	var setSql []string
	var argSql []interface{}
	if task.ParentId.Int64 > 0 {
		setSql = append(setSql, "parent_id = ?")
		argSql = append(argSql, task.ParentId.Int64)
	}
	if task.Prio.String != "" {
		setSql = append(setSql, "prio = ?")
		argSql = append(argSql, task.Prio.String)
	}
	if task.Status != "" {
		setSql = append(setSql, "status = ?")
		argSql = append(argSql, task.Prio.String)
	}
	if task.Desc != "" {
		setSql = append(setSql, "desc = ?")
		argSql = append(argSql, task.Desc)
	}

	if !task.DateDone.IsZero() {
		setSql = append(setSql, "date_done = ?")
		argSql = append(argSql, task.DateDone)
	} else {
		setSql = append(setSql, "date_done = ?")
		argSql = append(argSql, nil)
	}
	if task.Url.String != "" {
		setSql = append(setSql, "url = ?")
		argSql = append(argSql, task.Url.String)
	}
	if task.Note.String != "" {
		setSql = append(setSql, "note = ?")
		argSql = append(argSql, task.Note.String)
	}
	if task.Script.String != "" {
		setSql = append(setSql, "script = ?")
		argSql = append(argSql, task.Script.String)
	}

	sql += "set " + strings.Join(setSql, ", ")
	sql += " where id in (" + strings.Join(ids, ",") + ")"

	// execute update
	db := OpenDB()
	stmt, err := db.Prepare(sql)
	tools.CheckErr(err)
	_, err = stmt.Exec(argSql...)
	tools.CheckErr(err)
	termout.Verbose("Task updated: ", strings.Join(ids, ","))
}

func TskGetList() []Task {
	db := OpenDB()
	sql := "select id, parent_id, prio, status, desc, date_in, date_done, url, note, script from task where 1=1 "

	//if !timeFrom.IsZero() {
	//	sql += " and start >= ? and start <= ? "
	//}
	//
	//if onlyOpen {
	//sql += " and stop is null "
	//}

	sql += " order by CASE WHEN prio IS NULL THEN 'ZZ' ELSE prio END, date_in"

	rows, err := db.Query(sql)
	tools.CheckErr(err)

	var result []Task
	for rows.Next() {
		var task Task
		rows.Scan(&task.Id, &task.ParentId, &task.Prio, &task.Status, &task.Desc, &task.DateIn, &task.DateDone, &task.Url, &task.Note, &task.Script)
		result = append(result, task)
	}

	return result
}
