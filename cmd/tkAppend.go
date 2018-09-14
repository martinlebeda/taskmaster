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
)

// tkAppendCmd represents the tkAppend command
var tkAppendCmd = &cobra.Command{
	Use:     "append",
	Aliases: []string{"project", "context"},
	Args:    cobra.MinimumNArgs(2),
	Short:   "append project or context to task description",
	Long: `usage: tm tk append '@context' ID [ID ID ID ...]

For appended text use todo.txt syntax ie @context or +project or any string you want.`,
	Run: func(cmd *cobra.Command, args []string) {
		part := args[0]
		taskIds := args[1:]

		service.TskAppend(part, taskIds)

		if listAfterChange {
			service.TskListAfterChange()
		}
	},
}

func init() {
	taskCmd.AddCommand(tkAppendCmd)
}
