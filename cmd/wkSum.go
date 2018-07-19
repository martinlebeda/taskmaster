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

var wkSumBy string

// wkSumCmd represents the wkList command
var wkSumCmd = &cobra.Command{
	Use:   "sum",
	Short: "sumarize work report",
	//Aliases: []string{"ls"},
	//Long: ``, TODO Lebeda - add descroption
	Run: func(cmd *cobra.Command, args []string) {
		timeFrom, err := time.Parse("2006-01-02 15:04", wkListFromOpt)
		tools.CheckErr(err)
		timeTo, err := time.Parse("2006-01-02 15:04", wkListToOpt)
		tools.CheckErr(err)
		workSum := service.WrkGetWorkSum(timeFrom, timeTo, wkSumBy)
		termout.WrkSumWork(workSum)
	},
}

func init() {
	workCmd.AddCommand(wkSumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wkSumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wkSumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	wkSumCmd.Flags().StringVarP(&wkSumBy, "group-by", "g", "desc", "group by category/code/desc")

	wkSumCmd.Flags().StringVarP(&wkListFromOpt, "from", "f", now.BeginningOfDay().Format("2006-01-02 15:04"), "Record start after date and time")
	wkSumCmd.Flags().StringVarP(&wkListToOpt, "to", "t", now.EndOfDay().Format("2006-01-02 15:04"), "Record start before date and time")

}
