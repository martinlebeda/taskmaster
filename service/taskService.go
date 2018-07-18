package service

import (
	"fmt"
	. "github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/termout"
	"strconv"
	"time"
)

func TskAdd(task Task) {

	task = nullTask(task)

	fmt.Println(task)

	db := OpenDB()
	stmt, err := db.Prepare("insert into task (desc, status, date_in, prio, url, note, script) values (?, ?, ?, ?, ?, ?, ?)")
	CheckErr(err)
	result, err := stmt.Exec(task.Desc, "N", time.Now(), task.Prio, task.Url, task.Note, task.Script)
	CheckErr(err)
	count, err := result.RowsAffected()
	CheckErr(err)
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
