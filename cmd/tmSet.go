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
	"github.com/spf13/viper"
	"time"
)

var dateOpt string
var tmSetTagOpt string
var tmSetReplaceTagOpt bool

// tmSetCmd represents the tmSet command
var tmSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set timer for date and time",
	Args:  cobra.ExactArgs(2),
	Long:  `Set timer for concrete date and time.`,
	Run: func(cmd *cobra.Command, args []string) {
		service.TmrSet(tmSetReplaceTagOpt, tmSetTagOpt, dateOpt, args[0], args[1])
		if listAfterChange {
			service.TmrListAfterChange()
		}
		if viper.GetString("afterchange") != "" {
			service.SysAfterChange()
		}
	},
}

func init() {
	timerCmd.AddCommand(tmSetCmd)

	curDate := time.Now()
	tmSetCmd.Flags().StringVar(&dateOpt, "date", curDate.Format("2006-01-02"), "date for timer, default is current date")

	tmSetCmd.Flags().StringVarP(&tmSetTagOpt, "tag", "t", "", "Create timer with tag")
	tmSetCmd.Flags().BoolVar(&tmSetReplaceTagOpt, "replace-tag", false, "replace all timers with tag by this new")
}
