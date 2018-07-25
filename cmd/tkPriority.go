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

var tkPrioCleanOpt bool

// tkDoneCmd represents the tkDone command
var tkPriorityCmd = &cobra.Command{
	Use:     "priority",
	Aliases: []string{"prio", "p"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "set (or unset) task priority",
	// TODO Lebeda - add long description
	//Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if tkPrioCleanOpt {
			taskOpt.Prio.String = ""
		}

		service.TskUpdate(taskOpt, tkPrioCleanOpt, args)

		if listAfterChange {
			service.TkListAfterChange()
		}
	},
}

func init() {
	taskCmd.AddCommand(tkPriorityCmd)

	tkPriorityCmd.Flags().StringVarP(&taskOpt.Prio.String, "prio", "p", "", "task priority")
	tkPriorityCmd.Flags().BoolVarP(&tkPrioCleanOpt, "clean-priority", "c", false, "clean task priority")
}
