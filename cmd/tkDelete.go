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
)

// tkDeleteCmd represents the tkDelete command
var tkDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "delete task",
	Args:    cobra.MinimumNArgs(1),
	Long: `Delete tasks by ID.
	
	usage: tm tk del ID [ID ID ID ...]`,
	Run: func(cmd *cobra.Command, args []string) {
		service.TskDel(args)
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

func init() {
	taskCmd.AddCommand(tkDeleteCmd)
}
