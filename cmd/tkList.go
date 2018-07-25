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
	"time"
)

// tkListCmd represents the tkList command
var tkListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "A brief description of your command",
	// TODO Lebeda - add long description
	//Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		doneOpt, err := cmd.Flags().GetString("done-from")
		tools.CheckErr(err)
		tskDoneFrom, err := time.Parse("2006-01-02", doneOpt)
		tools.CheckErr(err)

		showMaybe, err := cmd.Flags().GetBool("maybe")
		tools.CheckErr(err)

		showNext, err := cmd.Flags().GetBool("next")
		tools.CheckErr(err)

		showPrio, err := cmd.Flags().GetStringArray("prio")
		tools.CheckErr(err)

		showCode, err := cmd.Flags().GetString("code")
		tools.CheckErr(err)

		showCategory, err := cmd.Flags().GetString("category")
		tools.CheckErr(err)

		tasks := service.TskGetList(tskDoneFrom, showMaybe, showPrio, showCode, showCategory, args)

		if showNext {
			termout.TskShowWork(tasks[0])
		} else {
			termout.TskListTasks(tasks)
		}

	},
}

func init() {
	taskCmd.AddCommand(tkListCmd)

	// TODO Lebeda - regex/glob pro prohledávání
	// TODO Lebeda - prio pro prohledávání, spec hodnoty -a -b -c -d
	// TODO Lebeda - prio pro prohledávání arra jako výstup

	tkListCmd.Flags().BoolP("next", "x", false, "show only work or next task")

	tkListCmd.Flags().StringArrayP("prio", "p", []string{}, "show only priority")

	tkListCmd.Flags().BoolP("maybe", "m", false, "show maybe tasks")
	tkListCmd.Flags().String("done-from", now.BeginningOfDay().Format("2006-01-02"), "show done from day")
	tkListCmd.Flags().StringP("code", "c", "", "show tasks with code")
	tkListCmd.Flags().StringP("category", "g", "", "show tasks with category")
}
