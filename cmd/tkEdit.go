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
	"github.com/martinlebeda/taskmaster/service"
	"github.com/spf13/cobra"
)

var selectByCategory, selectByCode bool

// tkUpdateCmd represents the tkEdit command
var tkUpdateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"edt", "edit", "upd"},
	Short:   "update values of task",
	Args:    cobra.MinimumNArgs(1), // TODO Lebeda - nemá smysl pro více ID
	Long: `usage: tm tk update [-d description] ID [ID ID ID ...]
	
	In task description can use priority, project and contexts in todo.txt format.
	Duration and description will be set for all IDs`,
	Run: func(cmd *cobra.Command, args []string) {
		service.TskUpdate(taskOpt, args)

		if listAfterChange {
			service.TskListAfterChange()
		}
	},
}

func init() {
	taskCmd.AddCommand(tkUpdateCmd)

	// TODO Lebeda - zajistit jednotné nápovědy pomocí konstant
	tkUpdateCmd.Flags().StringVarP(&taskOpt.Desc, "desc", "d", "", "task description")

	// TODO Lebeda - domyslet tkUpdateCmd.Flags().BoolVar(&tkPrioCleanOpt, "clean-priority", false, "force task priority (clean if set empty)")

	// TODO Lebeda - addFlagsSelectBy method for add flags on one place
	tkUpdateCmd.Flags().BoolVar(&selectByCategory, "by-category", false, "arguments are groups instead ID")
	tkUpdateCmd.Flags().BoolVar(&selectByCode, "by-code", false, "arguments are codes instead ID")
}
