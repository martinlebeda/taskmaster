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
	"github.com/spf13/cobra"
	"time"
)

var wkCategoryOpt, wkCodeOpt, wkBeforeOpt, wkTimeOpt, wkDateOpt string

// wkStartCmd represents the wkStart command
var wkStartCmd = &cobra.Command{
	Use:   "start",
	Short: "record start of task",
	Args:  cobra.ExactArgs(1),
	//Long: ``, TODO Lebeda - add description
	//Long: `A longer description that spans multiple lines and likely contains examples to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		service.WrkStart(args[0], wkCategoryOpt, wkCodeOpt, wkBeforeOpt, wkTimeOpt, wkDateOpt)
		if wklistAfterChange {
			workList := service.WrkGetWork(now.BeginningOfDay(), now.EndOfDay())
			termout.WrkListWork(workList)
		}
	},
}

func init() {
	workCmd.AddCommand(wkStartCmd)

	wkStartCmd.Flags().StringVarP(&wkCategoryOpt, "category", "g", "", "Category of record")
	wkStartCmd.Flags().StringVarP(&wkCodeOpt, "code", "e", "", "External code of record")

	wkStartCmd.Flags().StringVarP(&wkBeforeOpt, "before", "b", "", "Time shift of record")

	curDate := time.Now()
	wkStartCmd.Flags().StringVarP(&wkTimeOpt, "time", "t", curDate.Format("15:04"), "Time of begin record")
	wkStartCmd.Flags().StringVar(&wkDateOpt, "date", curDate.Format("2006-01-02"), "Time of begin record")
}
