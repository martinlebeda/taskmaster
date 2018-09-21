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

// tkRemoveCmd represents the tkRemove command
var tkRemoveCmd = &cobra.Command{
	Use:   "remove",
	Args:  cobra.MinimumNArgs(2),
	Short: "remove project or context to task description",
	Long: `usage: tm tk remove '@context' ID [ID ID ID ...]

Remove any text from description ie @context or +project or any string you want.`,
	Run: func(cmd *cobra.Command, args []string) {
		part := args[0]
		taskIds := args[1:]

		service.TskRemove(part, taskIds)

		if listAfterChange {
			service.TskListAfterChange()
		}
		if viper.GetString("afterchange") != "" {
			service.SysAfterChange()
		}
		if viper.GetString("exportafterchange") != "" {
			service.TskExportTasks(viper.GetString("exportafterchange"))
		}
	},
}

func init() {
	taskCmd.AddCommand(tkRemoveCmd)
}
