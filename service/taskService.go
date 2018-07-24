package service

import (
	"github.com/jinzhu/now"
	. "github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/termout"
	"github.com/martinlebeda/taskmaster/tools"
	"strconv"
	"strings"
	"time"
)

func TskAdd(task Task) {

	task = prepareTask(task)

	db := OpenDB()
	stmt, err := db.Prepare("insert into task (desc, status, date_in, prio, code, category, url, note, script) values (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	tools.CheckErr(err)
	result, err := stmt.Exec(task.Desc, "N", time.Now(), task.Prio, task.Code, task.Category, task.Url, task.Note, task.Script)
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
func prepareTask(task Task) Task {
	task.Prio.Valid = task.Prio.String != ""
	task.Code.Valid = task.Code.String != ""
	task.Category.Valid = task.Category.String != ""
	task.Url.Valid = task.Url.String != ""
	task.Note.Valid = task.Note.String != ""
	task.Script.Valid = task.Script.String != ""

	// TODO Lebeda - FIX: check PRIO only 1 character
	if task.Prio.Valid {
		task.Prio.String = strings.ToUpper(task.Prio.String)
	}

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
func TskUpdate(task Task, forcePriority bool, ids []string) {
	sql := "update task "

	task = prepareTask(task)

	// add parameters
	var setSql []string
	var argSql []interface{}

	if task.Prio.String != "" || forcePriority {
		setSql = append(setSql, "prio = ?")
		argSql = append(argSql, task.Prio.String)
	}
	if task.Category.String != "" {
		setSql = append(setSql, "category = ?")
		argSql = append(argSql, task.Category.String)
	}
	if task.Code.String != "" {
		setSql = append(setSql, "code = ?")
		argSql = append(argSql, task.Code.String)
	}
	if task.Status != "" {
		setSql = append(setSql, "status = ?")
		argSql = append(argSql, task.Status)
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

	// TODO Lebeda - check if setSql is not empty

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

func TskGetList(doneFrom time.Time, showMaybe bool, showPrio []string, showCode, showCategory string, args []string) []Task {
	db := OpenDB()
	sql := "select id, prio, code, category, status, desc, date_in, date_done, url, note, script from task where 1=1 "

	//if !timeFrom.IsZero() {
	//	sql += " and start >= ? and start <= ? "
	//}
	//
	//if onlyOpen {
	//sql += " and stop is null "
	//}

	sql += " and (status <> 'X' or date_done > ?) "
	if !showMaybe {
		sql += " and status <> 'M' "
	}
	if len(showPrio) > 0 {
		sql += " and prio in ('" + strings.Join(showPrio, "','") + "') "
	}
	if showCode != "" {
		sql += " and code = '" + showCode + "' "
	}
	if showCategory != "" {
		sql += " and category = '" + showCategory + "' "
	}

	if len(args) > 0 {
		sql += " and ( "

		var descWhere []string
		for _, val := range args {
			descWhere = append(descWhere, "desc like '%"+val+"%'")
		}
		sql += strings.Join(descWhere, " or ")

		sql += " ) "
	}

	sql += " order by CASE WHEN status = 'W' THEN 1 WHEN status = 'N' THEN 10 WHEN status = 'M' THEN 60 WHEN status = 'X' THEN 99 ELSE 80 END, " +
		" CASE WHEN (prio IS NULL) OR (PRIO = '') THEN 'W' ELSE prio END, " +
		" category, code, date_in"

	rows, err := db.Query(sql, doneFrom)
	tools.CheckErr(err)

	var result []Task
	for rows.Next() {
		var task Task
		rows.Scan(&task.Id, &task.Prio, &task.Code, &task.Category, &task.Status, &task.Desc, &task.DateIn, &task.DateDone, &task.Url, &task.Note, &task.Script)
		result = append(result, task)
	}

	return result
}

func TkGetById(id int) Task {
	db := OpenDB()
	sql := "select id, prio, code, category, status, desc, date_in, date_done, url, note, script from task where id = ?"

	rows, err := db.Query(sql, id)
	tools.CheckErr(err)

	var result []Task
	for rows.Next() {
		var task Task
		rows.Scan(&task.Id, &task.Prio, &task.Code, &task.Category, &task.Status, &task.Desc, &task.DateIn, &task.DateDone, &task.Url, &task.Note, &task.Script)
		result = append(result, task)
	}

	return result[0]
}

// reset status work to normal for all records
func TskResetWorkStatus() {
	db := OpenDB()
	db.Exec("update task set status = 'N' where status = 'W'")
}

func TkListAfterChange() {
	termout.EmptyLineOut()
	tasks := TskGetList(now.BeginningOfDay(), false, []string{}, "", "", []string{})
	termout.TskListTasks(tasks)
}
