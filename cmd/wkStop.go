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
	"github.com/spf13/viper"
	"time"
)

// wkStopCmd represents the wkStop command
var wkStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop work on current task",
	//Long: ``, TODO Lebeda - add description
	Run: func(cmd *cobra.Command, args []string) {
		service.WrkStop(wkBeforeOpt, wkTimeOpt, wkDateOpt)
		if listAfterChange {
			workList := service.WrkGetWork(now.BeginningOfDay(), now.EndOfDay(), false)
			termout.WrkListWork(workList, true)
		}
		if viper.GetString("afterchange") != "" {
			service.SysAfterChange()
		}
	},
}

func init() {
	workCmd.AddCommand(wkStopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wkStopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wkStopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	wkStopCmd.Flags().StringVarP(&wkBeforeOpt, "before", "b", "", "Time shift of record")

	curDate := time.Now()
	wkStopCmd.Flags().StringVarP(&wkTimeOpt, "time", "t", curDate.Format("15:04"), "Time of begin record")
	wkStopCmd.Flags().StringVar(&wkDateOpt, "date", curDate.Format("2006-01-02"), "Time of begin record")
}
