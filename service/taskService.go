// Copyright © 2018 Martin Lebeda <martin.lebeda@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package service

import (
	"github.com/jinzhu/now"
	. "github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/termout"
	"github.com/martinlebeda/taskmaster/tools"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func TskAdd(task Task) {

	task = prepareTask(task)

	db := OpenDB()
	stmt, err := db.Prepare("insert into task (desc, status, date_in) values (?, ?, ?)")
	tools.CheckErr(err)

	status := "N"
	if task.Status != "" {
		status = task.Status
	}

	result, err := stmt.Exec(task.Desc, status, time.Now())
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
	//task.Estimate.Valid = task.Estimate.String != ""

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

func TskPrio(priority string, ids []string) {
	db := OpenDB()

	// remove old priority
	sql := "update task set desc = substr(desc, 5) where id in (" + strings.Join(ids, ",") + ") and desc like '(_) %'"
	stmt, err := db.Prepare(sql)
	tools.CheckErr(err)
	_, err = stmt.Exec()
	tools.CheckErr(err)

	// add new priority
	if priority != "" {
		sql := "update task set desc = '(" + strings.ToUpper(priority) + ") '||desc where id in (" + strings.Join(ids, ",") + ")"
		stmt, err := db.Prepare(sql)
		tools.CheckErr(err)
		_, err = stmt.Exec()
		tools.CheckErr(err)
	}

	termout.Verbose("Task priority updated: ", strings.Join(ids, ","))
}

func TskAppend(part string, ids []string) {
	db := OpenDB()

	// add new part
	if part != "" {
		sql := "update task set desc = desc || ' " + part + "' where id in (" + strings.Join(ids, ",") + ") and not desc like '% " + part + "%'"
		stmt, err := db.Prepare(sql)
		tools.CheckErr(err)
		_, err = stmt.Exec()
		tools.CheckErr(err)
	}

	termout.Verbose("Task updated: ", strings.Join(ids, ","))
}

func TskRemove(part string, ids []string) {
	db := OpenDB()

	// remove old part
	sql := "update task set desc = trim(replace(desc, ' " + part + "', '')) where id in (" + strings.Join(ids, ",") + ") and desc like '% " + part + "%'"
	stmt, err := db.Prepare(sql)
	tools.CheckErr(err)
	_, err = stmt.Exec()
	tools.CheckErr(err)

	termout.Verbose("Task updated: ", strings.Join(ids, ","))
}

func TskGetList(doneFrom time.Time, showMaybe bool, showStatus string, showPrio []string, args []string) []Task {
	db := OpenDB()
	sql := "select id, status, desc, date_in, date_done from task where 1=1 "
	//sql := "select id, prio, code, category, status, desc, date_in, CASE WHEN date_done IS NULL THEN datetime('now') ELSE date_done END , url, note, estimate, script from task where 1=1 "

	//if !timeFrom.IsZero() {
	//	sql += " and start >= ? and start <= ? "
	//}
	//
	//if onlyOpen {
	//sql += " and stop is null "
	//}

	// usage showPrio
	if len(showPrio) > 0 {
		sql += " and ( "
		for _, prio := range showPrio {
			sql += "( desc like '(" + prio + ") %' ) or "
		}
		sql += " (1=0) ) "
	}

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
		err := rows.Scan(&task.Id, &task.Status, &task.Desc, &task.DateIn, &task.DateDoneRaw)
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

func TskGetById(id int) Task {
	db := OpenDB()
	sql := "select id, status, desc, date_in, date_done from task where id = ?"

	rows, err := db.Query(sql, id)
	tools.CheckErr(err)

	var result []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Status, &task.Desc, &task.DateIn, &task.DateDoneRaw)
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

func TskGetWork() Task {
	db := OpenDB()
	sql := "select id, status, desc, date_in, date_done from task where status = 'W'"

	rows, err := db.Query(sql)
	tools.CheckErr(err)

	var result []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Status, &task.Desc, &task.DateIn, &task.DateDoneRaw)
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

func TskRepairDoneDate() {
	db := OpenDB()
	db.Exec("update task set date_done = ? where status = 'X' and date_done is null", time.Now())
}

func TskListAfterChange() {
	termout.EmptyLineOut()
	tasks := TskGetList(now.BeginningOfDay(), false, "", []string{}, []string{})
	termout.TskListTasks(tasks)
}

func RemovePrioFromDesc(desc string) string {
	rp := regexp.MustCompile("^\\([A-Z]\\) ")
	s := rp.ReplaceAllString(desc, "")
	return s
}
func TskDone(args []string) {
	var tsk Task
	//tsk.Status = "N"
	//tsk.DateDone = tools.GetZeroTime()
	tsk.Status = "X"
	tsk.DateDone = time.Now()
	TskUpdate(tsk, args)
	// remove priority
	TskPrio("", args)
}
