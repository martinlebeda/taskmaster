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
)

// tkDoneCmd represents the tkDone command
var tkMaybeCmd = &cobra.Command{
	Use:     "maybe",
	Aliases: []string{"mb", "nonworkable", "inactive"},
	Short:   "set task status to 'maybe'",
	Args:    cobra.MinimumNArgs(1),
	Long:    `usage: tm tk maybe ID [ID ID ID ...]`,
	Run: func(cmd *cobra.Command, args []string) {
		var tsk model.Task
		tsk.Status = "M"
		tsk.DateDone = tools.GetZeroTime()
		service.TskUpdate(tsk, args)

		if listAfterChange {
			service.TskListAfterChange()
		}
	},
}

func init() {
	taskCmd.AddCommand(tkMaybeCmd)

	tkMaybeCmd.Flags().BoolVar(&selectByCategory, "by-category", false, "arguments are groups instead ID")
	tkMaybeCmd.Flags().BoolVar(&selectByCode, "by-code", false, "arguments are codes instead ID")
}
