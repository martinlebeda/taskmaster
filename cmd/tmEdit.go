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
	"github.com/martinlebeda/taskmaster/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

// tmEditCmd represents the tmEdit command
var tmEditCmd = &cobra.Command{
	Use:   "edit", // TODO Lebeda - rename edit to update
	Short: "Edit timer items",
	Args:  cobra.MinimumNArgs(1),
	Long:  `usage: tm tm update [--note 'note'] [--goal '--goal'] ID [ID ID ID ...]`,
	Run: func(cmd *cobra.Command, args []string) {
		note, err := cmd.Flags().GetString("note")
		tools.CheckErr(err)

		goalOpt, err := cmd.Flags().GetString("goal")
		tools.CheckErr(err)
		goal, err := tools.ParseDateTimeMinutes(goalOpt)
		tools.CheckErr(err)

		service.TmrUpdate(note, goal, args)
		if listAfterChange {
			service.TmrListAfterChange()
		}
		if viper.GetString("afterchange") != "" {
			service.SysAfterChange()
		}
	},
}

func init() {
	timerCmd.AddCommand(tmEditCmd)

	tmEditCmd.Flags().String("note", "", "new note value")
	tmEditCmd.Flags().String("goal", time.Time{}.Format(tools.BaseDateTimeFormat), "new note value")
}
