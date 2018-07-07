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
    "github.com/spf13/cobra"
    "github.com/martinlebeda/taskmaster/service"
    "github.com/martinlebeda/taskmaster/termout"
    "fmt"
)

var tmPastOpt, tmNextOpt, tmCntOpt, tmCndOut, tmNotify bool

// tmListCmd represents the tmList command
var tmListCmd = &cobra.Command{
    Use:     "list",
    Aliases: []string{"ls"},
    Short:   "list timer",
    Long:    `List of timer records.`,
    Args:    cobra.NoArgs,
    Run: func(cmd *cobra.Command, args []string) {
        timerDistances := service.TmrGetDistance(tmPastOpt, tmNextOpt)

        // notify
        if tmNotify {
            for _, distance := range timerDistances {
                if distance.Distance < 0 {
                    service.SysNotifyDistance(distance)
                }
            }
        }

        // output
        if tmCntOpt {
            fmt.Println(len(timerDistances))
        } else {
            termout.TmrListDistance(timerDistances, tmCndOut)
        }
    },
}

func init() {
    timerCmd.AddCommand(tmListCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // tmListCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    tmListCmd.Flags().BoolVarP(&tmNextOpt, "next", "n", false, "list only next timer")
    tmListCmd.Flags().BoolVarP(&tmPastOpt, "past", "p", false, "list only past timers")
    tmListCmd.Flags().BoolVar(&tmCntOpt, "count", false, "print only count of timers")
    tmListCmd.Flags().BoolVar(&tmCndOut, "condensed", false, "print timers without ID and no big spaces (ie for statusbar)")

    tmListCmd.Flags().BoolVar(&tmNotify, "notify", false, "system notify when past")
}
