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
	"github.com/martinlebeda/taskmaster/service"
	"github.com/martinlebeda/taskmaster/termout"
	"github.com/martinlebeda/taskmaster/tools"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var wkBeforeOpt, wkTimeOpt, wkDateOpt string

// wkStartCmd represents the wkStart command
var wkStartCmd = &cobra.Command{
	Use:   "start",
	Short: "record start of task",
	Args:  cobra.ExactArgs(1),
	Long: `usage: tm wk start 'description'
     tm wk start -T ID_TASK
     tm wk start -h ID_WORK
`,
	Run: func(cmd *cobra.Command, args []string) {

		byTask, err := cmd.Flags().GetBool("by-task")
		tools.CheckErr(err)

		byHistory, err := cmd.Flags().GetBool("by-history")
		tools.CheckErr(err)

		desc := args[0]
		if byTask {
			id, err := strconv.Atoi(args[0])
			tools.CheckErr(err)

			task := service.TskGetById(id)
			desc = service.RemovePrioFromDesc(task.Desc)
		}
		if byHistory {
			id, err := strconv.Atoi(args[0])
			tools.CheckErr(err)

			workList := service.WrkGetWorkById(id)
			desc = workList.Desc
		}

		service.WrkStart(desc, wkBeforeOpt, wkTimeOpt, wkDateOpt)

		if listAfterChange {
			workList := service.WrkGetWork(now.BeginningOfDay(), now.EndOfDay(), false)
			termout.WrkListWork(workList)
		}
	},
}

func init() {
	workCmd.AddCommand(wkStartCmd)

	wkStartCmd.Flags().StringVarP(&wkBeforeOpt, "before", "b", "", "Time shift of record")

	curDate := time.Now()
	wkStartCmd.Flags().StringVarP(&wkTimeOpt, "time", "t", curDate.Format("15:04"), "Time of begin record")
	wkStartCmd.Flags().StringVar(&wkDateOpt, "date", curDate.Format("2006-01-02"), "Time of begin record")

	wkStartCmd.Flags().BoolP("by-task", "T", false, "argument is ID of task and description used from task")
	wkStartCmd.Flags().BoolP("by-history", "H", false, "argument is ID of log and description used from this log")
}
