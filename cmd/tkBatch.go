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
	"bufio"
	"github.com/martinlebeda/taskmaster/service"
	"github.com/martinlebeda/taskmaster/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// tkBatchCmd represents the tkAdd command
var tkBatchCmd = &cobra.Command{
	Use: "batch",
	// TODO Lebeda - join with import command
	//Args:  cobra.ExactArgs(1),
	Aliases: []string{"feed"},
	Short:   "batch add multiple task",
	Long: `usage: tm tk batch [FILE ...]
	If not any file on command line, stdin is used.`,
	Run: func(cmd *cobra.Command, args []string) {
		//taskOpt.Desc = args[0]
		if len(args) > 0 {
			for _, arg := range args {
				// open the file filepath
				f, err := os.Open(arg)
				tools.CheckErr(err)
				fs := bufio.NewScanner(f)
				processLines(fs)
			}
		} else {
			fs := bufio.NewScanner(os.Stdin)
			processLines(fs)
		}

		if listAfterChange {
			service.TskListAfterChange()
		}
		if viper.GetString("exportafterchange") != "" {
			service.TskExportTasks(viper.GetString("exportafterchange"), []string{})
		}
		if viper.GetString("afterchange") != "" {
			service.SysAfterChange()
		}
	},
}

func processLines(scanner *bufio.Scanner) {
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		taskOpt.Desc = scanner.Text()
		//fmt.Println(taskOpt)
		service.TskAdd(taskOpt)
	}
}

func init() {
	taskCmd.AddCommand(tkBatchCmd)

}
