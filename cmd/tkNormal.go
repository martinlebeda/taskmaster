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
var tkNormalCmd = &cobra.Command{
	Use:     "normal",
	Aliases: []string{"nrm", "workable", "active"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "set task status to 'normal'",
	Long:    `usage: tm tk normal ID [ID ID ID ...]`,
	Run: func(cmd *cobra.Command, args []string) {
		var tsk model.Task
		tsk.Status = "N"
		tsk.DateDone = tools.GetZeroTime()
		service.TskUpdate(tsk, args)

		if listAfterChange {
			service.TskListAfterChange()
		}
		if viper.GetString("exportafterchange") != "" {
			service.TskExportTasks(viper.GetString("exportafterchange"))
		}
	},
}

func init() {
	taskCmd.AddCommand(tkNormalCmd)

	// TODO Lebeda - replace by query on task
	tkNormalCmd.Flags().BoolVar(&selectByCategory, "by-category", false, "arguments are groups instead ID")
	tkNormalCmd.Flags().BoolVar(&selectByCode, "by-code", false, "arguments are codes instead ID")
}
