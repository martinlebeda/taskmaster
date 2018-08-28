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
	stmt, err := db.Prepare("insert into task (desc, status, date_in, estimate) values (?, ?, ?, ?)")
	tools.CheckErr(err)
	result, err := stmt.Exec(task.Desc, "N", time.Now(), task.Estimate)
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
	task.Estimate.Valid = task.Estimate.String != ""

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
	sql := "update task "

	task = prepareTask(task)

	// add parameters
	var setSql []string
	var argSql []interface{}

	if task.Estimate.String != "" {
		setSql = append(setSql, "estimate = ?")
		argSql = append(argSql, task.Estimate.String)
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
	// TODO Lebeda - check if setSql is not empty

	sql += "set " + strings.Join(setSql, ", ")
	sql += " where 1=1 "

	// TODO Lebeda - select by query
	//if selectByCategory {
	//	sql += " and category in ('" + strings.Join(ids, "','") + "')"
	//} else if selectByCode {
	//	sql += " and code in ('" + strings.Join(ids, "','") + "')"
	//} else {
	sql += " and id in (" + strings.Join(ids, ",") + ")"
	//}

	// TODO Lebeda - debug for write sql
	//fmt.Println(sql)

	// execute update
	db := OpenDB()
	stmt, err := db.Prepare(sql)
	tools.CheckErr(err)
	_, err = stmt.Exec(argSql...)
	tools.CheckErr(err)
	termout.Verbose("Task updated: ", strings.Join(ids, ","))
}

func TskGetList(doneFrom time.Time, showMaybe bool, showPrio []string, showCode, showCategory, showStatus string, args []string) []Task {
	db := OpenDB()
	sql := "select id, status, desc, date_in, date_done, estimate from task where 1=1 "
	//sql := "select id, prio, code, category, status, desc, date_in, CASE WHEN date_done IS NULL THEN datetime('now') ELSE date_done END , url, note, estimate, script from task where 1=1 "

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
	if showStatus != "" {
		sql += " and status = '" + showStatus + "' "
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

	sql += " order by CASE WHEN status = 'W' THEN 1 WHEN status = 'N' THEN 10 WHEN status = 'M' THEN 60 WHEN status = 'X' THEN 99 ELSE 80 END, desc, date_in"

	rows, err := db.Query(sql, doneFrom)
	tools.CheckErr(err)

	var result []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Status, &task.Desc, &task.DateIn, &task.DateDoneRaw, &task.Estimate)
		tools.CheckErr(err)

		// null value
		if task.DateDoneRaw.Valid {
			task.DateDone, err = time.Parse("2006-01-02T15:04:05.999999-07:00", task.DateDoneRaw.String)
			tools.CheckErr(err)
		}

		result = append(result, task)
	}

	//for _, row := range rows {
	//    column1 = row.Str(0)
	//    column2 = row.Int(1)
	//    column3 = row.Bool(2)
	//    column4 = row.Date(3)
	//    // etc...
	//}

	return result
}

func TkGetById(id int) Task {
	db := OpenDB()
	sql := "select id, status, desc, date_in, date_done, estimate from task where id = ?"

	rows, err := db.Query(sql, id)
	tools.CheckErr(err)

	var result []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Status, &task.Desc, &task.DateIn, &task.DateDoneRaw, &task.Estimate)
		tools.CheckErr(err)

		// null value
		if task.DateDoneRaw.Valid {
			task.DateDone, err = time.Parse("2006-01-02T15:04:05.999999-07:00", task.DateDoneRaw.String)
			tools.CheckErr(err)
		}

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
	tasks := TskGetList(now.BeginningOfDay(), false, []string{}, "", "", "", []string{})
	termout.TskListTasks(tasks)
}
