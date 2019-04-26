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

package cmd

import (
	"github.com/jinzhu/now"
	"github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/service"
	"github.com/martinlebeda/taskmaster/termout"
	"github.com/martinlebeda/taskmaster/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var taskOpt model.Task

// tkAddCmd represents the tkAdd command
var tkAddCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(1),
	Short: "add new task",
	Long: `usage: tm tk add 'task description'
	
	In task description can use priority, project and contexts in todo.txt format.`,
	Run: func(cmd *cobra.Command, args []string) {
		startWork, err := cmd.Flags().GetBool("work")
		tools.CheckErr(err)

		startWorkLog, err := cmd.Flags().GetBool("worklog")
		tools.CheckErr(err)

		currentDone, err := cmd.Flags().GetBool("current-done")
		tools.CheckErr(err)

		currentDeffer, err := cmd.Flags().GetBool("current-defer")

		taskOpt.Desc = args[0]
		insertedId := service.TskAdd(taskOpt)

		if currentDone {
			task := service.TskGetWork()
			service.TskDone([]string{strconv.Itoa(task.Id)})
		}

		if currentDeffer {
			task := service.TskGetWork()
			service.TskDefer([]string{strconv.Itoa(task.Id)})
		}

		if startWork {
			var tsk model.Task
			tsk.Status = "W"
			tsk.DateDone = tools.GetZeroTime()
			service.TskResetWorkStatus()
			service.TskUpdate(tsk, []string{strconv.FormatInt(insertedId, 10)})
		}

		if startWorkLog {
			worklogBefore, err := cmd.Flags().GetString("worklog-before")
			tools.CheckErr(err)
			worklogDate, err := cmd.Flags().GetString("worklog-date")
			tools.CheckErr(err)
			worklogTime, err := cmd.Flags().GetString("worklog-time")
			tools.CheckErr(err)

			task := service.TskGetWork()
			taskDesc := service.RemovePrioFromDesc(task.Desc)
			service.WrkStart(taskDesc, worklogBefore, worklogTime, worklogDate)
			if listAfterChange {
				workList := service.WrkGetWork(now.BeginningOfDay(), now.EndOfDay(), false)
				termout.WrkListWork(workList, true)
			}
		}

		if listAfterChange {
			service.TskListAfterChange()
		}
		if viper.GetString("afterchange") != "" {
			service.SysAfterChange()
		}

		if viper.GetString("exportafterchange") != "" {
			service.TskExportTasks(viper.GetString("exportafterchange"), []string{})
		}
	},
}

func init() {
	taskCmd.AddCommand(tkAddCmd)

	tkAddCmd.Flags().BoolP("work", "k", false, "start work on added task")

	tkAddCmd.Flags().BoolP("current-done", "d", false, "mark current opened task as done")
	tkAddCmd.Flags().BoolP("current-defer", "r", false, "mark current opened task as defered")

	tkAddCmd.Flags().BoolP("worklog", "w", false, "automatic start worklog with description from task")
	tkAddCmd.Flags().StringP("timer", "t", "", "add new timer with desc from task")

	// options for worklog
	tkAddCmd.Flags().StringP("worklog-before", "b", "", "Time shift worklog")
	curDate := time.Now()
	tkAddCmd.Flags().String("worklog-time", curDate.Format("15:04"), "Time of begin worklog")
	tkAddCmd.Flags().String("worklog-date", curDate.Format("2006-01-02"), "Time of begin worklog")

	// options for timer
	tkAddCmd.Flags().String("timer-tag", "", "Timer tag for create timer from task")
	viper.BindPFlag("timer-tag", tkAddCmd.Flags().Lookup("timer-tag"))

	tkAddCmd.Flags().Bool("timer-replace-tag", false, "Replace other timers with tag when create timer from task")
	viper.BindPFlag("timer-replace-tag", tkWorkCmd.Flags().Lookup("timer-replace-tag"))
}

// TODO Lebeda - add if not exists
