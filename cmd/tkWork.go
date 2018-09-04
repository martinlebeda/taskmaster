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

// tkDoneCmd represents the tkDone command
var tkWorkCmd = &cobra.Command{
	Use:     "work",
	Aliases: []string{"wrk", "wk", "w"},
	Args:    cobra.ExactArgs(1),
	Short:   "set status as work",
	Long: `usage: tm tk work [flags] ID
	
	Only one task should be in status work. Other task in work status, will be reset to status normal.`,
	Run: func(cmd *cobra.Command, args []string) {

		startWorkLog, err := cmd.Flags().GetBool("worklog")
		tools.CheckErr(err)

		worklogBefore, err := cmd.Flags().GetString("worklog-before")
		tools.CheckErr(err)
		worklogDate, err := cmd.Flags().GetString("worklog-date")
		tools.CheckErr(err)
		worklogTime, err := cmd.Flags().GetString("worklog-time")
		tools.CheckErr(err)

		id, err := strconv.Atoi(args[0])
		tools.CheckErr(err)
		task := service.TkGetById(id)

		fromEstimate, err := cmd.Flags().GetBool("estimate")
		tools.CheckErr(err)

		timerDuration := ""
		if fromEstimate {
			timerDuration = task.Estimate.String
		}

		timerDuration, err = cmd.Flags().GetString("timer")
		tools.CheckErr(err)

		taskDesc := service.RemovePrioFromDesc(task.Desc)
		if startWorkLog {
			service.WrkStart(taskDesc, worklogBefore, worklogTime, worklogDate)
			if listAfterChange {
				workList := service.WrkGetWork(now.BeginningOfDay(), now.EndOfDay(), false)
				termout.WrkListWork(workList)
			}
		}

		if timerDuration != "" {
			service.TmrAdd(viper.GetBool("timer-replace-tag"), viper.GetString("timer-tag"), timerDuration, taskDesc)
			if listAfterChange {
				service.TmrListAfterChange()
			}
			if viper.GetString("afterchange") != "" {
				service.SysAfterChange()
			}
		}

		var tsk model.Task
		tsk.Status = "W"
		tsk.DateDone = tools.GetZeroTime()
		service.TskResetWorkStatus()
		service.TskUpdate(tsk, args)

		if listAfterChange {
			service.TkListAfterChange()
		}
	},
}

func init() {
	taskCmd.AddCommand(tkWorkCmd)

	tkWorkCmd.Flags().BoolP("worklog", "w", false, "automatic start worklog with group, code and description from task")
	tkWorkCmd.Flags().StringP("timer", "t", "", "add new timer with desc from task (owerride -T)")
	tkWorkCmd.Flags().BoolP("estimate", "T", false, "add new timer with desc from task estimate")

	// options for worklog
	tkWorkCmd.Flags().StringP("worklog-before", "b", "", "Time shift worklog")
	curDate := time.Now()
	tkWorkCmd.Flags().String("worklog-time", curDate.Format("15:04"), "Time of begin worklog")
	tkWorkCmd.Flags().String("worklog-date", curDate.Format("2006-01-02"), "Time of begin worklog")

	// options for timer
	tkWorkCmd.Flags().String("timer-tag", "", "Timer tag for create timer from task")
	viper.BindPFlag("timer-tag", tkWorkCmd.Flags().Lookup("timer-tag"))

	tkWorkCmd.Flags().Bool("timer-replace-tag", false, "Replace other timers with tag when create timer from task")
	viper.BindPFlag("timer-replace-tag", tkWorkCmd.Flags().Lookup("timer-replace-tag"))
}
