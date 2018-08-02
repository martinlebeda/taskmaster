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
	// TODO Lebeda - add long description
	//Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		timerDuration, err := cmd.Flags().GetString("timer")
		tools.CheckErr(err)

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

		if startWorkLog {
			service.WrkStart(task.Desc, task.Category.String, task.Code.String, worklogBefore, worklogTime, worklogDate)
			if listAfterChange {
				workList := service.WrkGetWork(now.BeginningOfDay(), now.EndOfDay(), false)
				termout.WrkListWork(workList)
			}
		}

		if timerDuration != "" {
			service.TmrAdd(false, "", timerDuration, task.Desc)
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
		service.TskUpdate(tsk, false, args)

		if listAfterChange {
			service.TkListAfterChange() // TODO Lebeda - whow only work task
		}
	},
}

func init() {
	taskCmd.AddCommand(tkWorkCmd)

	// TODO Lebeda - add functions for worklog and timer
	tkWorkCmd.Flags().BoolP("worklog", "w", false, "automatic start worklog with group, code and description from task")
	tkWorkCmd.Flags().StringP("timer", "t", "", "add new timer with desc from task")
	// TODO Lebeda - text worklog and timer

	// options for worklog
	tkWorkCmd.Flags().StringP("worklog-before", "b", "", "Time shift worklog")

	curDate := time.Now()
	tkWorkCmd.Flags().String("worklog-time", curDate.Format("15:04"), "Time of begin worklog")
	tkWorkCmd.Flags().String("worklog-date", curDate.Format("2006-01-02"), "Time of begin worklog")

}
