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
	"fmt"
	"github.com/jinzhu/now"
	"github.com/martinlebeda/taskmaster/service"
	"github.com/martinlebeda/taskmaster/termout"
	"github.com/martinlebeda/taskmaster/tools"
	"github.com/spf13/cobra"
	"time"
	"unicode"
)

// tkListCmd represents the tkList command
var tkListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list tasks",
	Long: `usage: tm tk ls [flags...] [patterns...]
	
	Pattern is string for search in description.
	If empty, list all tasks. Between more patterns is use OR operator.`,
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

		showPrioA, err := cmd.Flags().GetBool("prio-a")
		tools.CheckErr(err)
		if showPrioA {
			showPrio = append(showPrio, getAllPrioTo('A')...)
		}

		showPrioB, err := cmd.Flags().GetBool("prio-b")
		tools.CheckErr(err)
		if showPrioB {
			showPrio = append(showPrio, getAllPrioTo('B')...)
		}

		showPrioTo, err := cmd.Flags().GetString("prio-to")
		tools.CheckErr(err)
		if showPrioTo != "" {
			runes := []rune(showPrioTo)
			showPrio = append(showPrio, getAllPrioTo(runes[0])...)
		}

		showPrioZ, err := cmd.Flags().GetBool("prio-exists")
		tools.CheckErr(err)
		if showPrioZ {
			showPrio = append(showPrio, getAllPrioTo('Z')...)
		}

		showStatus, err := cmd.Flags().GetString("status")
		tools.CheckErr(err)

		tasks := service.TskGetList(tskDoneFrom, showMaybe, showStatus, showPrio, args)

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

	tkListCmd.Flags().BoolP("next", "x", false, "show only work or next task")

	tkListCmd.Flags().StringArrayP("prio", "p", []string{}, "show only priority")
	tkListCmd.Flags().BoolP("prio-a", "a", false, "show only priority A")
	tkListCmd.Flags().BoolP("prio-b", "b", false, "show only priority A-B")
	tkListCmd.Flags().StringP("prio-to", "t", "", "show only priority defined or less (ie. C = A-C)")
	tkListCmd.Flags().BoolP("prio-exists", "z", false, "show only task with any priority")

	tkListCmd.Flags().BoolP("maybe", "m", false, "show maybe tasks")
	tkListCmd.Flags().String("done-from", now.BeginningOfDay().Format("2006-01-02"), "show done from day")
	tkListCmd.Flags().StringP("status", "s", "", "show tasks with status")
}

// get all priority
func getAllPrioTo(prioTo rune) []string {
	prioTo = unicode.ToUpper(prioTo)
	var result []string
	for i := 'A'; i <= prioTo; i++ {
		result = append(result, fmt.Sprintf("%c", i))
	}
	return result
}
