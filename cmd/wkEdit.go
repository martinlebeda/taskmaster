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

// wkEditCmd represents the wkEdit command
var wkEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit worklog items",
	Args:  cobra.MinimumNArgs(1),
	// TODO Lebeda - add long description
	//	Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		code, err := cmd.Flags().GetString("code")
		service.CheckErr(err)

		category, err := cmd.Flags().GetString("category")
		service.CheckErr(err)

		desc, err := cmd.Flags().GetString("desc")
		service.CheckErr(err)

		startOpt, err := cmd.Flags().GetString("start")
		service.CheckErr(err)
		start, err := time.Parse(service.BaseDateTimeFormat, startOpt)
		service.CheckErr(err)

		stopOpt, err := cmd.Flags().GetString("stop")
		service.CheckErr(err)
		stop, err := time.Parse(service.BaseDateTimeFormat, stopOpt)
		service.CheckErr(err)

		service.WrkUpdate(code, category, desc, start, stop, args)

		if wklistAfterChange {
			workList := service.WrkGetWork(now.BeginningOfDay(), now.EndOfDay(), false)
			termout.WrkListWork(workList)
		}
	},
}

func init() {
	workCmd.AddCommand(wkEditCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wkEditCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wkEditCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	tmEditCmd.Flags().String("code", "", "new code value")
	tmEditCmd.Flags().String("category", "", "new category value")
	tmEditCmd.Flags().String("desc", "", "new desc value")

	tmEditCmd.Flags().String("start", time.Time{}.Format(service.BaseDateTimeFormat), "new start value")
	tmEditCmd.Flags().String("stop", time.Time{}.Format(service.BaseDateTimeFormat), "new stop value")
}