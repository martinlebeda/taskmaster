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
	"github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/service"
	"github.com/martinlebeda/taskmaster/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tkDoneCmd represents the tkDone command
var tkDeferCmd = &cobra.Command{
	Use:     "defer",
	Aliases: []string{"asdone", "shelve", "pass"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "unset task priority and set state as normal and optionaly append context",
	Long:    `usage: tm tk defer ID [ID ID ID ...]`,
	Run: func(cmd *cobra.Command, args []string) {
		service.TskPrio("", args)

		var tsk model.Task
		tsk.Status = "N"
		tsk.DateDone = tools.GetZeroTime()
		service.TskUpdate(tsk, args)

		part, err := cmd.Flags().GetString("context")
		tools.CheckErr(err)
		if part != "" {
			service.TskAppend(part, args)
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
	taskCmd.AddCommand(tkDeferCmd)

	tkDeferCmd.Flags().StringP("context", "o", "", "context (or any string) for append to task")
}
